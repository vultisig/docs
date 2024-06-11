---
description: The purpose of the Vultisig Token
---

# $VULT

Vultisig fulfils four important purposes:

1. To raise capital to accelerate the product during its bootstrap phase
2. To buy users and AUM as quickly as possible using an Airdrop
3. To provide incentives to contributors to build the product
4. To allow anyone to support the adoption trajectory of the product

## The Vultisig Token $VULT

- 100,000,000 starting supply
- 80% in launch liquidity
- 20% allocated to an airdrop
- ERC20 with ERC777 extensions: `approveAndCall()` and `_beforeTransferHook()`
- Ownable, but not mintable (to set launch params)
- Burnable

## Value Accrual

Vultisig operates on a buy-burn model. All affiliate, router, and bridge fees accumulated from platform usage will be used to buy and burn the asset.
{% hint style="info" %}
Through the buy and burn-burn model, the token automatically increases in value the more the app is used to swap and bridge.
{% endhint %}

| Fee                                     | Amount          |
| --------------------------------------- | --------------- |
| Swap Fees (Cross-chain and Token swaps) | 50 basis points |
| Bridge Fees (Between EVMs and IBC)      | 10 basis points |
