---
description: This is where you can export a vault share.
---

# Vault Backup

**How do Vault-Shares work in Vultisig?**

* Each share stores information that can be used to create the private key.
* &#x20;Each device in the vault has its own vault share, not one share for the whole vault.&#x20;
* Each vault share by itself does not have access to your assets, so it is safe to export and store them digitally.

{% hint style="danger" %}
Do not store vault shares of the same vault from different devices in the same location, as this means that a malicious party can reconstruct your vault.
{% endhint %}

**Digital Storage Suggetions**

* Store vault-shares on different cloud drives
* Have independent cloud drive accounts for each Vault share.
* Ensure that a device does not have access to multiple cloud drive accounts that store Vault shares.
* Store in separate password managers
* Use offline devices to store vault shares

{% hint style="warning" %}
Any backup is only as secure as the media and password used to store the vault share.
{% endhint %}

**Recovering a lost device**

Because each vault share is unique, the lost device can only be recovered with the corresponding vault share backup.

Example:\
Say we have 3 devices and 3 backups of each device. Then we lose device 3.\


<figure><img src="../../.gitbook/assets/reconstruct 1.jpg" alt="" width="375"><figcaption><p>3 devices, 3 shares, 1 device lost</p></figcaption></figure>

* If Device 3 is lost, it can only be recovered by using the backup share of Device 3 on a new device.
* If a share of device 1 or 2 is installed on the new device, the vault cannot be accessed because two of the same shares cannot sign the vault.
* Installing backup share 3 on device 1 or 2 will override that device as device 3.

<figure><img src="../../.gitbook/assets/reconstruct 2.jpg" alt="" width="375"><figcaption><p>Possible recreation</p></figcaption></figure>

{% hint style="info" %}
Two of the same vault shares can never access the same vault.
{% endhint %}
