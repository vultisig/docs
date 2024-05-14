---
description: How to manage a vault
---

# Managing your Vault

## Vault Functions

Select a vault. Then, click on the gear icon below the chains displayed, on the right.

### Details

This is where you see the details of your vault, including parties and public keys

### Backup

This is where you can export a vault share.

* For a m-of-n setup, there are n individual vault shares that can be backed-up. Assuming all your devices are lost, you would need to have at least m vault shares backed-up, which can be imported into fresh devices and then you can continue to sign, reshare, etc.
* Each vault share by itself has no access to your assets, thus it is save to export and store them digitally.
* However, if someone else has access to m vault shares, they WILL have access to your assets.

{% hint style="danger" %}
Do not store vault shares of the same vault from different devices in the same location as it means a malicous party can re-construct your vault.
{% endhint %}

* Digital storage suggestions:

\-- Store one vault share on your cloud drive, store another vault share on your partner's cloud drive.

\-- Or, have separate/independant cloud drive accounts for each vault share. Make sure one device do not have access to those multiple cloud drive accounts (thus access to multiple vault shares).

\-- Or, store the vault shares on separate Password Managers.

* Each of the vault shares are unique. For example:

\-- Let’s say the original 3 vault shares for a 2of3 vault are x, y & z; on devices a, b, c respectively.

\-- User imports vault share z into device d —> user cannot use only device c and d to sign (because they are the same vault share/signer).

\-- The signing devices can be a\&b, a\&c, a\&d, b\&c or b\&d.

### Rename

You can rename your vault. Note, it only renames the local vault name, and doesn't affect other devices.

### Reshare

Reshare is a function where you can kick out other devices in the vault, as well as upgrade or downgrade the number of parties.

<figure><img src="../../.gitbook/assets/ManageVault.png" alt="" width="375"><figcaption></figcaption></figure>
