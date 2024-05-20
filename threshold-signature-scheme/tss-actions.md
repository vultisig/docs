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
* They are simple wallet addresses called "externally-owned accounts" (EOA)



{% hint style="success" %}
**Importantly:** The individual secrets (vault shares) do not contain funds.\
The actual private key to the vault never exists in a normal vault user flow. \
Thus, individual vault shares can be safely emailed, saved, uploaded to websites, and more.

In fact, no part of the vault share indicates where the final vault is - there is no linkage between the vault share and final vault. So someone who discovers a single vault share will not know what to do with it or where to look.
{% endhint %}

{% hint style="danger" %}
It is critical that individual Vault shares are never stored together, otherwise a malicious party could recombine the Vault shares back into the Vault:

1. Do not back up more than 1 Vault share to the same device, email, Google Drive, or iCloud.
2. Do not upload more than 1 Vault share to the same website.
{% endhint %}

## **Key-Signing**

Key Signing is a process where a threshold (Vultisig TSS chooses 67%) of the parties must come together in another ceremony to again prove access to a secret and generate a signed transaction object. The transaction can then be sent to a crypto network. \
For example, for a 2 of 3 TSS vault, only 2 of the devices need to be together to sign an outgoing transaction.

## **Re-share**

Re-share is a process by which the signing threshold of a vault can work to kick out a non-responsive device and add another device, or even upgrade the vault from 2-of-2 to 3-of-4 (or downgrade).
