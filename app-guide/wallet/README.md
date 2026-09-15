---
description: >-
  Send, receive, and swap from the Vultisig Wallet tab across 36+ chains.
---

# Vultisig Wallet tab

The Wallet tab is the primary interface for managing crypto assets. It displays holdings across supported blockchains and the core transaction actions.

***

## Features

| Function | Description |
| --- | --- |
| **View assets** | Balances across chains in one view |
| **Receive** | Addresses and QR codes for deposits |
| **Send** | Transfer to an address. Address book and destination tags where the chain needs them |
| **Swap** | Same-chain and cross-chain. Vultisig picks the best quote |
| **Buy** | Fiat on-ramp on supported coins (not QBTC) |
| **Chain functions** | Extra actions via **Functions** on chains that offer them |

***

## Chain functions

Some chains expose extra actions beyond send / receive / swap. Open the chain in the Wallet tab, then tap **Functions**.

| Chain | Functions |
| --- | --- |
| [Cosmos Hub, Osmosis](../defi/cosmos.md) | IBC transfer |
| [dYdX](../defi/dydx.md) | Vote on governance proposals |
| Bitcoin, Bitcoin Cash, Litecoin, Dogecoin, Ethereum, Avalanche, BSC, Base, Ripple | Add THORChain LP |
| MayaChain | Leave, custom memo |

TON stake / unstake and Solana / Terra staking are on the [DeFi tab](../defi/), not Functions.

Kujira may still appear on some iOS builds. The network is wound down and is removed from the SDK. Do not treat it as a current chain.

***

## Supported chains

The apps expose these networks from one vault. Addresses are derived automatically.

* **UTXO:** Bitcoin, Bitcoin Cash, Litecoin, Dogecoin, Dash, Zcash
* **EVM:** Ethereum, Arbitrum, Optimism, Base, Polygon, BSC, Avalanche, Blast, Cronos, zkSync, Mantle, Hyperliquid, Sei, Robinhood
* **Cosmos:** Cosmos Hub, Osmosis, dYdX, Terra, Terra Classic, Noble, Akash, THORChain, MayaChain
* **Other:** Solana, Polkadot, Bittensor, TON, Ripple, Tron, Cardano, Sui

QBTC appears in some builds as a post-quantum testnet. It is not a WalletCore signing chain. Do not send mainnet funds there.

That is 36+ mainnet networks. The SDK `Chain` enum is the same set plus QBTC.

***

## Transaction security

Verify the destination, amount, and (on EVM) the decoded function on the sign screen. [Blockaid](sending.md#blockaid) can scan the transaction before you sign. Enable it in Settings.

***

## Related

* [Send](sending.md)
* [Swap](swapping.md)
* [DeFi](../defi/)
* [Referrals](../referrals.md)
