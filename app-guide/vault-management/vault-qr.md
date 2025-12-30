---
description: >-
  Vault QR code contains root public keys and vault ID. Use for pairing
  and verification—not a backup. Share safely for view-only access.
---

# Vault QR

The Vault QR is a QR code containing the vault's public information for safe sharing and verification.

<figure><img src="../../.gitbook/assets/VultisigQR-Main Vault-828.png" alt="" width="188"><figcaption><p>Example Vault QR</p></figcaption></figure>

***

## What's Included

| Data | Description |
|------|-------------|
| **ECDSA Public Key** | Root key for ECDSA chains |
| **EdDSA Public Key** | Root key for EdDSA chains |
| **Vault UID** | Unique identifier |
| **Vault Name** | Display name |

***

{% hint style="warning" %}
**The Vault QR is NOT a backup.** It contains only public information and cannot be used to sign transactions or recover funds.
{% endhint %}

***

## How to Export

1. Open the vault home screen
2. Tap the QR icon in the top right corner
3. Save or share the QR code

<figure><img src="../../.gitbook/assets/Simulator Screenshot - iPhone 15 Pro - 2024-10-15 at 20.21.18.png" alt="" width="188"><figcaption><p>QR export location</p></figcaption></figure>

***

## Use Cases

### Airdrop Registration

The Vault QR can be imported into the [Vultisig Airdrop page](https://airdrop.vultisig.com/) to register for the [$VULT airdrop](../../vultisig-token/airdrop/).

### View-Only Access

Share the QR to allow others to view your vault's public addresses without granting signing access.

### Verification

Compare Vault QRs across devices to verify they belong to the same vault.

***

## Related

- [Vault Details](vault-details.md) - View public keys directly
- [Airdrop](../../vultisig-token/airdrop/) - VULT token distribution
