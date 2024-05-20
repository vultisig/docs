---
description: What is the difference between TSS and Multi Signatures?
---

# Difference to Multi-Signatures

The main difference between Threshold Signature Schemes (TSS) and Multi-Signatures (MS) is that the TSS signature is created with a single private key, while MS uses multiple private keys to create a signature.\


<div align="left">

<figure><picture><source srcset="../.gitbook/assets/TSS.png" media="(prefers-color-scheme: dark)"><img src="../.gitbook/assets/TSS dark (1).png" alt="" width="375"></picture><figcaption><p>Threshold Signatures</p></figcaption></figure>

 

<figure><picture><source srcset="../.gitbook/assets/MS.png" media="(prefers-color-scheme: dark)"><img src="../.gitbook/assets/MS dark.png" alt="" width="375"></picture><figcaption><p>Multi-Sigantures</p></figcaption></figure>

</div>

Increasing the inconvenience that lost keys can't be re-shared and the funds that ultimately need to be migrated due to lower security and less redundancy. Here is a comparison of different factors for TSS and MS:

| Factors             | Threshold Signature Scheme                     | Multi Signature                         |
| ------------------- | ---------------------------------------------- | --------------------------------------- |
| Losing Shares/Keys  | New participants can be added                  | Funds need to migrate                   |
| On-Chain Footprint  | One Single Signature signing a Transaction     | More Signatures signing one Transaction |
| Adding participants | ✔                                              | ❌                                       |
| Compatibility       | <p>Multi-Chain:<br>ECDSA and EDDSA Support</p> | Single Chain Support                    |
| Private Key         | Never gets exposed                             | Potentionally exposed                   |
