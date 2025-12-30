---
description: >-
  TSS vs Multi-Sig explained. Why threshold signatures are superior: no
  private key reconstruction, chain-agnostic, lower fees, better privacy.
---

# Difference to Multi-Signatures

The main difference between Threshold Signature Schemes (TSS) and Multi-Signatures (MS) lies in their approach to generating signatures. In TSS, the signature is created by combining vault shares using Zero-Knowledge Proofs—no single private key is ever constructed. In contrast, MS involves multiple private keys, each contributing to the final signature.

***

## Visualization: Threshold Signatures

<div align="center">
  <figure>
    <picture>
      <source srcset="../.gitbook/assets/Tx white.png" media="(prefers-color-scheme: dark)">
      <img src="../.gitbook/assets/TX black.png" alt="" width="375">
    </picture>
    <figcaption><strong>Threshold Signatures</strong></figcaption>
  </figure>
</div>

## Visualization: Multi-Signatures

<figure>
  <picture>
    <source srcset="../.gitbook/assets/MS.png" media="(prefers-color-scheme: dark)">
    <img src="../.gitbook/assets/MS dark.png" alt="" width="375">
  </picture>
  <figcaption><strong>Multi-Signatures</strong></figcaption>
</figure>

***

## Comparison

| Factor | Threshold Signature Scheme | Multi-Signature |
|--------|---------------------------|-----------------|
| **Private Key Storage** | No private key constructed; uses vault shares and ZKP | Multiple private keys held by parties |
| **On-Chain Footprint** | Single signature on-chain | Multiple signatures visible on-chain |
| **Redundancy** | High; lost shares can be re-shared | Low; lost keys require fund migration |
| **Compatibility** | Multi-chain: ECDSA and EdDSA | Single chain support |
| **Flexibility** | Adjustable thresholds, replaceable devices | Fixed once created |
| **Adding Participants** | Supported | Not supported |

***

## Key Advantages of TSS

### No Key Reconstruction

In TSS, the private key never exists in one place. Multi-sig requires each signer to hold a complete private key, creating multiple single points of failure.

### Chain Agnostic

TSS works identically across all blockchains that support standard signatures. Multi-sig implementations vary by chain and often aren't available.

### Lower Fees

TSS produces a single signature regardless of threshold. Multi-sig requires multiple on-chain signatures, increasing transaction fees.

### Better Privacy

TSS transactions appear identical to normal single-signature transactions. Multi-sig transactions expose the signing structure on-chain.

### Flexibility

TSS allows resharing to add/remove devices without changing addresses. Multi-sig typically requires fund migration for any configuration change.

***

## Conclusion

Flexibility and security advantages make TSS the superior approach for managing digital assets compared to traditional multi-signature schemes.

***

## Related

- [How DKLS23 Works](how-dkls23-works.md)
- [TSS Actions](tss-actions.md)
