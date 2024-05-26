---
description: Exporting vault shares, backing up vaults, importing vault shares.
---

# Vault Backup

**How do Vault-Shares work in Vultisig?**

* Each vault share by itself does not have access to your assets, so it is safe to export and store them digitally.
* Each device in the vault has its own unique vault share.&#x20;
* Combining more than two vault shares together could be used by a malicious party to re-create the vault and generate the vault secrets.&#x20;

{% hint style="danger" %}
Do not store vault shares of the same vault from different devices in the same location, as this means that a malicious party can reconstruct your vault.
{% endhint %}

**Digital Storage Suggestions**

* Store vault-shares on different cloud drives
* Have independent cloud drive accounts for each Vault share.
* Ensure that a device does not have access to multiple cloud drive accounts that stores Vault shares.
* Store in separate password managers
* Use offline devices to store vault shares

{% hint style="warning" %}
Any backup is only as secure as the media and password used to store the vault share.
{% endhint %}

**Recovering a lost device**

If you lose a device you have two options:

Option 1) Import the original backed-up vault share into a new device.&#x20;

Option 2) Re-share the vault using the two other devices to include a new device.

***

**Example for Recovery:**\
Say we have 3 devices and 3 backups of each device in our set-up.  Due to a boating accident we lose device 3.\


<figure><img src="../../.gitbook/assets/reconstruct 1.jpg" alt="" width="375"><figcaption><p>3 devices, 3 shares, 1 device lost</p></figcaption></figure>

Recover Do's:

* If device 3 is lost, we can import backup of vault share 3 into the new device 4.

Recover Don'ts:

* Installing a backup of Share 1 or Share 2 on a new device allows vault access, but it doesn't recover the third device. The vault remains a 2/2 setup, with two devices holding the same share.
* Installing backup share 3 on device 1 or 2 will override that respective device as device 3.

<figure><img src="../../.gitbook/assets/reconstruct 2.jpg" alt="" width="375"><figcaption><p>Possible recreation</p></figcaption></figure>
