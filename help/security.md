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

### TSS Library (mobile-tss-lib)
<!--
TODO:
Add Report Link
-->
The core TSS library has been audited by:

| Auditor | Date | Scope | Report |
|---------|------|-------|--------|
| Kudelski Security | 2024 | GG20 implementation | [View Report](https://github.com/vultisig/mobile-tss-lib/blob/main/audit/Kudelski-Vultisig-Report.pdf) |

### DKLS23 Implementation
<!--
TODO:
Add Report Link
-->
The upgraded DKLS23 protocol (via Silence Laboratories):

| Auditor | Date | Scope | Report |
|---------|------|-------|--------|
| Trail of Bits | 2024 | DKLS23 protocol | [View Report](https://github.com/silence-laboratories/dkls23/blob/main/audits/) |

### Application Audits

Mobile and desktop application security assessments are conducted regularly. Reports are published upon completion.

***

## Bug Bounty

Vultisig operates a responsible disclosure program for security researchers.

### Scope

- Vultisig mobile applications (iOS, Android)
- Vultisig desktop applications (macOS, Windows, Linux)
- Vultisig browser extension
- TSS library (mobile-tss-lib)
- Backend infrastructure

### Rewards

Bounties are determined based on severity:

| Severity | Description | Reward Range |
|----------|-------------|--------------|
| Critical | Direct fund theft, key extraction | Up to $50,000 |
| High | Significant security bypass | Up to $10,000 |
| Medium | Limited security impact | Up to $2,500 |
| Low | Minor issues | Up to $500 |

### Reporting

Report vulnerabilities to: **security@vultisig.com**

Include:
- Detailed description of the vulnerability
- Steps to reproduce
- Potential impact assessment
- Your suggested fix (optional)

Do not publicly disclose vulnerabilities before they are addressed.

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
