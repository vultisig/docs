---
description: >-
  Vultisig security practices and audits. Independent assessments and
  responsible disclosure.
---

# Vultisig audits

Open-source threshold signatures, an audited signing library, and responsible disclosure.

There is no public bug bounty program today. Report a vulnerability through [GitHub](https://github.com/vultisig) or [Discord](https://discord.gg/Cugw9T2NrP).

***

## Security model

Vultisig uses threshold signature schemes (TSS). There is no complete private key in normal operation:

* Vault shares create signatures together
* Compromising one device does not move funds
* All app code is on [GitHub](https://github.com/vultisig)

For the protocol, see [Security & Technology](../security-and-technology/overview.md).

***

## Audits

### DKLS23 protocol library

The DKLS23 implementation Vultisig uses is from Silence Laboratories. Trail of Bits audited that **library**, not the Vultisig apps:

| Auditor | Date | Scope | Report |
| --- | --- | --- | --- |
| Trail of Bits | 2024 | DKLS23 protocol library | [View report](https://github.com/silence-laboratories/dkls23?tab=readme-ov-file#security-audit) |

### Application audits

Mobile and desktop application audit reports will be linked here when they are published.

***

## For users

1. Backup vault shares to separate, offline storage
2. Verify addresses before signing
3. Keep the apps updated
4. Use a Secure Vault for significant holdings
5. Never share vault shares or backup files

### What Vultisig cannot do

* Access your funds
* Recover lost vault shares
* Reverse blockchain transactions
* View your private keys (they never exist)

***

## Related

* [Security & Technology](../security-and-technology/overview.md)
* [Emergency Recovery](../security-and-technology/emergency-recovery.md)
* [Privacy Policy](privacy.md)
