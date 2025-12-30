---
description: >-
  View vault details: signing parties, ECDSA and EdDSA public keys, and vault
  configuration. Verify your multi-device setup at a glance.
---

# Vault Details

The Vault Details screen displays technical information about the current vault configuration.

<figure><img src="../../.gitbook/assets/Vault Details.png" alt="" width="188"><figcaption><p>Vault Details view</p></figcaption></figure>

***

## Information Displayed

| Field | Description |
|-------|-------------|
| **Vault Name** | Current name of the vault |
| **Vault Type** | TSS protocol used (GG20 or DKLS23) |
| **ECDSA** | Public key for ECDSA signatures (Bitcoin, Ethereum, etc.) |
| **EdDSA** | Public key for EdDSA signatures (Solana, Polkadot, etc.) |
| **M of N** | Signing threshold and total devices |

***

## Signing Parties

The M of N section shows all devices participating in the vault:
- **M** = Number of devices required to sign
- **N** = Total number of devices in the vault

Each device is listed with its identifier for verification.

***

## Related

- [Vault QR](vault-qr.md) - Export public keys
- [Vault Reshare](vault-reshare.md) - Modify device configuration
