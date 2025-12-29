---
description: >-
  Swap crypto across chains with Vultisig. Integrated with THORChain, Maya,
  1inch, Li.Fi, and Kyberswap for secure cross-chain exchanges.
---

# Swapping

Vultisig supports swapping assets both within the same chain and across different blockchains. All swaps are secured by your devices using threshold signatures.

***

## How to Swap

1. **Tap the "Swap" button** from your vault view
2. **Select the asset to swap from** (source)
3. **Select the asset to receive** (destination)
4. **Enter the amount** you want to swap
5. **Review the quote** - including rate, fees, and estimated output
6. **Initiate [signing](../signing-a-transaction/signing-a-transaction.md)** with your threshold of devices
7. **Monitor the swap** - cross-chain swaps depend on blockchain finality

<figure><img src="../../.gitbook/assets/image (1).png" alt=""><figcaption><p>Swap Flow</p></figcaption></figure>

***

## Swap Providers

Vultisig automatically selects the best route using these integrated providers:

| Provider | Type | Speciality |
|----------|------|------------|
| [THORChain](https://thorchain.org/) | Cross-chain DEX | Native cross-chain swaps (BTC, ETH, etc.) |
| [MAYA Protocol](https://www.mayaprotocol.com/) | Cross-chain DEX | Extended asset support |
| [1inch](https://1inch.io/) | DEX Aggregator | Best rates across EVM DEXs |
| [Li.Fi](https://li.fi/) | Bridge Aggregator | Cross-chain bridging |
| [Kyberswap](https://kyberswap.com) | DEX Aggregator | EVM chain swaps |

***

## Cross-Chain vs Same-Chain

* **Same-chain swaps** (e.g., ETH → USDC on Ethereum): Fast, single transaction
* **Cross-chain swaps** (e.g., BTC → ETH): Uses THORChain or Maya, time depends on blockchain finality

***

## Tips

* **Compare quotes** - Vultisig shows the best available rate
* **Check slippage** - larger swaps may have higher slippage
* **Cross-chain swaps have streaming** - large swaps are split for better rates
* **Transaction fees** are shown before confirmation

***

## Related

* [Signing a Transaction](../signing-a-transaction/signing-a-transaction.md)
* [Sending](sending.md)
* [Functions (DeFi)](functions.md)
