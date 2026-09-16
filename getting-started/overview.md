---
description: >-
  Download Vultisig, create a vault, back it up, and send. No seed phrase.
---

# Get started with Vultisig

You install the app, create a vault, back it up, then send. There is no seed phrase — each device holds a share of the key instead.

{% hint style="danger" %}
**Back up the vault before you put money in.**

If a device is lost and you have no backup, there is no recovery. Read [Backup & Recovery](backup-recovery.md) first.
{% endhint %}

***

## Start here

<table data-card-size="large" data-view="cards">
<thead>
<tr>
<th></th>
<th></th>
<th data-hidden data-card-target data-type="content-ref"></th>
</tr>
</thead>
<tbody>
<tr>
<td><strong>1. Download</strong></td>
<td>Install on the devices you will use to sign</td>
<td><a href="download-install.md">download-install.md</a></td>
</tr>
<tr>
<td><strong>2. Create a vault</strong></td>
<td>Fast Vault (one device) or Secure Vault (two or more)</td>
<td><a href="create-vault.md">create-vault.md</a></td>
</tr>
<tr>
<td><strong>3. Back it up</strong></td>
<td>Export a share from every device in the vault</td>
<td><a href="backup-recovery.md">backup-recovery.md</a></td>
</tr>
<tr>
<td><strong>4. First transaction</strong></td>
<td>Receive, then send a small test</td>
<td><a href="first-transaction.md">first-transaction.md</a></td>
</tr>
</tbody>
</table>

***

## Which vault?

If you have one device, start with a [Fast Vault](../app-guide/creating-a-vault/fast-vault.md). Your phone or laptop holds one share; VultiServer holds the other and signs with you. It can never sign alone. Setup takes about a minute.

If you want an extra device as a fallback, create a [Secure Vault](../app-guide/creating-a-vault/secure-vault.md). Two or more devices you control. 2-of-3 is the usual setup: you can lose one device and still sign.

A Fast Vault does not turn into a Secure Vault. If you want the higher threshold later, create a new Secure Vault and move the funds.

**Building with the SDK or CLI?**

- [SDK](../developer-docs/vultisig-sdk/) — TypeScript for apps, backends, and agents
- [CLI](../developer-docs/vultisig-sdk/CLI.md) — scripts and unattended agents

***

## Help

Stuck? [FAQ](../help/faq.md) or [Discord](https://discord.gg/Cugw9T2NrP).
