---
description: What is the difference between TSS and Multi Signatures?
---

# Difference to Multi-Signatures

The main difference between Threshold Signature Schemes (TSS) and Multi-Signatures (MS) lies in their approach to generating signatures. In TSS, the signature is created by combining vault shares using Zero-Knowledge Proofs, ensuring that no single private key is ever constructed. In contrast, MS involves multiple private keys, each contributing to the creation of the final signature.

## Visualization Threshold Signatures

<div align="center">
  <figure>
    <picture>
      <source srcset="../.gitbook/assets/Tx white.png" media="(prefers-color-scheme: dark)">
      <img src="../.gitbook/assets/TX black.png" alt="" width="375">
    </picture>
    <figcaption><strong>Threshold Signatures</strong></figcaption>
  </figure>
</div>

## Visualization Multi-Signatures

<figure>
  <picture>
    <source srcset="../.gitbook/assets/MS.png" media="(prefers-color-scheme: dark)">
    <img src="../.gitbook/assets/MS dark.png" alt="" width="375">
  </picture>
  <figcaption><strong>Multi-Signatures</strong></figcaption>
</figure>

Increasing the inconvenience of Multi-Signatures (MS) is the fact that lost keys cannot be re-shared, and funds often need to be migrated to maintain security, leading to lower redundancy.

Here is a comparison of different factors for TSS and MS:

| Factors             | Threshold Signature Scheme                                                 | Multi-Signature                                               |
| ------------------- | -------------------------------------------------------------------------- | ------------------------------------------------------------- |
| Private Key Storage | No private key is constructed; uses vault shares and Zero-Knowledge Proofs | Parties having multiple private keys                          |
| On-Chain Footprint  | One Single Signature signing a Transaction on-chain                        | Multiple Signatures signing one Transaction; visible on-chain |
| Redundancy          | High redundancy; lost shares can be re-shared and reconfigured             | Low redundancy; lost keys require fund migration              |
| Compatibility       | Multi-Chain: ECDSA and EDDSA Support                                       | Single Chain Support                                          |
| Flexibility         | Can adjust signing thresholds and replace lost devices                     | Fixed once created; each party having one private key         |
| Adding participants | ✔                                                                         | ❌                                                            |

This comparison shows that flexibility and security prevail in TSS over MS, making it a better application for managing digital assets.
