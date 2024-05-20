---
description: What is the difference between TSS and Multi Signatures?
---

# Difference to Multi-Signatures

The most significant difference between Threshold-Signatures-Schemes (TSS) and Multi-Signatures (MS) are that the TSS-Signature is constructed with one private key while MS uses multiple Private keys to generate a signature.\


<div align="left">

<figure><picture><source srcset="../.gitbook/assets/TSS.png" media="(prefers-color-scheme: dark)"><img src="../.gitbook/assets/TSS dark (1).png" alt="" width="375"></picture><figcaption><p>Threshold Signatures</p></figcaption></figure>

 

<figure><picture><source srcset="../.gitbook/assets/MS.png" media="(prefers-color-scheme: dark)"><img src="../.gitbook/assets/MS dark.png" alt="" width="375"></picture><figcaption><p>Multi-Sigantures</p></figcaption></figure>

</div>

Raising the inconvenience that lost keys can't be reshared and the funds ultimatley need to migrate due to lower security and less redundancy.&#x20;

Here is a comparison with different Factors for TSS and MS:

| Factors             | Threshold Signature Scheme                     | Multi Signature                         |
| ------------------- | ---------------------------------------------- | --------------------------------------- |
| Losing Shares/Keys  | New participants can be added                  | Funds need to migrate                   |
| On-Chain Footprint  | One Single Signature signing a Transaction     | More Signatures signing one Transaction |
| Adding participants | ✔                                              | ❌                                       |
| Compatability       | <p>Multi-Chain:<br>ECDSA and EDDSA Support</p> | Single Chain Support                    |
