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
	// PluginID defined in verifier database 'plugins'. You can find it in your local version while testing or request
	// from Vultisig team before going live
	PluginID = "<your-plugin-id>"
)

type Trigger struct {
	logger                *logrus.Logger
	vaultStorage          vault.Storage
	vaultEncryptionSecret string
	signer                *keysign.Signer
	eth                   *evm.SDK
}

func NewTrigger(
	logger *logrus.Logger,
	vaultStorage vault.Storage,
	vaultEncryptionSecret string,
	ethRpc *ethclient.Client,
	signer *keysign.Signer,
) *Trigger {
	ethEvmChainID, err := common.Ethereum.EvmID()
	if err != nil {
		return nil
	}
	eth := evm.NewSDK(ethEvmChainID, ethRpc, ethRpc.Client())
	return &Trigger{
		logger:                logger,
		signer:                signer,
		vaultStorage:          vaultStorage,
		vaultEncryptionSecret: vaultEncryptionSecret,
		eth:                   eth,
	}
}

func (tr *Trigger) Wait(ctx context.Context) {
	http.HandleFunc("/trigger", func(w http.ResponseWriter, r *http.Request) {
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

		toAddress := r.URL.Query().Get("fromAddress")
		if toAddress == "" {
			http.Error(w, "Missing fromAddress parameter", http.StatusBadRequest)
			return
		}

		amount := r.URL.Query().Get("amount")
		if amount == "" {
			http.Error(w, "Missing amount parameter", http.StatusBadRequest)
			return
		}

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

	go func() {
		log.Println("Server starting on port", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("Server error: %v", err)
		}
	}()

	select {
	case <-ctx.Done():
		log.Println("Shutting down server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		srv.Shutdown(shutdownCtx)
	}
}

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

	vault, err := getVaultFromPolicy(tr.vaultStorage, pluginPolicy, tr.vaultEncryptionSecret)
	if err != nil {
		return fmt.Errorf("failed to get vault from policy: %w", err)
	}

	ethAddress, _, _, err := address.GetAddress(vault.PublicKeyEcdsa, vault.HexChainCode, common.Ethereum)
	if err != nil {
		return fmt.Errorf("failed to get eth address: %w", err)
	}

	chain := common.Ethereum
	amountInt, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return fmt.Errorf("invalid amount")
	}

	tx, e := tr.genUnsignedTx(
		ctx,
		ethAddress,
		toAddress,
		amountInt,
	)
	if e != nil {
		return fmt.Errorf("p.genUnsignedTx: %w", e)
	}

	signRequest, err := vtypes.NewPluginKeysignRequestEvm(
		pluginPolicy, "", chain, tx)
	if err != nil {
		return fmt.Errorf("vtypes.NewPluginKeysignRequestEvm: %w", e)
	}

	err = tr.signTx(ctx, *signRequest)
	if err != nil {
		return fmt.Errorf("failed to init sign: %w", err)
	}

	return nil
}

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

func (tr *Trigger) signTx(
	ctx context.Context,
	req vtypes.PluginKeysignRequest,
) error {
	sigs, err := tr.signer.Sign(ctx, req)
	if err != nil {
		tr.logger.WithError(err).Error("Keysign failed")
		return fmt.Errorf("failed to sign transaction: %w", err)
	}

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

	_, err = tr.broadcastTx(ctx, sig, req)
	if err != nil {
		tr.logger.WithError(err).Error("failed to complete signing process (broadcast tx)")
		return fmt.Errorf("failed to complete signing process: %w", err)
	}

	return nil
}

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

func (tr *Trigger) broadcastTx(
	ctx context.Context,
	signature tss.KeysignResponse,
	signRequest vtypes.PluginKeysignRequest,
) (*types.Transaction, error) {
	txBytes, err := base64.StdEncoding.DecodeString(signRequest.Transaction)
	if err != nil {
		return nil, fmt.Errorf("failed to decode b64 proposed tx: %w", err)
	}
	txHex := gcommon.Bytes2Hex(txBytes)

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
