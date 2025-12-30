---
description: >-
  Cosmos IBC in Vultisig: transfer assets between IBC-enabled chains, switch
  assets to THORChain, and execute custom transactions.
---

# Cosmos (IBC)

{% hint style="warning" %}
**Migration in Progress**

Cosmos IBC features are being migrated to the DeFi tab. Currently, these functions are still accessible via the **Functions** button on each respective chain in the Wallet tab.
{% endhint %}

Vultisig supports IBC (Inter-Blockchain Communication) transfers for Cosmos ecosystem chains. Users can move assets between IBC-enabled chains and execute custom transactions directly from the DeFi tab.

***

## Supported Chains

| Chain | Features |
|-------|----------|
| **Cosmos Hub** | IBC Transfer, Switch, Custom |
| **Kujira** | IBC Transfer, Custom |
| **Osmosis** | IBC Transfer, Custom |

***

## IBC Transfer

Transfer assets between IBC-enabled Cosmos chains.

### How to Transfer

1. Open **DeFi tab** → Select chain (Cosmos/Kujira/Osmosis)
2. Select **IBC Transfer**
3. Enter the **destination chain**
4. Enter the **destination address**
5. Enter the **amount**
6. (Optional) Add a memo
7. Sign with your devices

***

## Switch (Cosmos Hub)

The Switch function transfers assets from Cosmos Hub to THORChain. This is part of the RUJIRA token migration flow.

### How to Switch

1. Open **DeFi tab** → **Cosmos** → **Switch**
2. Select the **asset to switch**
3. Enter the **amount**
4. Sign with your devices

{% hint style="info" %}
For details on the Switch function and RUJI merge, see the [RUJIRA documentation](https://docs.rujira.network/understanding-ruji-token#merge-flow).
{% endhint %}

***

## Custom Transactions

Advanced users can create custom memo transactions to interact with each Cosmos chain's ecosystem.

### How to Execute Custom Transaction

1. Open **DeFi tab** → Select chain
2. Select **Custom**
3. Enter the **memo** according to chain specifications
4. Sign with your devices

{% hint style="warning" %}
Custom transactions are for advanced users only. Incorrect memos can result in errors or loss of funds.
{% endhint %}

***

## Related

- [DeFi Overview](README.md)
- [THORChain](thorchain.md)
