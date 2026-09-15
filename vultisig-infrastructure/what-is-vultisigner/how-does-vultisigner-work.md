---
description: >-
  How a Vultisig Fast Vault is created and signed. 2-of-2 keygen with
  Vultiserver, where the server share is stored, and what to do if the server
  is offline.
---

# How a Vultisig Fast Vault works

## Set Up

### 2-of-2 Fast Vault

In this setup, the Vault is created together with a user's device and the Vultisigner server.\
The user registers a keygen request with the Vultisigner server, which creates a new Vault that records the connection to the user's device. A keygen ceremony is then executed, which creates the Vault shares on the user's device and on the Vultisigner server. \\

## Vault Share storage

The Vultiserver vault share is stored on a dedicated server and encrypted with the Fast Vault password. During setup, a backup of that share is emailed to you. Configurable transaction policies are [not in the apps yet](what-can-be-configured.md).\\

For security reasons, the Vultisigner's Vault share is password encrypted and the user must provide an email address to which the Vault share will be sent during setup, for backup and independent access.

<figure><img src="../../.gitbook/assets/Vultisigner storage 2-2.png" alt="" width="563"><figcaption></figcaption></figure>

## Transaction Signing

When you approve a transaction on your device, the request is sent to Vultiserver. The server joins the keysign with its share. You still have to approve on the device. Configurable policies are not applied today.

<figure><img src="../../.gitbook/assets/default vultisigner.png" alt=""><figcaption><p>Vultisigner Flowchart</p></figcaption></figure>

{% hint style="info" %}
In case the Vultisigner server is offline and unable to sign, the user needs to import the Vultisigner share received via email into another device and sign like a normal Keysign.
{% endhint %}
