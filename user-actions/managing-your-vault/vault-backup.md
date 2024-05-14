---
description: This is where you can export a vault share.
---

# Vault Backup

**How do Vault-Shares work in Vultisig?**

* Each share has information stored that can construct the private key
* Each device in the vault has it´s unique vault-share, not one share for the whole vault.
* Each vault share by itself has no access to your assets, thus it is save to export and store them digitally.

{% hint style="danger" %}
Do not store vault shares of the same vault from different devices in the same location as it means a malicous party can re-construct your vault.
{% endhint %}

***

**Digital Storage Suggetions**

* Store vault-shares on different cloud drives
* Have independant cloud drive accounts for each vault share.&#x20;
* Make sure one device do not have access to multiple cloud drive accounts where vault shares are stored.
* Store in seperate Password Managers
* Use offline devices to store vault-share

{% hint style="warning" %}
Each backup is only as safe as the medium and password used where the vault-share is stored
{% endhint %}

***

**Recovering a lost device**

As each vault share is unique the lost device can only be recovered with the respective backup of the vault-share. \


Example:\
Lets say we have 3 devices and 3 backups of each device. Then we loose device 3.

<figure><img src="../../.gitbook/assets/reconstruct 1 (1).jpg" alt="" width="375"><figcaption><p>3 devices, 3 shares, 1 device lost</p></figcaption></figure>

* If device 3 is lost, it can only be recovered with using the backup-share of device 3 on a new device.
* If installing a share of device 1 or 2 on the new device the vault can´t be accessed, as two of the same shares can´t sign the vault.
* If installing backup-share 3 on device 1 or 2 the respective device gets overridden as device 3.

<figure><img src="../../.gitbook/assets/reconstruct 2 (1).jpg" alt="" width="375"><figcaption><p>Possible recreations</p></figcaption></figure>

{% hint style="info" %}
Two of the same vault-shares can never access a vault
{% endhint %}
