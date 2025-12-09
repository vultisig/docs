package main

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	gcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	v1 "github.com/vultisig/commondata/go/vultisig/vault/v1"
	"github.com/vultisig/mobile-tss-lib/tss"
	"github.com/vultisig/recipes/sdk/evm"
	"github.com/vultisig/verifier/plugin/keysign"
	vtypes "github.com/vultisig/verifier/types"
	"github.com/vultisig/verifier/vault"
	"github.com/vultisig/vultiserver/contexthelper"
	"github.com/vultisig/vultisig-go/address"
	"github.com/vultisig/vultisig-go/common"
)

const (
	// PluginID is the unique identifier for this plugin.
	// It must match the ID registered in the Vultisig Verifier database.
	PluginID = "<your-plugin-id>"
)

// Trigger is the main service struct that handles external requests to initiate transactions.
// It acts as a bridge between an HTTP API and the Vultisig signing infrastructure.
type Trigger struct {
	logger                *logrus.Logger
	vaultStorage          vault.Storage   // Access to the encrypted S3/Minio vault storage
	vaultEncryptionSecret string          // Secret key used to decrypt the vault data
	signer                *keysign.Signer // Vultisig client for requesting TSS signatures
	eth                   *evm.SDK        // Helper SDK for constructing Ethereum transactions
}

// NewTrigger initializes the Trigger service.
// It sets up the Ethereum SDK connection which is required for building and broadcasting transactions.
func NewTrigger(
	logger *logrus.Logger,
	vaultStorage vault.Storage,
	vaultEncryptionSecret string,
	ethRpc *ethclient.Client,
	signer *keysign.Signer,
) *Trigger {
	// Determine the Chain ID (Mainnet, Sepolia, etc.) from the common configuration
	ethEvmChainID, err := common.Ethereum.EvmID()
	if err != nil {
		return nil
	}
	// Initialize the SDK with the correct chain ID and RPC client
	eth := evm.NewSDK(ethEvmChainID, ethRpc, ethRpc.Client())
	return &Trigger{
		logger:                logger,
		signer:                signer,
		vaultStorage:          vaultStorage,
		vaultEncryptionSecret: vaultEncryptionSecret,
		eth:                   eth,
	}
}

// Wait starts a simple HTTP server to listen for transaction triggers.
// This function blocks until the context is cancelled.
// ENDPOINT: GET /trigger
// PARAMS: uuid (policy ID), pubkey, fromAddress (to), amount
func (tr *Trigger) Wait(ctx context.Context) {
	http.HandleFunc("/trigger", func(w http.ResponseWriter, r *http.Request) {
		// --- Request Validation ---
		uuidStr := r.URL.Query().Get("uuid")
		if uuidStr == "" {
			http.Error(w, "Missing uuid parameter", http.StatusBadRequest)
			return
		}

		pubkey := r.URL.Query().Get("pubkey")
		if pubkey == "" {
			http.Error(w, "Missing pubkey parameter", http.StatusBadRequest)
			return
		}

		toAddress := r.URL.Query().Get("fromAddress") // Note: Logic uses this as 'toAddress' despite param name
		if toAddress == "" {
			http.Error(w, "Missing fromAddress parameter", http.StatusBadRequest)
			return
		}

		amount := r.URL.Query().Get("amount")
		if amount == "" {
			http.Error(w, "Missing amount parameter", http.StatusBadRequest)
			return
		}

		// --- Transaction Execution ---
		err := tr.createTx(ctx,
			uuidStr,
			pubkey,
			toAddress,
			amount,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Tx triggered for UUID: %s\n", uuidStr)
	})

	srv := &http.Server{
		Addr:    ":8090",
		Handler: nil,
	}

	// Start HTTP server in a goroutine
	go func() {
		log.Println("Server starting on port", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("Server error: %v", err)
		}
	}()

	// Wait for shutdown signal
	select {
	case <-ctx.Done():
		log.Println("Shutting down server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		srv.Shutdown(shutdownCtx)
	}
}

// createTx orchestrates the entire flow of creating and signing a transaction.
// Steps:
// 1. Fetch and decrypt the Vault (to get the sender's address).
// 2. Construct an unsigned Ethereum transaction.
// 3. Request a TSS signature from the Vultisig network.
func (tr *Trigger) createTx(c context.Context, policyUuid, pubkey, toAddress, amount string) error {
	ctx, cancel := context.WithTimeout(c, 5*time.Minute)
	defer cancel()

	if err := contexthelper.CheckCancellation(ctx); err != nil {
		tr.logger.WithError(err).Warn("Context cancelled, skipping trigger")
		return err
	}

	policyId, err := uuid.Parse(policyUuid)
	if err != nil {
		return err
	}

	pluginPolicy := vtypes.PluginPolicy{
		ID:        policyId,
		PublicKey: pubkey,
		PluginID:  PluginID,
	}

	// Step 1: Retrieve the vault.
	// We need the vault to derive the on-chain address associated with the public key.
	vault, err := getVaultFromPolicy(tr.vaultStorage, pluginPolicy, tr.vaultEncryptionSecret)
	if err != nil {
		return fmt.Errorf("failed to get vault from policy: %w", err)
	}

	// Derive the Ethereum address from the vault's ECDSA public key.
	ethAddress, _, _, err := address.GetAddress(vault.PublicKeyEcdsa, vault.HexChainCode, common.Ethereum)
	if err != nil {
		return fmt.Errorf("failed to get eth address: %w", err)
	}

	chain := common.Ethereum
	amountInt, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return fmt.Errorf("invalid amount")
	}

	// Step 2: Generate the unsigned transaction (payload).
	tx, e := tr.genUnsignedTx(
		ctx,
		ethAddress,
		toAddress,
		amountInt,
	)
	if e != nil {
		return fmt.Errorf("p.genUnsignedTx: %w", e)
	}

	// Step 3: Wrap the transaction in a Vultisig Keysign Request.
	// This object tells the TSS nodes "Sign this data for this policy on this chain".
	signRequest, err := vtypes.NewPluginKeysignRequestEvm(
		pluginPolicy, "", chain, tx)
	if err != nil {
		return fmt.Errorf("vtypes.NewPluginKeysignRequestEvm: %w", e)
	}

	// Initiate the signing process.
	err = tr.signTx(ctx, *signRequest)
	if err != nil {
		return fmt.Errorf("failed to init sign: %w", err)
	}

	return nil
}

// getVaultFromPolicy fetches the encrypted vault file from storage and decrypts it.
// This is necessary because the raw public key alone isn't enough to sign; we need the
// vault context (though for address derivation, we just need the keys).
func getVaultFromPolicy(s vault.Storage, policy vtypes.PluginPolicy, encryptionSecret string) (*v1.Vault, error) {
	vaultFileName := common.GetVaultBackupFilename(policy.PublicKey, policy.PluginID.String())
	vaultContent, err := s.GetVault(vaultFileName)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault")
	}

	if vaultContent == nil {
		return nil, fmt.Errorf("vault not found")
	}

	return common.DecryptVaultFromBackup(encryptionSecret, vaultContent)
}

// signTx handles the communication with the Vultisig signing nodes.
// It sends the request, waits for the signature, and then broadcasts the result.
func (tr *Trigger) signTx(
	ctx context.Context,
	req vtypes.PluginKeysignRequest,
) error {
	// Send request to TSS nodes.
	sigs, err := tr.signer.Sign(ctx, req)
	if err != nil {
		tr.logger.WithError(err).Error("Keysign failed")
		return fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Validate we received exactly one signature (standard for EVM).
	if len(sigs) != 1 {
		tr.logger.
			WithField("sigs_count", len(sigs)).
			Error("expected only 1 message+sig per request for evm")
		return fmt.Errorf("failed to sign transaction: invalid signature count: %d", len(sigs))
	}
	var sig tss.KeysignResponse
	for _, s := range sigs {
		sig = s
	}

	// Broadcast the signed transaction to the blockchain.
	_, err = tr.broadcastTx(ctx, sig, req)
	if err != nil {
		tr.logger.WithError(err).Error("failed to complete signing process (broadcast tx)")
		return fmt.Errorf("failed to complete signing process: %w", err)
	}

	return nil
}

// genUnsignedTx uses the EVM SDK to build the raw transaction bytes.
// It handles nonce management, gas estimation, and RLP encoding internally.
func (tr *Trigger) genUnsignedTx(
	ctx context.Context,
	senderAddress string,
	toAddress string,
	amount *big.Int,
) ([]byte, error) {
	tx, err := tr.eth.MakeTxTransferNative(
		ctx,
		gcommon.HexToAddress(senderAddress),
		gcommon.HexToAddress(toAddress),
		amount,
	)
	if err != nil {
		return nil, fmt.Errorf("p.eth.MakeTxTransferNative: %v", err)
	}
	return tx, nil
}

// broadcastTx combines the original unsigned transaction with the signature
// (R, S, V components) and submits it to the network.
func (tr *Trigger) broadcastTx(
	ctx context.Context,
	signature tss.KeysignResponse,
	signRequest vtypes.PluginKeysignRequest,
) (*types.Transaction, error) {
	// Decode the original transaction bytes stored in the request.
	txBytes, err := base64.StdEncoding.DecodeString(signRequest.Transaction)
	if err != nil {
		return nil, fmt.Errorf("failed to decode b64 proposed tx: %w", err)
	}
	txHex := gcommon.Bytes2Hex(txBytes)

	// Reconstruct the signed transaction using the TSS signature parts.
	// Note: RecoveryID is used to calculate the 'V' value for EIP-155.
	tx, err := tr.eth.Send(
		ctx,
		txBytes,
		gcommon.Hex2Bytes(signature.R),
		gcommon.Hex2Bytes(signature.S),
		gcommon.Hex2Bytes(signature.RecoveryID),
	)
	if err != nil {
		tr.logger.WithError(err).WithField("tx_hex", txHex).Error("p.eth.Send")
		return nil, fmt.Errorf("p.eth.Send(tx_hex=%s): %w", txHex, err)
	}

	tr.logger.WithFields(logrus.Fields{
		"from_public_key": signRequest.PublicKey,
		"to_address":      tx.To().Hex(),
		"hash":            tx.Hash().Hex(),
		"chain":           common.Ethereum.String(),
	}).Info("tx successfully signed and broadcasted")
	return tx, nil
}
