---
description: >-
  Reshare vault keys to add or remove devices. Upgrade security without
  changing addresses or moving funds. Advanced feature for vault management.
---

# Vault Reshare

{% hint style="danger" %}
**Advanced feature.** Access to funds is at risk if not properly managed.
{% endhint %}

<figure><img src="../../.gitbook/assets/image (1) (1) (1) (1) (1).png" alt=""><figcaption><p>Reshare in settings</p></figcaption></figure>

***

## What is Resharing?

Resharing modifies the number of devices in a vault without changing wallet addresses or moving funds. This feature enables:

- Adding new devices to increase security
- Removing devices that are no longer needed
- Including Vultiserver in an existing Secure Vault
- Replacing a lost device with a new one

<figure><img src="../../.gitbook/assets/Reshare.png" alt="" width="563"><figcaption></figcaption></figure>

***

## Requirements

{% hint style="info" %}
A threshold majority is **always** required. For a 2-of-3 vault, at least 2 devices must be present.
{% endhint %}

***

## Video Guides

**Resharing from 3-of-4 to 2-of-3:**

[![Reshare 3-of-4](../../.gitbook/assets/TwitterVideoThumbnail.jpeg)](https://x.com/iceman00008/status/1958446928271900905/video/1)

**Resharing from 2-of-2 to 2-of-3:**

[![Reshare 2-of-2](../../.gitbook/assets/TwitterVideoThumbnail.jpeg)](https://twitter.com/iceman00008/status/1825339005673857356/video/1)

***

## How to Reshare

1. Go to **Settings** → **Vault Settings** → **Reshare**
2. Choose whether to include Vultiserver
3. Start the reshare ceremony on the initiating device
4. Join with other devices (scan QR or use relay)
5. Add new devices or exclude devices as needed
6. Complete the keygen process

<figure><img src="../../.gitbook/assets/Reshare Vault - Start Screen.png" alt="" width="188"><figcaption></figcaption></figure>

***

## Use Cases

### Adding Devices

Join additional devices during the reshare ceremony to increase security:
- 2-of-2 → 2-of-3 (adds backup device)
- 2-of-3 → 3-of-4 (increases threshold)

### Removing Devices

Exclude a device by not joining it in the ceremony:
- 3-of-4 → 2-of-3 (reduces complexity)
- Remove a compromised device

### Replacing Lost Devices

If a device is lost but you have threshold access:
1. Initiate reshare with remaining devices
2. Add the replacement device
3. Complete reshare to issue new shares

***

## Critical Warnings

{% hint style="danger" %}
**Backup immediately after resharing!**

After resharing, all vault shares change. Old backups are NOT compatible with new shares.
{% endhint %}

{% hint style="info" %}
Old vault shares can still sign with each other. Resharing does not invalidate the old set—it creates a new parallel set.
{% endhint %}

***

## Related

- [Vault Backup](vault-backup.md) - Essential after resharing
- [Vault Upgrade](vault-upgrade.md) - Change TSS protocol
