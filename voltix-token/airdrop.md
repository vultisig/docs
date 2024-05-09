---
description: Aggressively buying Users and AUM
---

# Airdrop

## TLDR

1. Install Vultisig App and create your vault.
2. Register your vault for the Airdrop. You can register as many vaults as you wish.
3. Deposit funds to your vault, aka use your wallet! Funds are any supported Layer1 or Layer2 assets, and any token on the 1inch token list.
4. Airdrop will be based on your total vault amounts, multiplied by the length of time your funds are kept in the vault. The largest holders for the longest amount of time earn the most.

## Airdrop Process

To register their vaults for the Airdrop, users send their ECDSA and EdDSA public keys to an airdrop registry, which will scan for presence of funds on chains, and then begin counting Airdrop Value.

A new scan will be performed each month, and the User Airdrop Value accumulates. Users that keep the funds the longest in Vultisig will earn the highest Airdrop Value.

$$
user\_airdrop\_value += \$\text{total\_vault\_value}
$$

The final share of the airdrop is the pro-rata share of the Airdrop Allocation.

$$
\text{user\_airdrop\_share} = 10,000,000 \times \left( \frac{\text{user\_airdrop\_value}}{\text{total\_airdrop\_value}} \right)
$$

<figure><img src="../.gitbook\assets\airdrop.png" alt="" width="143"><figcaption></figcaption></figure>

### Initial Airdrop

The initial airdrop is 10,000,000 VLTI. After which, all public keys are purged and a new airdrop period begins.

### Ongoing Airdrop

The Airdrop Process will continue for another 5 years, and 2% of the supply of VLTI will be airdropped on the anniversary of the token launch every year.

### Privacy

{% hint style="warning" %}
Do not register for the airdrop if you do not wish for your public keys to be sent to the Airdrop Registry. Note, after the Airdrop has finished, your public keys will be purged. No other user information is collected.
{% endhint %}

Read the [#privacy-policy](../other/privacy.md#privacy-policy "mention")
