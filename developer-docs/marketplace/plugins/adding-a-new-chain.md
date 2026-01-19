---
description: The how to steps to add a missing chains to the Marketplace ecosystem
---

# Adding a New Chain to the Vultisig Plugin Ecosystem

This document provides a high-level overview of the steps required to add support for a new blockchain across the Vultisig Marketplace ecosystem. The process involves multiple repositories and follows a dependency chain where each layer builds upon the previous.

***

## Overview

Adding a new chain requires changes across five main repositories:

1. **vultisig-go** - Core address generation and chain definitions
2. **recipes** - Transaction parsing, validation rules, and policy engine
3. **verifier** - Transaction indexing and status monitoring
4. **app-recurring** - Automation services (send/swap execution)
5. **plugin-marketplace** - Frontend UI support

***

## Step by step Guide

_Note that if you run into any issues please contact us in the Vultisig discord._

### 1. Vultisig-go (Foundation Layer)

**Purpose**: Provides the foundational chain definitions and address derivation logic.

#### What Needs to Be Added

**Chain Enumeration (`common/chain.go`)**

* Add the new chain to the `Chain` enum constant
* Add the chain to `chainToString` mapping for serialization
* Add the derivation path to `chainDerivePath` mapping (HD wallet path)
* Add the native symbol to `NativeSymbol()` method
* If EVM-compatible, add the chain ID to `EvmID()` method
* If EdDSA-based (like Solana, Sui), update `IsEdDSA()` method

**Address Generation (`address/`)**

* Create a new file (e.g., `newchain.go`) implementing address derivation
* Implement the address encoding specific to the chain (e.g., bech32, base58check, hex)
* Add the chain case to `GetAddress()` in `address.go` to route to the new implementation
* Create corresponding test file with address derivation test vectors

#### Trust Wallet Core Dependency

**Important**: Address generation may or may not require Trust Wallet Core support depending on the chain type:

* **ECDSA chains** (Bitcoin-like, EVM): Use `tss.GetDerivedPubKey()` from `mobile-tss-lib` for key derivation, then custom address encoding
* **EdDSA chains** (Solana, Sui, etc.): Typically use the root public key directly without derivation
* **Cosmos chains**: Use bech32 encoding with chain-specific HRP (human-readable prefix)
* **Custom curves**: May require Trust Wallet Core or custom implementation

The `mobile-tss-lib` package handles the TSS key derivation, while address encoding is typically implemented directly in vultisig-go.

***

### 2. Recipes (Policy & Validation Layer)

**Purpose**: Defines how transactions are parsed, validated, and what policy rules can be applied.

#### What Needs to Be Added

**Chain Implementation (`chain/{chaintype}/`)**

Create a new directory structure (or add to existing category like `utxo/`, `evm/`):

1. **`chain.go`** - Implements the `types.Chain` interface:
   * `ID()` - Unique chain identifier (lowercase, e.g., "zcash")
   * `Name()` - Human-readable name
   * `Description()` - Brief description
   * `SupportedProtocols()` - List of protocol IDs (e.g., \["zec"])
   * `ParseTransaction()` - Decode raw transaction bytes into structured format
   * `GetProtocol()` - Return protocol handler by ID
   * `ComputeTxHash()` - Compute transaction hash with signatures applied
2. **`decode.go`** - Transaction deserialization logic:
   * Parse raw transaction bytes into chain-specific transaction structure
   * Handle version detection, input/output parsing, signatures
3. **`protocol.go`** - Implements the `types.Protocol` interface:
   * Define supported functions (transfer, swap, etc.)
   * Implement `MatchFunctionCall()` for policy constraint validation
   * Extract and validate transaction parameters against policy rules

**Registry Registration (`chain/registry.go`)**

* Import the new chain package
* Add `RegisterChain(newchain.NewChain())` in `init()`

**Engine Implementation (`engine/{chaintype}/`)**

Create the rule evaluation engine:

1. **Engine struct** - Wraps chain-specific validation logic:
   * `Supports()` - Return true for the supported chain(s)
   * `Evaluate()` - Validate a transaction against policy rules
   * Parse transaction, extract parameters, check constraints
2. **Constraint validation** - Validate output addresses, amounts, data fields

**Engine Registry (`engine/registry.go`)**

* Import the new engine package
* Add engine registration in `NewChainEngineRegistry()`

#### Why Each Component Is Needed

| Component | Purpose                                                  |
| --------- | -------------------------------------------------------- |
| Chain     | Identifies the blockchain and routes to correct handlers |
| Protocol  | Defines what operations are possible (transfer, swap)    |
| Engine    | Validates transactions match allowed policy rules        |
| Decoder   | Parses raw bytes into structured transaction data        |

***

### 3. Verifier (Transaction Monitoring Layer)

**Purpose**: Tracks transaction status on-chain and computes final transaction hashes.

#### What Needs to Be Added

**TX Indexer Chain Implementation (`plugin/tx_indexer/pkg/chain/`)**

1. **`newchain.go`** - Implements `chain.Indexer` interface:
   * `ComputeTxHash()` - Compute transaction hash from proposed transaction and signatures
   * Typically wraps the recipes chain implementation

**RPC Client (`plugin/tx_indexer/pkg/rpc/`)**

1. **`newchain.go`** - Blockchain RPC client:
   * `GetTxStatus()` - Query transaction confirmation status
   * Return `TxOnChainPending`, `TxOnChainSuccess`, or `TxOnChainFailed`
   * Use appropriate API (native RPC, block explorer API like Blockchair, etc.)

**Chains List (`plugin/tx_indexer/chains_list.go`)**

* Register the new chain in the supported chains list
* Wire up the indexer and RPC client

#### What the Verifier Checks

The verifier performs several critical functions:

1. **Policy Validation** - Ensures transaction matches the approved policy rules
2. **Transaction Tracking** - Monitors transaction from proposal to confirmation
3. **Hash Computation** - Computes final transaction hash after signatures are applied
4. **Status Monitoring** - Polls blockchain for confirmation status

***

### 4. App-recurring (Automation Layer)

**Purpose**: Executes automated transactions (sends, swaps) based on scheduled policies.

#### What Needs to Be Added

**Network Package (`internal/{chain}/`)**

Create a new directory with:

1. **`types.go`** - Define chain-specific types:
   * `From` struct (address, amount, public key)
   * `To` struct (chain, asset ID, address)
   * `TxInput`, `TxOutput` structures
   * Any chain-specific interfaces (e.g., `TxBroadcaster`)
2. **`network.go`** - Network service orchestration:
   * Initialize all chain services
   * Provide `Send()` and `Swap()` entry points
   * Coordinate between services
3. **`send_service.go`** - Build send transactions:
   * `BuildTransfer()` - Create transaction outputs for a simple transfer
   * Handle change output calculation
4. **`swap_service.go`** - Build swap transactions:
   * `FindBestAmountOut()` - Query swap providers for best rate
   * Aggregate results from multiple providers
5. **`swap_provider.go`** - DEX/swap provider integration:
   * `MakeOutputs()` - Build transaction outputs for swap
   * Integrate with Thorchain, MayaChain, or other DEXs
6. **`signer_service.go`** - Transaction signing and broadcast:
   * `SignAndBroadcast()` - Coordinate TSS signing
   * Build keysign request with message hashes
   * Apply signatures and broadcast to network
7. **`client.go`** - Blockchain RPC client:
   * `GetUTXOs()` - Fetch unspent outputs (for UTXO chains)
   * `GetBalance()` - Query account balance
   * `BroadcastTransaction()` - Submit signed transaction
   * `GetFeeRate()` - Query current network fees
8. **`fee_provider.go`** - Fee estimation:
   * Query network for current fee rates
   * Calculate transaction fees
9. **`address.go`** (if needed) - Address utilities:
   * Address validation
   * Script generation (P2PKH, P2SH, etc.)

**Consumer Integration (`internal/recurring/consumer.go`)**

1. **Add Network to Consumer struct**:
   * Add field for the new chain's Network service
   * Add to `NewConsumer()` constructor
2. **Add pubkey-to-address helper**:
   * `newchainPubToAddress()` - Derive chain address from vault public key
3. **Add operation handlers**:
   * `handleNewchainSend()` - Execute send operations
   * `handleNewchainSwap()` - Execute swap operations
4. **Update handle() routing**:
   * Add chain detection in the main handler
   * Route to appropriate send/swap handler

**Worker Initialization (`cmd/worker/main.go`)**

* Initialize the new chain's Network service
* Add to DCA Consumer constructor
* Configure RPC endpoints and dependencies

**Supported Chains (`internal/recurring/spec.go`)**

* Add the chain to `supportedChains` slice
* This enables dynamic policy generation for the new chain

#### What the App Does

The app-recurring service:

1. **Schedules** - Triggers transactions based on policy schedule (hourly, daily, etc.)
2. **Builds Transactions** - Constructs appropriate transaction for send or swap
3. **Signs** - Coordinates TSS signing across vault participants
4. **Broadcasts** - Submits signed transaction to the blockchain
5. **Tracks** - Monitors transaction until confirmation

***

### 5. Plugin-marketplace (Frontend Layer)

**Purpose**: Provides the user interface for creating and managing policies.

#### What Needs to Be Added

**Chain Configuration (`src/utils/chain.ts`)**

1. **Add to chain category**:
   * For EVM: Add to `evmChains` and `evmChainInfo` with viem chain config
   * For UTXO: Add to `utxoChains`
   * For other: Add to `otherChains`
2. **Add ticker**: Add native token ticker to `tickers` mapping
3. **Add RPC URL** (if EVM): Add to `evmRpcUrls` mapping

**Storage (`src/storage/chain.ts`)**

* Usually no changes needed - uses the chain types from utils

**UI Components**

Depending on chain characteristics, may need:

* Token images in `public/tokens/`
* Chain-specific validation logic
* Custom UI for chain-specific features

***

### Summary Checklist

#### vultisig-go

* [ ] Add chain enum and mappings
* [ ] Implement address derivation
* [ ] Add tests with known test vectors

#### recipes

* [ ] Create chain implementation (ID, name, protocols)
* [ ] Implement transaction parser/decoder
* [ ] Create protocol with function definitions
* [ ] Implement constraint matching
* [ ] Create validation engine
* [ ] Register in chain and engine registries

#### verifier

* [ ] Create chain indexer for hash computation
* [ ] Create RPC client for status checking
* [ ] Register in chains list

#### app-recurring

* [ ] Create network package with all services
* [ ] Integrate swap provider (Thorchain/MayaChain if supported)
* [ ] Add to consumer with send/swap handlers
* [ ] Initialize in worker
* [ ] Add to supported chains list

#### plugin-marketplace

* [ ] Add chain definition
* [ ] Add ticker mapping
* [ ] Add any chain-specific UI components

***

### Chain Type Patterns

#### UTXO Chains (Bitcoin, Zcash, Litecoin, etc.)

* Require UTXO management (fetching, selection, change calculation)
* Transaction format: inputs referencing previous outputs, new outputs
* Typically use base58check or bech32 addresses
* May have chain-specific features (Zcash's shielded transactions, BCH's CashAddr)

#### EVM Chains (Ethereum, Arbitrum, Base, etc.)

* Simpler: single account model, no UTXO
* Use standard Ethereum address format (0x prefixed hex)
* Can share most infrastructure with existing EVM implementation
* Just need chain ID and RPC configuration

#### Account-Based Non-EVM (Solana, XRP, etc.)

* Account model but different transaction format
* May require chain-specific SDKs
* Often have unique address formats

***

### Testing Recommendations

1. **Unit Tests**: Address derivation from known test vectors
2. **Integration Tests**: Transaction building and parsing round-trips
3. **End-to-End Tests**: Full send/swap flow on testnet
4. **Policy Tests**: Constraint validation with various policy rules
