---
description: What is the difference between TSS and Multi Signatures?
---

# Difference to Multi-Signatures

The main difference between Threshold Signature Schemes (TSS) and Multi-Signatures (MS) lies in their approach to generating signatures. In TSS, the signature is created by combining vault shares using Zero-Knowledge Proofs, ensuring that no single private key is ever constructed. In contrast, MS involves multiple private keys, each contributing to the creation of the final signature.

#### **Visualization** Threshold Signatures**:**

<div align="center">

<figure><picture><source srcset="../.gitbook/assets/Tx white.png" media="(prefers-color-scheme: dark)"><img src="../.gitbook/assets/TX black.png" alt="" width="375"></picture><figcaption><p><strong>Threshold Signatures</strong></p></figcaption></figure>

</div>

### **Visualization** Multi-Signatures**:**

<figure><picture><source srcset="../.gitbook/assets/MS.png" media="(prefers-color-scheme: dark)"><img src="../.gitbook/assets/MS dark.png" alt="" width="375"></picture><figcaption><p><strong>Multi-Sigantures</strong></p></figcaption></figure>

Increasing the inconvenience of Multi-Signatures (MS) is the fact that lost keys cannot be re-shared, and funds often need to be migrated to maintain security, leading to lower redundancy.

Here is a comparison of different factors for TSS and MS:

<table><thead><tr><th>Factors</th><th width="270.3333333333333">Threshold Signature Scheme</th><th>Multi Signature</th></tr></thead><tbody><tr><td>Private Key Storage</td><td>No private key is constructed; uses vault shares and Zero-Knowledge Proofs</td><td>Parties having multiple private keys</td></tr><tr><td>On-Chain Footprint</td><td>One Single Signature signing a Transaction on-chain</td><td>Multiple Signatures signing one Transaction; <br>visable on-chain</td></tr><tr><td>Redundancy</td><td>High redundancy; <br>lost shares can be re-shared and reconfigured</td><td>Low redundancy; lost keys require fund migration</td></tr><tr><td>Compatibility</td><td>Multi-Chain:<br>ECDSA and EDDSA Support</td><td>Single Chain Support</td></tr><tr><td>Flexibility</td><td>Can adjust signing thresholds and replace lost devices</td><td>Fixed once created; <br>each party having one private key</td></tr><tr><td>Adding participants</td><td>✔</td><td>❌</td></tr></tbody></table>

This comparison shows that flexibility and security prevail in TSS over MS, making it a better application for managing digital assets.
