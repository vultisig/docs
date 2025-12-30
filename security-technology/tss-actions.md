---
description: >-
  TSS operations explained: Key Generation (vault creation), Key Signing
  (transactions), and Resharing (device changes). Core MPC actions.
---

# TSS Actions

Vultisig's Threshold Signature Scheme supports three core operations:

1. Key Generation
2. Key Signing
3. Re-Sharing

***

## Key Generation

Key generation (keygen) is the process by which devices prove access to a secret and generate a shared public key. This requires 100% of devices to be online.

The shared public key creates on-chain addresses (the "vault") for receiving funds. These addresses resemble simple wallet addresses (externally-owned accounts) and do not appear as special contracts or scripts.

{% hint style="success" %}
**Important:** Individual vault shares do not contain funds. The actual private key never exists in normal operation. Vault shares can be safely stored, emailed, or uploaded—they reveal nothing about the vault's location or contents.
{% endhint %}

{% hint style="danger" %}
**Never store multiple vault shares together.** A malicious party with enough shares could recombine them to access the vault:

1. Do not back up more than one vault share to the same device, email, or cloud storage
2. Do not upload more than one vault share to the same website

Following these guidelines prevents unauthorized vault access.
{% endhint %}

***

## Key Signing

Key signing is the process where a threshold of parties (67% minimum) collaborate to prove access to a secret and generate a signed transaction.

For a 2-of-3 vault, only 2 of 3 devices need to participate. This ensures the signing process remains secure and efficient while maintaining vault integrity.

The signing process:
1. Initiating device creates transaction payload
2. Threshold devices join via QR code or relay
3. Devices jointly compute signature shares
4. Signature is assembled and transaction broadcasts

See [Keysign](keysign.md) for detailed signing procedures.

***

## Re-share

Re-sharing allows the vault configuration to adapt by replacing devices or modifying the threshold.

Use cases:
- **Adding devices**: Upgrade from 2-of-2 to 2-of-3 or 3-of-4
- **Removing devices**: Exclude a lost or compromised device
- **Replacing devices**: Substitute a non-responsive device with a new one

{% hint style="info" %}
Re-sharing requires the current signing threshold. For a 2-of-3 vault, at least 2 devices must participate in the reshare ceremony.
{% endhint %}

{% hint style="warning" %}
After resharing, all vault shares change. Old backups are NOT compatible with new shares. Always backup immediately after resharing.
{% endhint %}

See [Vault Reshare](../app-guide/vault-management/vault-reshare.md) for step-by-step instructions.

***

## Related

- [Keysign](keysign.md) — Detailed signing process
- [Vault Reshare](../app-guide/vault-management/vault-reshare.md) — How to reshare
- [How DKLS23 Works](how-dkls23-works.md) — Technical protocol details
