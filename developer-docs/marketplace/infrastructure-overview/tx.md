# Transaction Proposing Flow Documentation

## TL;DR

This guide explains how plugin developers in the Vultisig ecosystem can programmatically initiate and sign transactions based on automated conditions (schedules, events, etc.). Plugins must:

1. **Define automation rules** via Recipe specifications that declare supported transaction types, parameters, and constraints
2. **Implement `ProposeTransactions`** to validate policies, fetch vault data, generate unsigned transactions based on recipe rules, index them, and create keysign requests
3. **Implement transaction generation** (`genUnsignedTx`) for specific chains (EVM, etc.) with appropriate parameters
4. **Handle signing** (`initSign` and `SigningComplete`) to request signatures from the signer service and broadcast signed transactions to the network
5. **Pass verifier validation** - all proposed transactions must comply with automation constraints and recipe rules

**Key flow**: Trigger Event → Fetch Automation → Generate Unsigned TX → Create Keysign Request → Verifier Validates → Sign → Broadcast → Index TX 

## Overview

This documentation describes the transaction proposing flow for plugin developers in the Vultisig ecosystem. Plugins have the capability to initiate transactions based on their own conditions, which can be triggered by schedules, events, or other conditions.

## Key Concepts

- **Plugin automation**: Defines the rules and constraints for transactions that a plugin can propose
- **Recipe**: Contains the specific transaction rules derived from the automation
- **Keysign Request**: A request to sign a transaction that meets all automation requirements
- **Verifier**: A secondary participant that validates proposed transactions against the automation

## Transaction Flow Sequence

1. **Trigger Event**: A scheduler, event, or condition triggers the plugin
2. **Automation Retrieval**: The plugin fetches its associated automation
3. **Transaction Proposal**: The plugin generates unsigned transactions based on automation rules
4. **Signing Initiation**: The plugin creates keysign requests for each transaction
5. **Verification & Signing**: The verifier validates and signs the transactions
6. **Broadcast**: Signed transactions are broadcast to the network
7. **Transaction Indexing**: Proposed transactions are recorded in the transaction indexer


### Transaction Proposal

You should implement:
- Plugin automation validation
- Getting pk info from vault
- Processing each rule in the recipe and generate transactions if it's needed
- Create transaction records in the indexer
- Prepare keysign requests for each transaction

Example:
```go 
func (p *Plugin) ProposeTransactions(ctx context.Context, policy vtypes.PluginPolicy) ([]vtypes.PluginKeysignRequest, error) {
	// Plugin policy validation
	err := p.ValidatePluginPolicy(policy)
	if err != nil {
		return nil, fmt.Errorf("failed to validate app policy: %w", err)
	}
    
	// Getting pk info from vault
	vault, err := common.GetVaultFromPolicy(p.vaultStorage, policy, p.vaultEncryptionSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault from policy: %w", err)
	}

	recipe, err := policy.GetRecipe()
	if err != nil {
		return nil, fmt.Errorf("failed to get recipe from policy: %w", err)
	}
	
	var (
		txs = make([]vtypes.PluginKeysignRequest, 0)
	)

	// Processing each rule in the recipe and generate transactions if it's needed
	for _, _rule := range recipe.Rules {
		    if !someRuleBasedCondition {
				continue
            }           
			tx, e := p.genUnsignedTx(
				ctx,
				...
			)
			if e != nil {
				return fmt.Errorf("p.genUnsignedTx: %w", e)
			}

			txHex := ecommon.Bytes2Hex(tx)

			txData, e := ethereum.DecodeUnsignedPayload(tx)
			if e != nil {
				return fmt.Errorf("ethereum.DecodeUnsignedPayload: %w", e)
			}
			txHashToSign := etypes.LatestSignerForChainID(ethEvmID).Hash(etypes.NewTx(txData))

            // Create transaction records in the indexer
			txToTrack, e := p.txIndexerService.CreateTx(ctx, storage.CreateTxDto{
				PluginID:      policy.PluginID,
				PolicyID:      policy.ID,
				ChainID:       vcommon.Chain(chain),
				TokenID:       tokenID,
				FromPublicKey: policy.PublicKey,
				ToPublicKey:   recipient,
				ProposedTxHex: txHex,
			})
			if e != nil {
				return fmt.Errorf("p.txIndexerService.CreateTx: %w", e)
			}

			msgHash := sha256.Sum256(txHashToSign.Bytes())

			// Create signing request
			signRequest := vtypes.PluginKeysignRequest{
				KeysignRequest: vtypes.KeysignRequest{
					PublicKey: policy.PublicKey,
					Messages: []vtypes.KeysignMessage{
						{
							TxIndexerID:  txToTrack.ID.String(),
							Message:      base64.StdEncoding.EncodeToString(txHashToSign.Bytes()),
							Chain:        vcommon.Chain(chain),
							Hash:         base64.StdEncoding.EncodeToString(msgHash[:]),
							HashFunction: vtypes.HashFunction_SHA256,
						},
					},
					PolicyID: policy.ID,
					PluginID: policy.PluginID.String(),
				},
				Transaction: txHex,
			}

			txs = append(txs, signRequest)
			return nil
	}
	
	return txs, nil
}
```

For preparing `signRequest` you may also use builder from *vultisig/verifier* repo:
```go
func NewPluginKeysignRequestEvm(policy PluginPolicy, txToTrack string, chain vgcommon.Chain, tx []byte) (
	*PluginKeysignRequest, error) 
```

### Transaction Generation

You should implement:
- Creating appropriate transactions based on chain and tx type
- Args may vary depending on your tx type and network

Example:
```go
func (p *Plugin) genUnsignedTx(
    ctx context.Context,
    chain rcommon.Chain,
    senderAddress, tokenID, amount, to string,
) ([]byte, error) {
	switch chain {
	case rcommon.Ethereum:
		amt, ok := new(big.Int).SetString(amount, 10)
		if !ok {
			return nil, fmt.Errorf("failed to parse amount: %s", amount)
		}
		
		tx, err := p.eth.MakeAnyTransfer(
			ctx,
			ecommon.HexToAddress(senderAddress),
			ecommon.HexToAddress(to), 
			ecommon.HexToAddress(tokenID), 
			amt, 
			)
		if err != nil {
			return nil, fmt.Errorf("p.eth.MakeAnyTransfer: %v", err)
		}
		return tx, nil
	default:
	}
	return nil, fmt.Errorf("unsupported chain: %s", chain)
}
```


### Signing Process (`initSign` and `SigningComplete`)

These method should handle:
- Requesting signatures from the signer service
- Validating signature responses
- Broadcasting signed transactions to the network

Example:
```go
func (p *Plugin) initSign(
	ctx context.Context,
	req vtypes.PluginKeysignRequest,
	pluginPolicy vtypes.PluginPolicy,
) error {
	sigs, err := p.signer.Sign(ctx, req)
	if err != nil {
		p.logger.WithError(err).Error("Keysign failed")
		return fmt.Errorf("failed to sign transaction: %w", err)
	}

	if len(sigs) != 1 {
		p.logger.
			WithField("sigs_count", len(sigs)).
			Error("expected only 1 message+sig per request for evm")
		return fmt.Errorf("failed to sign transaction: invalid signature count: %d", len(sigs))
	}
	var sig tss.KeysignResponse
	for _, s := range sigs {
		sig = s
	}

	err = p.SigningComplete(ctx, sig, req, pluginPolicy)
	if err != nil {
		p.logger.WithError(err).Error("failed to complete signing process (broadcast tx)")
		return fmt.Errorf("failed to complete signing process: %w", err)
	}
	return nil
}

func (p *Plugin) SigningComplete(
	ctx context.Context,
	signature tss.KeysignResponse,
	signRequest vtypes.PluginKeysignRequest,
	_ vtypes.PluginPolicy,
) error {
	tx, err := p.eth.Send(
		ctx,
		gcommon.FromHex(signRequest.Transaction),
		gcommon.Hex2Bytes(signature.R),
		gcommon.Hex2Bytes(signature.S),
		gcommon.Hex2Bytes(signature.RecoveryID),
	)
	if err != nil {
		p.logger.WithError(err).WithField("tx_hex", signRequest.Transaction).Error("p.eth.Send")
		return fmt.Errorf("p.eth.Send(tx_hex=%s): %w", signRequest.Transaction, err)
	}

	p.logger.WithFields(logrus.Fields{
		"from_public_key": signRequest.PublicKey,
		"to_address":      tx.To().Hex(),
		"hash":            tx.Hash().Hex(),
		"chain":           vcommon.Ethereum.String(),
	}).Info("tx successfully signed and broadcasted")
	return nil
}
```

## Automation Verification

All proposed transactions must pass verification by the verifier service, which will:
- Validate the transaction matches all automation constraints
- Check the transaction parameters against the recipe rules
- Reject any transactions that don't comply with the automation

## Recipe Specification

The plugin defines its supported transaction types through a recipe specification:

```go
func (p *Plugin) GetRecipeSpecification() (*rtypes.RecipeSchema, error)
```

Key aspects:
- Defines supported resources (e.g., Ethereum ERC20 transfers)
- Specifies required parameters (recipient, amount)
- Configures scheduling options (frequency, start date)
- Declares chain support requirements