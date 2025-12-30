---
description: >-
  MiCAR regulatory compliance white paper. EU Markets in Crypto-Assets
  Regulation compliance documentation for Vultisig.
---

# MiCAR White Paper

This document addresses Vultisig's compliance considerations under the European Union's Markets in Crypto-Assets Regulation (MiCAR).

***

## Overview

The Markets in Crypto-Assets Regulation (MiCAR) establishes a comprehensive regulatory framework for crypto-assets in the European Union. This white paper outlines how Vultisig's architecture and operations relate to MiCAR requirements.

***

## Vultisig Classification

### Non-Custodial Software

Vultisig is **non-custodial wallet software**—it provides tools for users to manage their own digital assets without Vultisig taking custody or control.

Key characteristics:

- **No custody**: Vultisig never holds user funds
- **No private key access**: Keys are never constructed; vault shares remain with users
- **User sovereignty**: Users maintain complete control over their assets

### Software Provider vs. CASP

Under MiCAR, Crypto-Asset Service Providers (CASPs) require authorization. Vultisig's role as a non-custodial software provider differs from regulated custody services:

| Aspect | CASP (Custodial) | Vultisig (Non-Custodial) |
|--------|------------------|--------------------------|
| Asset Control | Provider holds assets | User holds assets |
| Key Management | Provider manages keys | User manages vault shares |
| Recovery | Provider can recover | Only user can recover |
| Regulatory Status | Requires CASP license | Software tool |

***

## Technical Architecture

### Threshold Signatures

Vultisig uses threshold signature schemes (TSS) where:

1. No single private key is ever created
2. Vault shares are distributed across user devices
3. Signatures require threshold collaboration
4. Vultisig has no access to signing capability

### Data Handling

Vultisig minimizes data collection:

- No user registration required
- No KYC/identity verification
- No transaction data storage
- Vault data stored locally on user devices

***

## User Protections

### Transparency

- **Open source**: All code publicly auditable
- **Documentation**: Comprehensive user guides
- **Risk disclosure**: Clear communication of self-custody responsibilities

### Security

- **Audited code**: Independent security assessments
- **No single point of failure**: TSS architecture
- **User education**: Best practices documentation

***

## Regulatory Engagement

Vultisig is committed to:

- Monitoring regulatory developments
- Engaging with regulators constructively
- Adapting to evolving requirements
- Maintaining transparency with users

***

## Disclaimers

This white paper is for informational purposes only and does not constitute legal advice. Regulatory interpretations may vary by jurisdiction and may change over time.

Users are responsible for:

- Understanding applicable regulations in their jurisdiction
- Complying with local laws regarding digital assets
- Seeking professional legal and tax advice as needed

***

## Contact

For regulatory inquiries:

- **Email**: legal@vultisig.com

***

## Related

- [Terms of Use](terms.md)
- [Privacy Policy](privacy.md)
- [Security](security.md)
