---
description: Exporting vault shares, backing up or importing vaults shares.
---

# Vault Backup

<figure><img src="../../.gitbook/assets/image.png" alt=""><figcaption><p>Backup in Settings</p></figcaption></figure>

{% hint style="success" %}
**Vault shares should always be backed up.**&#x20;
{% endhint %}

## How do Vault-Shares work in Vultisig?

* Each device in the vault has its **own unique** vault share. So each device needs its **own backup**.
* Each vault share by itself does not have access to any assets, so it is safe to export and store them digitally.

{% hint style="danger" %}
**Never** store different vault shares of the same vault in the same location, as this could potentially lead to loss of funds. \
This can happen if a malicious party gains access to that location and is able to reconstruct the vault.
{% endhint %}

### Digital Storage Suggestions

* [x] Storing vault-shares on different cloud drives/locations
* [x] Having independent cloud drives or password manager for each Vault share.
* [x] Ensuring that a device does not have access to multiple cloud drives or password managers that store Vault shares.
* [x] Use offline devices to store vault shares

## Backup Encryption

It is possible to encrypt the Vault share with an optional password, adding an extra layer of security to the shares.

The Password option can either activated with putting in a strong password and **saving** it or can be **skipped**.

<figure><img src="../../.gitbook/assets/Enrypt Backup.png" alt="" width="188"><figcaption></figcaption></figure>

## **Recovering a lost device**

If a device is lost, there are two ways to recover the device:

* **Option 1:** Import the original backed-up vault share into a new device.
* **Option 2:** Reshare the vault using the two other devices to include a new device. This is only possible on a 2-of-3 setup.

{% hint style="warning" %}
**Backups are reshare sensitive**. \
This means that backups are incompatible with each other after using the reshare feature, see [here](vault-reshare.md#important-note-for-reshare) for more information.
{% endhint %}
