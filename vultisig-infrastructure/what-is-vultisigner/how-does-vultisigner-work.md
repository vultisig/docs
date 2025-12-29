---
description: >-
  Technical overview of Vultiserver Fast Vault setup. 2-of-2 keygen ceremony,
  vault share storage in secure enclave, and signing flow.
---

# How does a Fast Vault work?

## Set Up

### 2-of-2 Fast Vault

In this setup, the Vault is created together with a user's device and the Vultisigner server.\
The user registers a keygen request with the Vultisigner server, which creates a new Vault that records the connection to the user's device. A keygen ceremony is then executed, which creates the Vault shares on the user's device and on the Vultisigner server. \\

## Vault Share storage

The Vultisigner's Vault Shares are stored on a dedicated Vultisigner server, which contains the Vault Shares along with configured instructions for signing transactions.\\

For security reasons, the Vultisigner's Vault share is password encrypted and the user must provide an email address to which the Vault share will be sent during setup, for backup and independent access.

<figure><img src="../../.gitbook/assets/Vultisigner storage 2-2.png" alt="" width="563"><figcaption></figcaption></figure>

## Transaction Signing

When a user wishes to sign a transaction, the request is sent to the Vultisigner server.\
The appropriate Vultisigner Vault will verify that the configured transaction policies are met before participating in the Keysign process.\
If they are met, the Vultisigner joins the Keysign ceremony and the transaction is broadcasted to the blockchain.

<figure><img src="../../.gitbook/assets/default vultisigner.png" alt=""><figcaption><p>Vultisigner Flowchart</p></figcaption></figure>

{% hint style="info" %}
In case the Vultisigner server is offline and unable to sign, the user needs to import the Vultisigner share received via email into another device and sign like a normal Keysign.
{% endhint %}
