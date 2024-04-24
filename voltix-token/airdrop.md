---
description: Aggressively buying Users and AUM
---

# Airdrop

Users have an option to register their vaults for an Airdrop. To do this, they send their ECDSA and EdDSA public keys to an airdrop registry, which will scan for presence of funds on chains, and then begin counting Airdrop Value.&#x20;

A new scan will be performed each month, and the User Airdrop Value accumulates. Users that keep the funds the longest in Voltix will earn the highest Airdrop Value.

```
user_airdrop_value += $total_vault_value
```

The final share of the airdrop is the pro-rata share of the Airdrop

```
user_share = 20,000,000 * user_airdrop_value / total_aidrop_value
```

{% hint style="warning" %}
Do not register for the airdrop if you do not wish for your public keys to be sent to the&#x20;
{% endhint %}

<figure><img src="../.gitbook/assets/image.png" alt="" width="143"><figcaption></figcaption></figure>
