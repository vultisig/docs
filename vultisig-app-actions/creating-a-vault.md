---
description: How to create a vault.
---

# Creating a Vault

## Overview

Vultisig vaults are multi-factor by nature, meaning that at least 2 devices are required to create a vault.

There are different types of vaults that can be created in the Vultisig App, designed to meet all users' needs and configuration preferences. From a 'hot wallet' option to a 'cold wallet' option, users have full flexibility in creation and modification afterward.

Creating a Vault is also called a Key Generation (KeyGen) event, see more information [here](../threshold-signature-scheme/tss-actions.md#key-generation).

***

## Setup - Vault Types

{% hint style="success" %}
iOS, MacOS, Android, Windows and Linux are supported. Check the [website](https://vultisig.com/) for the latest update on distribution.
{% endhint %}

There are two general setup types to choose from:

* Fast Vaults
* Secure Vaults

### Fast Vaults

This setup is the equivalent of a "hot wallet" in Vultisig, allowing vaults to be created on-the-go **without** requiring **multiple devices** from the user.

<figure><img src="../.gitbook/assets/Setup Fast (1).png" alt="" width="375"><figcaption><p>Fast Vaults in app</p></figcaption></figure>

Fast Vaults are configured as a two-factor vault, where one device is held by the user and the other part is the Vultiserver, which automatically co-signs the user's requests (learn more about what the Vultiserver is [here](../vultisig-infrastructure/what-is-vultisigner/)), making it a single signature experience.\
In the future, [transaction policies](../vultisig-infrastructure/what-is-vultisigner/what-can-be-configured.md) will allow users to specify parameters for co-signing.

{% hint style="info" %}
It is recommended not to store large amounts in these vaults and to use them as a daily wallet or 'hot wallet.'
{% endhint %}

***

### Secure Vault

This setup offers the highest level of security and is considered the 'cold wallet' equivalent in the Vultisig App.

<figure><img src="../.gitbook/assets/Setup Secure (1).png" alt="" width="375"><figcaption><p>Secure Vault in app</p></figcaption></figure>

**Secure Vaults consist solely of user devices.**

Users need at least two devices, although minimum three is recommended. More devices can be added to their vaults, increasing the signing threshold and enhancing security with each additional device.

{% hint style="info" %}
This option is also ideal for shared wallets among multiple users and DAOs.
{% endhint %}

The vaults will have a `m`-of-`n` Threshold to sign transactions, where `m` is at least 2/3rds (Threshold is 67%) of `n`, and no maximum number of `n` devices.&#x20;

{% hint style="info" %}
The more devices you use, the longer it will take to process any transactions.
{% endhint %}

The following are the most common vaults:

1. **2-of-3 vault** - three devices to create a vault and two devices can sign a transaction.\
   This vault is automatically redundant, which means you can lose one device and still have access to your vault. To make sure you are fully protected, please [back up](managing-your-vault/vault-backup.md) the Vault shares of every device.\
   **This vault type is recommended as a secure vault setup.**
2. **3-of-4 vault** - four devices to create a vault and three devices to sign a transaction.\
   This vault is automatically redundant, which means you can lose one device and still have access to your vault. To make sure you are fully protected, please [back up](managing-your-vault/vault-backup.md) the Vault shares of every device.
3. **2-of-2 vault** - two devices to create a vault and two devices to sign a transaction.\
   This vault is vunerable if you lose one device, you can lose access to the funds.\
   To make sure you are fully protected, please [back up](managing-your-vault/vault-backup.md) the Vault shares of every device.\
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

## Generating A Vault

<figure><img src="../.gitbook/assets/How Keygen works.png" alt=""><figcaption><p>Keygen Flowchart</p></figcaption></figure>

***

### How to generate a Vault

#### Videos

_Click on the image below to watch an explanation video on Twitter on creating a 2/2 Secure Vault_

[![](../.gitbook/assets/TwitterVideoThumbnail.jpeg)](https://twitter.com/iceman00008/status/1824683328085299628/video/1)

_Click on the image below to watch an explanation video on Twitter on creating a 2/2 Fast Vault_

[![](../.gitbook/assets/TwitterVideoThumbnail.jpeg)](https://twitter.com/iceman00008/status/1843118661076455906/video/1)

***

#### Step by Step Guide

Get your devices ready and create a vault.

<figure><img src="../.gitbook/assets/Create Vault Start (1).png" alt="" width="375"><figcaption><p>Create Vault</p></figcaption></figure>

Select the Vault setup of preference:

* Fast Vault
* Secure Vault

<figure><img src="../.gitbook/assets/Setup Secure (1).png" alt="" width="375"><figcaption></figcaption></figure>

After selecting your preferred setup, follow the steps for the following Setups:

#### Fast Vault

* Select "Fast Vault" in the Setup selection
* Choose a Vault Name
* Enter your email to receive the Server Vault Share for complete self-custody!
* Set a password for the **server vault share**
* Set an optional hint in case you temporarily forget the server share password.
* Start the keygen (Vault creation) and let it finish
* Enter the 4 digit verification code, received via email
* Backup your device share
* Use your new Vultisig Vault

#### Secure Vault

Start Vault creation with main device and join with all pairing devices.

**Main Device:** START -> will show a QR Code to scan with your pairing device(s)

<figure><img src="../.gitbook/assets/1 Device Scanned.png" alt="" width="281"><figcaption></figcaption></figure>

**Pairing Devices:** Select "Scan QR code" and scan the QR code from the main device

<figure><img src="../.gitbook/assets/QR Code Sheet.png" alt=""><figcaption></figcaption></figure>

When all wanted pairing devices are present, create the vault with the `Next` button.

### Network Types

You can choose Internet or WiFI.

1. **Internet:** Using the Vultisig relay server. Encrypted packages are routed through the Vultisig relay server over the Internet. Each device can be on different networks / Internet providers.
2. **Local**: Using local Wi-Fi Network, however may not work on some Wi-Fi networks (since they may block mDNS packets).

The Vault Creation step may fail if the Internet/network connection is not stable.

<figure><img src="../.gitbook/assets/Local Mode.png" alt="" width="281"><figcaption></figcaption></figure>

### Keygen

Once you click NEXT, the keygen process will begin. First it will create the pre-parameters (your vault shares and some other aspects, about 10 seconds), then it will create the ECDSA and EdDSA keys (another 10 seconds).\
Finally, it is done! Make sure that all devices show the done screen.

<div align="left"><figure><img src="../.gitbook/assets/Keygen (2).png" alt=""><figcaption><p>Ongoing Keygen</p></figcaption></figure> <figure><img src="../.gitbook/assets/Creation successful.png" alt="" width="375"><figcaption><p>Creation sucessful</p></figcaption></figure></div>

{% hint style="warning" %}
After creating a vault, **ALWAYS** [back up](managing-your-vault/vault-backup.md) every device.
{% endhint %}

### Troubleshooting

If a Keygen fails, it may be because you have an unreliable network and the devices dropped connections.

1. Quit the apps.
2. Change networks.
3. Start again.

Another reason are low spec android devices.\
Please ensure that your device has at **least 4 GB RAM**
