---
description: Aggressively buying Users and AUM
---

# Airdrop

### Airdrop Process

Users have an option to register their vaults for an Airdrop. To do this, they send their ECDSA and EdDSA public keys to an airdrop registry, which will scan for presence of funds on chains, and then begin counting Airdrop Value.&#x20;

A new scan will be performed each month, and the User Airdrop Value accumulates. Users that keep the funds the longest in Voltix will earn the highest Airdrop Value.

```
user_airdrop_value += $total_vault_value
```

The final share of the airdrop is the pro-rata share of the Airdrop Allocation.

```
user_airdrop_share = 10,000,000 * (user_airdrop_value / total_airdrop_value)
```

<figure><img src="../.gitbook/assets/image (1).png" alt="" width="143"><figcaption></figcaption></figure>

### Initial Airdrop

The initial airdrop is 10,000,000 VLTX. After which, all public keys are purged and a new airdrop period begins.&#x20;

### Ongoing Airdrop

The Airdrop Process will continue for another 5 years, and 2% of the supply of VLTX will be airdropped on the anniversary of the token launch every year.&#x20;

### Privacy

{% hint style="warning" %}
Do not register for the airdrop if you do not wish for your public keys to be sent to the Airdrop Registry. Note, after the Airdrop has finished, your public keys will be purged. No other user information is collected.
{% endhint %}

Read the [#privacy-policy](../other/privacy.md#privacy-policy "mention")
