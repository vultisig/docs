---
description: >-
  Upgrade your Vultisig vault from GG20 to DKLS23 protocol. Faster signing,
  better security. Same addresses, no fund migration needed.
---

# Vault Upgrade

{% hint style="info" %}
Upgrades are only available for vaults created with the GG20 protocol.
{% endhint %}

***

## What is Upgrading?

Upgrading changes the underlying TSS protocol from GG20 to DKLS23. Benefits include:

- **Faster signing** — improved keygen and signing speed
- **Extension compatibility** — required for [Vultisig Extension](../../vultisig-ecosystem/vultisig-extension/)
- **Same addresses** — no need to migrate funds

Technical details are available in the [TSS section](../../security-technology/).

***

## How to Upgrade

The upgrade can be initiated from:
- The banner on the home screen
- **Settings** → **Vault Settings** → **Upgrade**

<div><figure><img src="../../.gitbook/assets/image (2).png" alt=""><figcaption><p>Settings option</p></figcaption></figure> <figure><img src="../../.gitbook/assets/Upgrade.png" alt=""><figcaption><p>Home banner</p></figcaption></figure></div>

### Requirements

- All original devices from vault creation must be present
- This is a full keygen ceremony (similar to creating a new vault)

### Steps

1. Initiate upgrade on one device
2. Join with all other devices
3. Complete the keygen process
4. **Backup all new vault shares immediately**

***

## Backup After Upgrade

{% hint style="warning" %}
New vault shares are created during upgrade. Store them with the same security precautions as original shares.
{% endhint %}

### Distinguishing Backups

| Protocol | File naming |
|----------|-------------|
| GG20 | `vault-xxxx-part1of2.vult` |
| DKLS23 | `vault-xxxx-share1of2.vult` |

***

## Upgrading Active Vaults

{% hint style="info" %}
Active Vaults are a discontinued feature with a different upgrade procedure.
{% endhint %}

To upgrade an Active Vault:

1. Import the server share onto a third device
2. Initiate upgrade from other devices
3. Join with all three devices
4. Complete upgrade

The result is a standard [Secure Vault](../creating-a-vault/secure-vault.md) without Vultiserver involvement.

{% hint style="info" %}
If Fast signing is still desired, create a new [Fast Vault](../creating-a-vault/fast-vault.md) after upgrading.
{% endhint %}

***

## Related

- [TSS Protocols](../../security-technology/) - Technical details
- [Vultisig Extension](../../vultisig-ecosystem/vultisig-extension/) - Requires DKLS23
