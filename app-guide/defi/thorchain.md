---
description: >-
  THORChain DeFi in Vultisig: bond to nodes, stake RUNE/TCY/RUJI, and add LP.
---

# THORChain in Vultisig

Bond, stake, add LP, and merge RUJI from the DeFi tab. You can also add THORChain LP from Wallet → Functions on the listed L1s.

***

## Available Actions

| Action | Asset | Description |
|--------|-------|-------------|
| Bond | RUNE | Bond to THORChain nodes |
| Unbond | RUNE | Remove bond from nodes |
| Stake | RUNE, TCY, yRUNE, yTCY, RUJI | Stake tokens for protocol fees |
| Unstake | RUNE, TCY, yRUNE, yTCY, RUJI | Remove staked tokens |
| Merge / withdraw RUJI | RUJI | RUJIRA merge and withdraw |
| Add LP | Various pairs | Provide liquidity to pools |
| Remove LP | Various pairs | Withdraw liquidity |
| Custom | — | Advanced memo transactions |

***

## Bond / Unbond

Bond RUNE to a THORChain node. Only do this if you are talking to a node operator who has already whitelisted your address.

### How to Bond

1. Open **DeFi tab** → **THORChain** → **Bond**
2. Enter the **node address** (must have whitelisted your address)
3. Enter the **RUNE amount** to bond
4. (Optional) Specify provider and operator fee
5. Sign with your devices

{% hint style="info" %}
Bonding is an advanced feature. Only use if you are in direct contact with a node operator who has whitelisted your address.
{% endhint %}

### How to Unbond

1. Open **DeFi tab** → **THORChain** → **Unbond**
2. Enter the **node address**
3. Enter the **RUNE amount** to unbond
4. (Optional) Specify provider
5. Sign with your devices

### Leave

The "Leave" function is for node operators who wish to exit or disable a node. See [THORChain documentation](https://docs.thorchain.org/thornodes/leaving#leaving) for details.

***

## Stake / unstake

1. Open **DeFi tab** → **THORChain** → **Stake**
2. Select the asset
3. Enter the amount
4. Sign with your devices

Unstake is the same path.

| Asset | Notes |
|-------|-------------|
| **RUNE** | Native THORChain token |
| **TCY** | THORChain yield token |
| **yRUNE** | Yield-bearing RUNE |
| **yTCY** | Yield-bearing TCY |
| **RUJI** | RUJIRA. Tokenomics: [RUJIRA docs](https://docs.rujira.network/understanding-ruji-token) |

***

## RUJI Merge

Merge eligible THORChain tokens into RUJI.

### Merge Tokens

1. Open **DeFi tab** → **THORChain** → **Merge RUJI**
2. Select the **token to merge**
3. Enter the **amount**
4. Sign with your devices

### Withdraw Merged RUJI

After merging, withdraw RUJI from the merge contract:

1. Open **DeFi tab** → **THORChain** → **Withdraw RUJI**
2. Select the **token**
3. Enter the **amount of shares**
4. Sign with your devices

### Withdraw RUJI Rewards

Staked RUJI accumulates USDC rewards. To withdraw:

1. Open **DeFi tab** → **THORChain** → **Withdraw RUJI Rewards**
2. Enter the **USDC amount** to withdraw
3. Sign with your devices

***

## Add THORChain LP from an L1

You can also add THORChain LP from the **Wallet tab** on Bitcoin, Bitcoin Cash, Litecoin, Dogecoin, Ethereum, Avalanche, BSC, Base, and Ripple.

1. Open **Wallet tab** → the chain that holds the asset
2. Tap **Functions** → **Add THORChain LP**
3. Select the pool
4. Enter the amount
5. Sign with your devices

The deposit uses the asset you opened Functions from, unless you change it in the form.

***

## Liquidity Pools

Add or remove liquidity on a THORChain pair.

### How to Add Liquidity

1. Open **DeFi tab** → **THORChain** → **LPs**
2. Select a **trading pair**
3. Enter the **amount** to provide
4. Sign with your devices

### How to Remove Liquidity

1. Open **DeFi tab** → **THORChain** → **LPs**
2. Select your **active position**
3. Enter the **amount** to withdraw
4. Sign with your devices

***

## Custom transactions

Custom memo transactions for THORChain.

{% hint style="warning" %}
Wrong memos can fail or lose funds.
{% endhint %}

***

## Related

- [DeFi Overview](README.md)
- [MayaChain](maya.md)
- [THORChain Docs](https://docs.thorchain.org/)
