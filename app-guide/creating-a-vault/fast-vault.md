---
description: >-
  Vultisig Fast Vault: one device plus VultiServer (2-of-2). Setup in about
  a minute. VultiServer co-signs and can never sign alone.
---

# Vultisig Fast Vault

## Overview

Your device holds one share. VultiServer holds the other and signs with you. It can never sign alone. Same math as a Secure Vault, one-device signing.

How the server share works: [VultiServer](../../vultisig-infrastructure/what-is-vultisigner/).

<figure><img src="../../.gitbook/assets/image (24).png" alt=""><figcaption></figcaption></figure>

{% hint style="info" %}
Want every share on devices you hold? Create a [Secure Vault](secure-vault.md) and move the funds. Fast Vault does not convert in place.
{% endhint %}

***

## Creation Steps

<figure><img src="../../.gitbook/assets/image (26).png" alt="" width="375"><figcaption></figcaption></figure>

1. Download Vultisig and open the app
2. Select **Fast Vault**

<figure><img src="../../.gitbook/assets/image (28).png" alt="" width="375"><figcaption></figcaption></figure>

3. Choose a vault name
4. Enter your email — the server share is sent there
5. Set a password for the server share

{% hint style="warning" %}
Remember this password. It encrypts the server backup. Vultisig cannot recover it for you.
{% endhint %}

6. Optionally set a hint for that password
7. Start keygen and wait until it finishes
8. Enter the 4-digit code from the email
9. Export your device share

{% hint style="danger" %}
Do not keep the device backup on the same phone. Use a different cloud account or offline media.
{% endhint %}

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
