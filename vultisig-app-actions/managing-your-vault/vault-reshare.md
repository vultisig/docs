---
description: >-
  Resharing Vault Shares with paired devices and adding new devices to the
  vault.
---

# Vault Reshare

{% hint style="danger" %}
T**his feature is only intended for advanced users! Access to funds are at risk if not properly managed!**
{% endhint %}

<figure><img src="../../.gitbook/assets/image (1) (1) (1) (1) (1).png" alt=""><figcaption><p>Feature in settings</p></figcaption></figure>

## What is the Reshare feature

Resharing is a feature that can be used to increase or decrease the number of participating devices in a Vault setup. In this setup also a Vultisigner can be included in the Vault setup without moving funds.

<figure><img src="../../.gitbook/assets/Reshare.png" alt="" width="563"><figcaption></figcaption></figure>

## How to use the Feature

{% hint style="info" %}
A threshold majority is **always** required to use this feature.\
For example, for a 2-of-2 vault, both devices must be present.
{% endhint %}

Do this if a device has been lost in a secure setup and you need to add a new, unique device or update the current shares.\
&#xNAN;_&#x49;f a device of a 2-of-2 setup is lost please use the_ [_backup feature_](vault-backup.md)_._

_Click on the image below to watch an explanation video on Twitter for Resharing from 3-of-4 to 2-of-3 (new version workflow)_

[![](../../.gitbook/assets/TwitterVideoThumbnail.jpeg)](https://x.com/iceman00008/status/1958446928271900905/video/1)

_Click on the image below to watch an explanation video on Twitter for Resharing from 2-of-2 to 2-of-3 (old version workflow)_

[![](../../.gitbook/assets/TwitterVideoThumbnail.jpeg)](https://twitter.com/iceman00008/status/1825339005673857356/video/1)

After clicking "Reshare" in the Vultisig Vault menu, the next steps are very similar to creating a Vault, as the reshare feature is just another keygen ceremony.

Start the Reshare with or without the Vultisigner present.

<figure><img src="../../.gitbook/assets/Reshare Vault - Start Screen.png" alt="" width="188"><figcaption></figcaption></figure>

You can increase a vault setup by simply adding a new device, or decrease a vault setup by excluding (not joining with) a device from the keygen ceremony.

{% hint style="info" %}
You always need the threshold of the current vault setup to change the configuration. Similar to when signing a transaction.
{% endhint %}

When resharing with Vultisigner, the steps are similar, but with the inclusion of the Vultisigner.\
The server checks if a Vault Share is currently present and requests the password.\
If a Vault Share is not present, a new share will be created on the server.

## When to use the Reshare feature

Like described in the section about [what the feature is](vault-reshare.md#what-is-the-reshare-feature), it can be used for the following:

### Changing the Number of Parties

Changing the number of parties is as simple as joining another keygen session.\
This gives the ability to change the number of devices participating in a setup.\
New unique devices can be added or old devices can be excluded.

## **IMPORTANT NOTE FOR RESHARE**

After resharing a vault, the vault shares of each device will also change. The old vault shares are not compatible with the new vault shares. Any device that was not present in a reshare will still have the old vault share, and will not be able to keysign with the devices with the new vault shares.

{% hint style="danger" %}
**Be sure to ALWAYS back up your new vault shares after using the reshare feature!**
{% endhint %}

{% hint style="info" %}
The set of old vault shares can still keysign with each other and access the funds linked to the vault address. Resharing does not "invalidates" the set of old vault shares.
{% endhint %}
