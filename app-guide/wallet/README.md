---
description: >-
  Vultisig Wallet tab: send, receive, and swap crypto across 30+ blockchains.
  Multi-chain asset management with threshold signature security.
---

# Wallet

The Wallet tab is the primary interface for managing crypto assets. It displays all holdings across supported blockchains and provides access to core transaction functions.

***

## Features

| Function | Description |
|----------|-------------|
| **View Assets** | See balances across all chains in one view |
| **Receive** | Generate addresses and QR codes for deposits |
| **Send** | Transfer assets to any address |
| **Swap** | Exchange assets within or across chains |

***

## Supported Chains

Vultisig supports 30+ blockchains including:

- **UTXO**: Bitcoin, Litecoin, Dogecoin, Bitcoin Cash, Dash
- **EVM**: Ethereum, Arbitrum, Optimism, Base, Polygon, BSC, Avalanche
- **Cosmos**: Cosmos Hub, Osmosis, Kujira, Dydx, THORChain, Maya
- **Other**: Solana, Polkadot, Sui, TON, Ripple

All chains are accessible from a single vault. Each vault generates addresses for all supported chains automatically.

***

## Transaction Security

All transactions from the Wallet tab require threshold signatures:

1. Transaction is initiated on one device
2. Other devices join the signing session
3. Each device independently verifies and approves
4. Threshold signatures combine to create a valid transaction
5. Transaction broadcasts to the blockchain

No single device ever holds the complete private key. See [How Keysigning Works](../../security-technology/keysign.md) for technical details.

***

## Guides

{% content-ref url="sending.md" %}
[sending.md](sending.md)
{% endcontent-ref %}

{% content-ref url="swapping.md" %}
[swapping.md](swapping.md)
{% endcontent-ref %}
