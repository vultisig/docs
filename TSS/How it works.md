---
description: How TSS works
---

# How it works

Threshold Signature Scheme (TSS) is a multi-party computation protocol that enables different computers (devices) with separate "secrets" to come together in a ceremony to perform three functions:

1. Key-generation
2. Key-signing
3. Re-sharing

**Key-gen** is a process by which the devices prove access to a secret, and generate a shared public key. 100% of all devices must be online. The shared public key which is generated can then be used to create on-chain addresses to receive funds - the "vault".

The vault addresses do not look like special contracts or scripts; they are simple wallet addresses "externally-owned accounts (EOAs).&#x20;

{% hint style="success" %}
Importantly: the individual secrets (vault shares) do not contain funds. The actual private key to the vault never exists in a normal vault user flow. Thus the individual vault shares are safe to be emailed, stored, uploaded to websites and more.&#x20;

In fact, no part of the vault share indicates where the final vault is - there is no linkage between the vault share and final vault. Thus someone who discovers an individual vault share will not know what to do with it or where to look.&#x20;
{% endhint %}

{% hint style="danger" %}
It is critically important that individual vault shares are never stored together; else a malicous party could recombine the vault shares into the vault:

1\) Do not backup more than 1 vault share to the same device, email, google drive or iCloud.

2\) Do not upload more than 1 vault share to the same website
{% endhint %}

**Key-sign** is a process by which a threshold (Voltix TSS chooses 67%) of the parties must come together in another ceremony to once again prove access to a secret and generated a signed transaction object. The transaction can then be broadcast to a crypto network.&#x20;

Eg, for a 2-of-3 TSS vault, only 2 of the devices need to be together to sign an outgoing transaction.&#x20;

**Re-share** is a process by which the signing theshold of a vault can work to kick out a non-responsive device and add a different device, or even upgrade the vault from 2-of-2 to 3-of-4, (or downgrade).&#x20;

