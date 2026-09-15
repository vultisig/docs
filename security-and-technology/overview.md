---
description: >-
  How Vultisig security works. Threshold Signature Schemes (TSS), DKLS23
  protocol, keysigning process, and cryptographic foundations explained.
---

# Vultisig security and TSS

Vultisig uses threshold signatures (TSS), a form of multi-party computation (MPC). There is no complete private key in normal operation.

***

## Core Concepts

### No Private Key Ever Exists

Unlike traditional wallets, Vultisig never constructs a complete private key. Instead, cryptographic operations are performed across distributed vault shares using zero-knowledge proofs. Even during signing, the key remains split—only the signature is assembled.

### Threshold Security

Vultisig uses a `t`-of-`n` threshold model. For a 2-of-3 vault, any 2 devices can sign, but no single device can act alone. This provides both security (no single point of compromise) and redundancy (one device can be lost).

***

## TSS Protocols

Vultisig supports two TSS protocols:

| Protocol | Status | Signing Rounds | Speed |
|----------|--------|----------------|-------|
| [GG20](how-gg20-works.md) | Legacy | 6 rounds | Baseline |
| [DKLS23](how-dkls23-works.md) | Current | 3 rounds | 5-10x faster |

New vaults use DKLS23 by default. Existing GG20 vaults can be [upgraded](../app-guide/vault-management/vault-upgrade.md).

***

## Key Operations

| Operation | Description | Guide |
|-----------|-------------|-------|
| **Key Generation** | Creating vault shares across devices | [TSS Actions](tss-actions.md) |
| **Key Signing** | Threshold devices signing transactions | [Keysign](keysign.md) |
| **Re-sharing** | Adding/removing devices from a vault | [TSS Actions](tss-actions.md) |

***

## Security Comparisons

- [Vultisig vs multisig](difference-to-multi-sig.md)
- [Vultisig vs passkeys](difference-to-passkeys.md)

***

## Emergency Procedures

In the unlikely event that Vultisig software becomes unavailable, vault shares can be recombined to extract a traditional private key:

- [Emergency Recovery](emergency-recovery.md) — Last-resort key extraction

{% hint style="danger" %}
Emergency recovery permanently converts a TSS vault to a single-signature wallet. Only use if Vultisig software is completely unavailable.
{% endhint %}

***

## Technical Deep Dives

{% content-ref url="how-dkls23-works.md" %}
[how-dkls23-works.md](how-dkls23-works.md)
{% endcontent-ref %}

{% content-ref url="how-gg20-works.md" %}
[how-gg20-works.md](how-gg20-works.md)
{% endcontent-ref %}

{% content-ref url="tss-actions.md" %}
[tss-actions.md](tss-actions.md)
{% endcontent-ref %}

{% content-ref url="keysign.md" %}
[keysign.md](keysign.md)
{% endcontent-ref %}
