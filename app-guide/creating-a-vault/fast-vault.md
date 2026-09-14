---
description: >-
  Vultisig Fast Vault is a one-device seedless crypto wallet. Your phone plus
  Vultiserver as co-signer. Setup in about a minute. For daily amounts, not
  long-term storage.
---

# Vultisig Fast Vault

## Overview

A Vultisig Fast Vault is a one-device seedless crypto wallet. Your device holds one share; Vultiserver holds the other and co-signs. Same threshold math as a Secure Vault, one-device signing.

Learn more about Vultiserver in the [Infrastructure documentation](../../vultisig-infrastructure/what-is-vultisigner/).

<figure><img src="../../.gitbook/assets/image (24).png" alt=""><figcaption></figcaption></figure>

{% hint style="info" %}
Fast Vaults are recommended for daily transactions and smaller amounts. For larger holdings, use a [Secure Vault](secure-vault.md).
{% endhint %}

***

## Creation Steps

<figure><img src="../../.gitbook/assets/image (26).png" alt="" width="375"><figcaption></figcaption></figure>

1. Download Vultisig and open the app
2. Select **Fast Vault** option

<figure><img src="../../.gitbook/assets/image (28).png" alt="" width="375"><figcaption></figcaption></figure>

3. Choose a vault name
4. Enter your email to receive the server vault share
5. Set a password for the server share

{% hint style="warning" %}
Remember this password. It encrypts your backup and is the only way to access it. Vultisig cannot help with password recovery.
{% endhint %}

6. Set an optional hint for the password
7. Start keygen and wait for completion
8. Enter the 4-digit verification code received via email
9. Backup your device share

{% hint style="danger" %}
Do not store the device backup on the device itself. Use separate cloud storage or offline media.
{% endhint %}

10. Start using your Fast Vault

***

## Backups

Fast Vaults have **two backup shares** despite using one user device for signing:

- **Device share**: Your device's vault share
- **Server share**: Sent to your email during creation

{% hint style="info" %}
Both shares must be backed up in separate locations. See [Backup & Recovery](../../getting-started/backup-recovery.md) for storage recommendations.
{% endhint %}

***

## Troubleshooting

**Unable to connect to the server**
The server may be temporarily unavailable. Try again later.

**Unstable Network**
Keygen may fail with unstable connections. Quit the app, change networks, and restart.

**Unable to sign with Fast Vault**

1. *Server not available*: Try again later
2. *QR code appears instead of password prompt*: You imported the server share instead of the device share. Import the correct device share.
