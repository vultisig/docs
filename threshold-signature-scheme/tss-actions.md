---
description: Possible Actions
---

# TSS Actions

1. Key-Generation
2. Key-Signing
3. Re-Sharing

## **Key-Generation**

**Key-gen** is a process by which the devices prove access to a secret, and generate a shared public key. 100% of all devices must be online. The shared public key which is generated can then be used to create on-chain addresses to receive funds - the "vault".

* The vault addresses do not look like special contracts or scripts;&#x20;
* They are simple wallet addresses "externally-owned accounts" (EOAs).



{% hint style="success" %}
**Importantly:** The individual secrets (vault shares) do not contain funds.\
The actual private key to the vault never exists in a normal vault user flow. \
Thus the individual vault shares are safe to be emailed, stored, uploaded to websites and more.

In fact, no part of the vault share indicates where the final vault is - there is no linkage between the vault share and final vault. Thus someone who discovers an individual vault share will not know what to do with it or where to look.
{% endhint %}

{% hint style="danger" %}
It is critically important that individual vault shares are never stored together; else a malicous party could recombine the vault shares into the vault:

1\) Do not backup more than 1 vault share to the same device, email, google drive or iCloud.

2\) Do not upload more than 1 vault share to the same website
{% endhint %}

## **Key-Signing**

**Key-sign** is a process by which a threshold (Vultisig TSS chooses 67%) of the parties must come together in another ceremony to once again prove access to a secret and generated a signed transaction object. The transaction can then be broadcast to a crypto network.

Eg, for a 2-of-3 TSS vault, only 2 of the devices need to be together to sign an outgoing transaction.

## **Re-share**

**Re-share** is a process by which the signing theshold of a vault can work to kick out a non-responsive device and add a different device, or even upgrade the vault from 2-of-2 to 3-of-4, (or downgrade).
