---
description: >-
  Create a Vultisig vault. Fast Vault is one device plus VultiServer.
  Secure Vault is two or more devices you control. No seed phrase.
---

# Create your first Vultisig vault

Think of the vault as a safe that needs a threshold of keys to open. Your assets live in that safe. The full private key never exists in one place.

***

## Vault types

<table data-card-size="large" data-view="cards">
<thead>
<tr>
<th></th>
<th></th>
<th></th>
<th data-hidden data-card-target data-type="content-ref"></th>
</tr>
</thead>
<tbody>
<tr>
<td><strong>Fast Vault</strong></td>
<td>One device + VultiServer (2-of-2)</td>
<td>VultiServer signs with you and can never sign alone</td>
<td><a href="../app-guide/creating-a-vault/fast-vault.md">fast-vault.md</a></td>
</tr>
<tr>
<td><strong>Secure Vault</strong></td>
<td>Two or more devices you control</td>
<td>No VultiServer. An extra device in 2-of-3 is the fallback</td>
<td><a href="../app-guide/creating-a-vault/secure-vault.md">secure-vault.md</a></td>
</tr>
</tbody>
</table>

***

## Comparison

| | Fast Vault | Secure Vault |
|---------|------------|--------------|
| Setup | About a minute | About three minutes |
| Devices | 1 + VultiServer | 2 or more you control |
| VultiServer | Signs with you. Cannot sign alone. | Not used |
| If one device is lost | Your device backup + VultiServer | Extra device (2-of-3 or 3-of-4) |

If you only have one device, start with Fast Vault. If you want every share on hardware you hold, create a Secure Vault.

A Fast Vault does not turn into a Secure Vault. For the higher threshold later, create a **new** Secure Vault and move the funds.

**Secure Vault thresholds:**

- **2-of-3**: the usual setup. You can lose one device and still sign.
- **3-of-4**: three of four devices have to agree.
- **2-of-2**: both devices required. No fallback if one is gone.

***

## Video guides

**Creating a Fast Vault:**

[![Fast Vault Setup](../.gitbook/assets/TwitterVideoThumbnail.jpeg)](https://x.com/iceman00008/status/1955828412876312653/video/1)

**Creating a 2-of-2 Secure Vault:**

[![2-of-2 Setup](../.gitbook/assets/TwitterVideoThumbnail.jpeg)](https://x.com/iceman00008/status/1955865336341041197/video/1)

**Creating a 3-of-4 Secure Vault:**

[![3-of-4 Setup](../.gitbook/assets/TwitterVideoThumbnail.jpeg)](https://x.com/iceman00008/status/1958428338915287477/video/1)

***

## After creation

{% hint style="danger" %}
**Export a backup before you deposit.**

A lost device with no backup means the funds are gone.

[Backup & Recovery](backup-recovery.md)
{% endhint %}

***

## Step-by-step

- [Fast Vault](../app-guide/creating-a-vault/fast-vault.md)
- [Secure Vault](../app-guide/creating-a-vault/secure-vault.md)
