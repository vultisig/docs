---
description: >-
  Vultisig security practices, audits, and bug bounty program.
  Independent security assessments and responsible disclosure.
---

# Security

Vultisig prioritizes security through open-source transparency, professional audits, and responsible disclosure practices.

***

## Security Model

Vultisig's security is built on threshold signature schemes (TSS), eliminating single points of failure:

- **No single private key**: Keys are never constructed; vault shares create signatures collaboratively
- **Distributed trust**: Compromising one device does not compromise funds
- **Open source**: All code publicly auditable on [GitHub](https://github.com/vultisig)

For technical details, see [Security & Technology](../security-technology/README.md).

***

## Audits

### DKLS23 Implementation
<!--
TODO:
Add Report Link
-->
The upgraded DKLS23 protocol (via Silence Laboratories):

| Auditor | Date | Scope | Report |
|---------|------|-------|--------|
| Trail of Bits | 2024 | DKLS23 protocol | [View Report](https://github.com/silence-laboratories/dkls23?tab=readme-ov-file#security-audit) |

### Application Audits

Mobile and desktop application security assessments are conducted regularly. Reports are published upon completion.

***

## Security Best Practices

### For Users

1. **Backup vault shares** to secure, offline storage
2. **Verify addresses** before signing transactions
3. **Keep apps updated** for latest security patches
4. **Use Secure Vault** for significant holdings
5. **Never share vault shares** or backup files

### What Vultisig Cannot Do

- Access your funds
- Recover lost vault shares
- Reverse blockchain transactions
- View your private keys (they never exist)

***

## Related

- [Security & Technology](../security-technology/README.md) — Technical documentation
- [Emergency Recovery](../security-technology/emergency-recovery.md) — Fund recovery if software unavailable
- [Privacy Policy](privacy.md) — Data handling practices
