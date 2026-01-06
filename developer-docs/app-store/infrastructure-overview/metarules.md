# Rules System: MetaRules vs Direct Rules

## TL;DR

Vultisig's automation engine uses two rule types: **MetaRules** (high-level like `ethereum.send` or `solana.swap`) that automatically expand to protocol-specific implementations, and **Direct Rules** (low-level, ABI/IDL-mapped like `ethereum.erc20.transfer`) for precise control. Rules support seven constraint types (FIXED, MAX, MIN, MAGIC_CONSTANT, ANY, REGEXP, UNSPECIFIED) and use Magic Constants from `vultisig/recipes` for transparent address resolution. All app transactions require Verifier approval based on these rules before TSS signing.


## Overview

Vultisig's policy engine supports two rule types that govern how apps interact with blockchain networks:

- **MetaRules**: High-level, chain-agnostic abstractions automatically mapped to specific implementations.
- **Direct Rules**: Low-level, protocol-specific rules that strictly follow blockchain or contract semantics.

All app transactions must comply with these rules and be approved by the Verifier before signing.

***

## MetaRules: Protocol-Agnostic Abstractions

MetaRules summarize cross-chain actions (like send, swap) via simplified resource identifiers such as `{chain}.send`. The system expands these abstractions into concrete, protocol-specific Direct Rules tailored for each network.

### Example MetaRules

| MetaRule | Expands to (Examples)                                               | Chains Supported |
| :-- |:--------------------------------------------------------------------| :-- |
| `ethereum.send` | `ethereum.eth.transfer` (native), `ethereum.erc20.transfer` (ERC20) | Ethereum |
| `solana.send` | `solana.system.transfer`, `solana.spl_token.transfer`               | Solana |
| `bitcoin.send` | `bitcoin.btc.transfer`                                              | Bitcoin |
| `ethereum.swap` | 1inch, Uniswap V2/V3 calls                                          | Ethereum |
| `solana.swap` | Jupiter aggregator calls                                            | Solana |
| `ethereum.bridge` | LiFi, Across, deBridge, Native L2 bridge calls                      | Ethereum, Arbitrum, Optimism, Base, Polygon, etc. |

### MetaRule Expansion (Go Example)

```go
func (m *MetaRule) TryFormat(resource string, constraints map[string]*types.Constraint) ([]string, error) {
	parts := strings.Split(resource, ".")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid meta-rule format: %s", resource)
	}
	chain := parts[0]
	protocol := metaProtocol(parts[1])
	switch protocol {
	case send:
		return m.expandSendRule(chain, constraints)
	case swap:
		return m.expandSwapRule(chain, constraints)
	case bridge:
		return m.expandBridgeRule(chain, constraints)
	default:
		return nil, fmt.Errorf("unsupported meta-protocol: %s", protocol)
	}
}
```

***

## Direct Rules: Protocol-Specific Operations

Direct Rules are mapped directly to the target protocol's ABI (EVM) or IDL (Solana), providing fine-grained policy enforcement for app calls.

### Rule Structure (`protobuf`)

```protobuf
message Rule {
  string resource = 1;
  Effect effect = 2;
  string description = 3;
  map<string, Constraint> constraints = 4;
  Authorization authorization = 5;
  string id = 6;
  repeated ParameterConstraint parameter_constraints = 7;
  Target target = 13;
}
```


### Example Direct Rules

#### ERC-20 Transfer (Ethereum)

```json
{
  "resource": "ethereum.erc20.transfer",
  "effect": "EFFECT_ALLOW",
  "target": {
    "target_type": "TARGET_TYPE_ADDRESS",
    "address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
  },
  "parameter_constraints": [
    {
      "parameter_name": "to",
      "constraint": {
        "type": "CONSTRAINT_TYPE_WHITELIST",
        "whitelist": ["0x742b15...a92E"],
        "required": true
      }
    }
  ]
}
```


#### Solana SPL Token Transfer

```json
{
  "resource": "solana.spl_token.transfer",
  "effect": "EFFECT_ALLOW",
  "parameter_constraints": [
    {
      "parameter_name": "amount",
      "constraint": {
        "type": "CONSTRAINT_TYPE_RANGE",
        "min_value": "1000000",
        "max_value": "100000000",
        "required": true
      }
    }
  ]
}
```


***

## Supported Blockchains \& Protocols

| Chain | Features Supported |
| :-- | :-- |
| Ethereum | ERC-20, Uniswap, 1inch, custom contracts |
| Bitcoin | Native BTC transfers |
| Solana | SPL tokens, Jupiter aggregator, Metaplex |
| Polygon | Polygon bridges, DEXs |
| ARB, BSC etc | Layer2 \& EVM-specific features |
| THORChain | Cross-chain liquidity/swaps |
| XRPL | XRP Ledger native operations |


***

## Constraint Types

Based on `constraint.proto`:


| Type | Description                                               |
| :-- |:----------------------------------------------------------|
| CONSTRAINT_TYPE_FIXED | Enforce fixed value for parameter                         |
| CONSTRAINT_TYPE_MAX | Maximum allowed value                                     |
| CONSTRAINT_TYPE_MIN | Minimum allowed value                                     |
| CONSTRAINT_TYPE_MAGIC_CONSTANT | Special constant (system address, treasury, router, etc.) |
| CONSTRAINT_TYPE_ANY | Accept any value                                          |
| CONSTRAINT_TYPE_REGEXP | Regular expression match for string params                |
| CONSTRAINT_TYPE_UNSPECIFIED | Not specified - usually treated as deny                   |

### Magic Constants
Magic Constants are **predefined system addresses and values** maintained in the `vultisig/recipes` library that provide standardized references to important protocol addresses, treasury wallets, and infrastructure endpoints. They enable transparent, auditable automations while simplifying the user experience.

| Name | Value | Purpose                    |
| :-- | :-- |:---------------------------|
| VULTISIG_TREASURY | 1 | Treasury address           |
| THORCHAIN_VAULT | 2 | Router/vault for THORChain |
| THORCHAIN_ROUTER | 3 | Router for THORChain swaps |
| LIFI_ROUTER | 6 | LiFi cross-chain aggregator router |
| ARBITRUM_L1_GATEWAY | 7 | Arbitrum L1 Gateway Router (on Ethereum) |
| OPTIMISM_L1_BRIDGE | 8 | Optimism L1 Standard Bridge (on Ethereum) |
| BASE_L1_BRIDGE | 9 | Base L1 Standard Bridge (on Ethereum) |
| ARBITRUM_L2_GATEWAY | 10 | Arbitrum L2 Gateway Router |
| OPTIMISM_L2_BRIDGE | 11 | Optimism L2 Standard Bridge |
| BASE_L2_BRIDGE | 12 | Base L2 Standard Bridge |
| ACROSS_SPOKE_POOL | 15 | Across Protocol SpokePool contract |
| DEBRIDGE_DLN_SOURCE | 16 | deBridge DLN Source contract |

#### What Are Magic Constants?

Magic Constants are symbolic references that resolve to actual blockchain addresses or protocol-specific values at runtime. Instead of hardcoding addresses directly in rules or requiring users to input complex addresses, developers reference these constants by name, and the Vultisig system automatically resolves them to the correct values based on the target chain and current protocol state.

**Key characteristics:**

- **Centrally Defined**: All magic constants are defined and maintained in the `vultisig/recipes` repository, ensuring consistency across all apps
- **Chain-Aware**: The same constant can resolve to different addresses on different chains (e.g., VULTISIG_TREASURY resolves to the appropriate treasury address for Ethereum, Solana, etc.)
- **Version-Safe**: When protocol addresses change (upgrades, migrations), constants are updated centrally without breaking existing app rules
- **Auditable**: Users can verify exactly which addresses are being used by inspecting the recipes library


#### When to Use Magic Constants

Use magic constants in these scenarios:

1. **Protocol Infrastructure Addresses**
    - DEX routers (Uniswap, THORChain, Jupiter)
    - Vault contracts
    - Official protocol treasury addresses
2. **Cross-Chain Operations**
    - Bridge contracts
    - Cross-chain routers
    - Liquidity pool addresses
3. **User-Friendly Automations**
    - When users shouldn't need to know specific addresses
    - For operations that always use the same infrastructure
    - To prevent user errors from incorrect address input

***

## Rule Development Guidelines

- **MetaRules**: Use for simple, cross-chain actions or consistent business logic.
- **Direct Rules**: Apply for advanced protocol features, fine-grained control, or where ABI details matter.
- **Constraints**: Use principle of least privilege; always review all allowed parameters and document constraints.
- **Testing**: Validate with protocol ABIs/IDLs and ensure all value constraints are respected before allow.
- **Magic Constants**: Use system-defined constants for seamless app integrations and parameter substitutions.

