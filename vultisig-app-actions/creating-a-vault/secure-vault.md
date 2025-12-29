---
description: >-
  Create a Secure Vault with multiple devices for cold storage security.
  Flexible threshold signing (2-of-2 up to m-of-n). Ideal for holdings and DAOs.
---

# Secure Vault

## Overview

This setup offers the highest level of security and is considered the 'cold wallet' equivalent in the Vultisig App.

**Secure Vaults consist solely of user devices.**

Users need at least two devices, although minimum three is recommended. More devices can be added to their vaults, increasing the signing threshold and enhancing security with each additional device.

<figure><img src="../../.gitbook/assets/image (25).png" alt=""><figcaption></figcaption></figure>

{% hint style="info" %}
This option is also ideal for shared wallets among multiple users and DAOs.
{% endhint %}

***

The vaults will have a `m`-of-`n` Threshold to sign transactions, where `m` is at least 2/3rds (Threshold is 67%) of `n`, and no maximum number of `n` devices.&#x20;

The following are the most common vault setups:

1. **2-of-3 vault** - three devices to create a vault and two devices can sign a transaction.\
   This vault is automatically redundant, which means you can lose one device and still have access to your vault. To make sure you are fully protected, please [back up](../managing-your-vault/vault-backup.md) the Vault shares of every device.\
   **This vault type is recommended as a secure vault setup.**
2. **3-of-4 vault** - four devices to create a vault and three devices to sign a transaction.\
   This vault is automatically redundant, which means you can lose one device and still have access to your vault. To make sure you are fully protected, please [back up](../managing-your-vault/vault-backup.md) the Vault shares of every device.
3. **2-of-2 vault** - two devices to create a vault and two devices to sign a transaction.\
   This vault is vunerable if you lose one device, you can lose access to the funds.\
   To make sure you are fully protected, please [back up](../managing-your-vault/vault-backup.md) the Vault shares of every device.\
   **This vault type is not recommended as a secure vault setup.**

{% hint style="info" %}
What is the most redundant vault that allows you maximum flexibility?

Try this:\
1\) Use 3 different builds (Mac, iOS, Android) on 3 different devices.

2\) Export vaults shares with 3 different passwords to encrypt.

3\) Save vault shares in 3 different Cloud Storage options (Google, iCloud, Proton, Dropbox etc), each with a unique email address per Cloud Storage.

4\) Ensure each email has 2FA.\
\
To compromise this vault, an attacker would need to\
\
1\) Break into 2 different emails, intercept your 2FA, AND crack 2 different passwords, or\
2\) Compromise 2 different devices (get past passcodes and biometrics).

If you practice good security, the likelihood of this is significantly low, almost zero. (How many times has someone broken into 1 of your devices/email/storage accounts, let alone 2).

The advantage of this setup is you can re-spawn anywhere in the world with just your email accounts and passwords, without carrying around hardware wallets and seed-phrases.
{% endhint %}

## Creating a Secure Vault

* Download Vultisig and get your devices ready. You will need:
  * One initiating device
  * Multiple pairing devices

<figure><img src="../../.gitbook/assets/image (26).png" alt="" width="375"><figcaption></figcaption></figure>

* Select Secure Vault option on one device, this will be the initiating device.

<figure><img src="../../.gitbook/assets/image (29).png" alt="" width="375"><figcaption></figcaption></figure>

* **Initiating device: Next** -> Select a vault name. -> it will show a QR Code to scan with your pairing device(s)

<figure><img src="../../.gitbook/assets/image (30).png" alt=""><figcaption></figcaption></figure>

* **Pairing Devices:** Select "Scan QR code" and scan the QR code from the main device

<figure><img src="../../.gitbook/assets/image (31).png" alt=""><figcaption></figcaption></figure>

* When all wanted pairing devices are present, create the vault with the `Next` button on the initiating device.

### Network Types

You can choose Internet or WiFI.

1. **Internet:** Using the Vultisig relay server. Encrypted packages are routed through the Vultisig relay server over the Internet. Each device can be on different networks / Internet providers.
2. **Local**: Using local Wi-Fi Network, however may not work on some Wi-Fi networks (since they may block mDNS packets).

The Vault Creation step may fail if the Internet/network connection is not stable.

<figure><img src="../../.gitbook/assets/Local Mode.png" alt="" width="281"><figcaption></figcaption></figure>

### Keygen/Vault creation

Once you click NEXT, the keygen process will begin. First it will create the pre-parameters (your vault shares and some other aspects, about 5 seconds), then it will create the ECDSA and EdDSA keys (another 5 seconds).\
\
Finally, it is done! Make sure that all devices show the done screen.

<figure><img src="../../.gitbook/assets/image (32).png" alt=""><figcaption></figcaption></figure>

### Backups

Follow the onboarding process to back up each device and its shares. To ensure constant access to your funds, keep backups of each device at all times.\
\
Read more about Backups and their storage [here](../managing-your-vault/vault-backup.md).

{% hint style="warning" %}
After creating a vault, **ALWAYS** [back up](../managing-your-vault/vault-backup.md) **every** device.
{% endhint %}

