---
description: >-
  Swap crypto across chains with Vultisig. Integrated with THORChain, Maya,
  1inch, Li.Fi, and Kyberswap for secure cross-chain exchanges.
---

# Swapping

Vultisig supports swapping assets both within the same chain and across different blockchains. All swaps are secured by threshold signatures.

***

## How to Swap

1. **Tap "Swap"** from your vault view
2. **Select the asset to swap from** (source)
3. **Select the asset to receive** (destination)
4. **Enter the amount** to swap
5. **Review the quote** — rate, fees, and estimated output
6. **Sign with your devices**
7. **Monitor the swap** — cross-chain swaps depend on blockchain finality

<figure><img src="../../.gitbook/assets/image (1).png" alt=""><figcaption><p>Swap Flow</p></figcaption></figure>

***

## Swap Providers

Vultisig automatically selects the best route using integrated providers:

| Provider | Type | Speciality |
|----------|------|------------|
| [THORChain](https://thorchain.org/) | Cross-chain DEX | Native cross-chain swaps (BTC, ETH, etc.) |
| [MAYA Protocol](https://www.mayaprotocol.com/) | Cross-chain DEX | Extended asset support |
| [1inch](https://1inch.io/) | DEX Aggregator | Best rates across EVM DEXs |
| [Li.Fi](https://li.fi/) | Bridge Aggregator | Cross-chain bridging |
| [Kyberswap](https://kyberswap.com) | DEX Aggregator | EVM chain swaps |

***

## Cross-Chain vs Same-Chain

| Type | Example | Speed |
|------|---------|-------|
| **Same-chain** | ETH → USDC on Ethereum | Fast, single transaction |
| **Cross-chain** | BTC → ETH | Depends on blockchain finality |

Cross-chain swaps use THORChain or Maya Protocol. Large swaps are automatically streamed for better rates.

***

## Tips

- **Compare quotes** — Vultisig shows the best available rate
- **Check slippage** — larger swaps may have higher slippage
- **Streaming swaps** — large amounts are split for better execution
- **Transaction fees** are shown before confirmation

***

## Related

- [Keysign](../../security-technology/keysign.md) — How transaction signing works
- [Sending](sending.md)
- [DeFi](../defi/README.md)
