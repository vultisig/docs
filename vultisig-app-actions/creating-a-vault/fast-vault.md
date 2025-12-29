---
description: >-
  Create a Fast Vault in seconds—single-device hot wallet powered by
  Vultisigner. Perfect for daily transactions with seedless MPC security.
---

# Fast Vault

## Overview

In Vultisig this setup is the equivalent of a "hot wallet", allowing vaults to be created on-the-go with just one device, **without** requiring **multiple devices** from the user.

Fast Vaults are configured as a two-factor vault, where one device is held by the user and the other part is the Vultiserver, which automatically co-signs the user's requests (learn more about what the Vultiserver is [here](../../vultisig-infrastructure/what-is-vultisigner/)), making it a single signature experience.\
In the future, [transaction policies](../../vultisig-infrastructure/what-is-vultisigner/what-can-be-configured.md) will allow users to specify parameters for co-signing.

<figure><img src="../../.gitbook/assets/image (24).png" alt=""><figcaption></figcaption></figure>

{% hint style="info" %}
It is recommended not to store large amounts in these vaults and to use them as a daily wallet or 'hot wallet.'
{% endhint %}

***

## Fast Vault Creation

* Download Vultisig and get your device ready

<figure><img src="../../.gitbook/assets/image (26).png" alt="" width="375"><figcaption></figcaption></figure>

*   Select Fast Vault option

    <figure><img src="../../.gitbook/assets/image (28).png" alt="" width="375"><figcaption></figcaption></figure>

After selecting Fast Vault follow the onboarding flow with these simple steps:

### Steps

1. Choose a Vault Name
2. Enter your email to receive the Server Vault Share for complete self-custody! (In case the co-signing server is not available)
3. Set a password for the **server share**

{% hint style="warning" %}
Remember this password very well, as it encrypts your backup and is the only way to access it. Vultisig does not have access to user data and therefore cannot help with recovery.
{% endhint %}

4. Set an optional hint in case you temporarily forget the password of the server share.
5. Start the keygen (Vault creation) and let it finish
6. Enter the 4 digit verification code, received via email to verify that you received the server backup.
7. Backup your device share

{% hint style="danger" %}
Do not store the device backup on the device itself.
{% endhint %}

8. Start using your new Vultisig Fast Vault

***

## Backups

Even though Fast Vault uses one user device for signing transactions, it has **2 backup shares.** This is due to the co-signing server being the second part of the vault.&#x20;

{% hint style="info" %}
Always have both backups stored securely in separate locations. Read more about Backups [here](../managing-your-vault/vault-backup.md)
{% endhint %}

***

## Troubleshooting

* **Unable to connect to the server.**\
  The server may be temporarily unavailable, so establishing a connection is not possible. Please try again later.
* **Unstable Network**\
  A keygen may fail if you have an unstable network connection and the device drops the connection to the server.
  * Quit the app.
  * Change networks.
  * Start again.
* **Unable to sign with my Fast Vault**\
  This may has multiple reasons:

1. Q: Server not available. \
   A: Please try again later
2. Q: I directly get a QR code presented instead of a password.\
   A: You have imported the server share of the Fast Vault. To sign transactions with the password and only one device, please import the device share.
