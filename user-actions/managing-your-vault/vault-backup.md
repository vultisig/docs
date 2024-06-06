---
description: Exporting vault shares, backing up vaults, importing vault shares.
---

# Vault Backup

{% hint style="success" %}
**Vault shares should always be backed up**
{% endhint %}

### **How do Vault-Shares work in Vultisig?**

* Each vault share by itself does not have access to your assets, so it is safe to export and store them digitally.
* Each device in the vault has its own unique vault share. \
  So each device needs its own backup.

{% hint style="danger" %}
Do not store vault shares of the same vault from different devices in the same location, as this means that a malicious party can potentially reconstruct your vault.
{% endhint %}

**Digital Storage Suggestions**

* Store vault-shares on different cloud drives
* Have independent cloud drive accounts for each Vault share.
* Ensure that a device does not have access to multiple cloud drive accounts that stores Vault shares.
* Use offline devices to store vault shares

### **Recovering a lost device**

If you lose a device you have two options:

**Option 1)** Import the original backed-up vault share into a new device.&#x20;

**Option 2)** Re-share the vault using the two other devices to include a new device.
