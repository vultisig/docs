---
description: >-
  Create a Secure Vault with multiple devices for cold storage security.
  Flexible threshold signing (2-of-2 up to m-of-n). Ideal for holdings and DAOs.
---

# Secure Vault

## Overview

Secure Vault offers the highest level of security—the "cold wallet" equivalent in Vultisig. These vaults consist solely of user-controlled devices with no third-party involvement.

A minimum of two devices is required; three or more is recommended for redundancy. Additional devices increase the signing threshold and enhance security.

<figure><img src="../../.gitbook/assets/image (25).png" alt=""><figcaption></figcaption></figure>

{% hint style="info" %}
Secure Vaults are ideal for shared wallets among multiple users and DAOs.
{% endhint %}

***

## Threshold Configurations

Secure Vaults use an `m`-of-`n` threshold where `m` is at least 67% of `n`. Common configurations:

| Setup | Devices | Signing | Redundancy | Recommendation |
|-------|---------|---------|------------|----------------|
| **2-of-3** | 3 | 2 | Yes (1 device) | **Recommended** |
| **3-of-4** | 4 | 3 | Yes (1 device) | Maximum security |
| **2-of-2** | 2 | 2 | No | Not recommended |

{% hint style="info" %}
**Best practice for maximum redundancy:**

1. Use 3 different platforms (Mac, iOS, Android)
2. Encrypt vault shares with 3 different passwords
3. Store shares in 3 different cloud providers (Google, iCloud, Dropbox)
4. Use unique email addresses with 2FA for each cloud account

This setup allows recovery from anywhere in the world with just email access and passwords—no hardware wallets or seed phrases required.
{% endhint %}

***

## Creation Steps

**Requirements:**
- One initiating device
- One or more pairing devices

<figure><img src="../../.gitbook/assets/image (26).png" alt="" width="375"><figcaption></figcaption></figure>

### On the Initiating Device

1. Select **Secure Vault** option

<figure><img src="../../.gitbook/assets/image (29).png" alt="" width="375"><figcaption></figcaption></figure>

2. Choose a vault name
3. A QR code will appear for pairing devices

<figure><img src="../../.gitbook/assets/image (30).png" alt=""><figcaption></figcaption></figure>

### On Pairing Devices

4. Select **Scan QR code** and scan the initiating device's QR

<figure><img src="../../.gitbook/assets/image (31).png" alt=""><figcaption></figcaption></figure>

### Complete Setup

5. When all devices are connected, tap **Next** on the initiating device
6. Wait for keygen to complete on all devices

<figure><img src="../../.gitbook/assets/image (32).png" alt=""><figcaption></figcaption></figure>

***

## Network Options

| Mode | Description | Use Case |
|------|-------------|----------|
| **Internet** | Uses Vultisig relay server | Devices on different networks |
| **Local** | Uses local WiFi (mDNS) | Same network, may not work on all WiFi |

<figure><img src="../../.gitbook/assets/Local Mode.png" alt="" width="281"><figcaption></figcaption></figure>

***

## After Creation

{% hint style="warning" %}
**Immediately backup every device** after vault creation. Without backups, device loss can mean permanent fund loss.

See [Backup & Recovery](../../getting-started/backup-recovery.md) for detailed instructions.
{% endhint %}
