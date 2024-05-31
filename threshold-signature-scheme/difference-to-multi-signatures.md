---
description: What is the difference between TSS and Multi Signatures?
---

# Difference to Multi-Signatures

The main difference between Threshold Signature Schemes (TSS) and Multi-Signatures (MS) is that the TSS signature is created using the combination of vault shares through Zero-Knowledge Proof, where no private key is constructed, while MS uses multiple private keys to create a signature.\


<div align="left">

<figure><picture><source srcset="../.gitbook/assets/Tx white.png" media="(prefers-color-scheme: dark)"><img src="../.gitbook/assets/TX black.png" alt="" width="375"></picture><figcaption><p>Threshold Signatures</p></figcaption></figure>

 

<figure><picture><source srcset="../.gitbook/assets/MS.png" media="(prefers-color-scheme: dark)"><img src="../.gitbook/assets/MS dark.png" alt="" width="375"></picture><figcaption><p>Multi-Sigantures</p></figcaption></figure>

</div>

Increasing the inconvenience of MS, that lost keys can't be re-shared and the funds ultimately need to be migrated due to lower security and less redundancy. \
Here is a comparison of different factors for TSS and MS:

***

<table><thead><tr><th>Factors</th><th width="270.3333333333333">Threshold Signature Scheme</th><th>Multi Signature</th></tr></thead><tbody><tr><td>Private Key Storage</td><td>Parties holding shares</td><td>Parties having multiple keys</td></tr><tr><td>On-Chain Footprint</td><td>One Single Signature signing a Transaction</td><td>Multiple Signatures signing one Transaction</td></tr><tr><td>Losing Shares/Keys</td><td>New participants can be added</td><td>Funds need to migrate</td></tr><tr><td>Compatibility</td><td>Multi-Chain:<br>ECDSA and EDDSA Support</td><td>Single Chain Support</td></tr><tr><td>Key Construction</td><td>Private Key never constructed</td><td>Parties having each one private key</td></tr><tr><td>Adding participants</td><td>✔</td><td>❌</td></tr></tbody></table>
