---
description: How do Threshold Signatures work ?
---

# How it works

## Concept

The Threshold Signature scheme introduced in 2018 by Gerrano-Goldfeder uses a cryptographically secure zero-knowledge proof to distribute vault shares and sign transactions.\
Vultisig uses the improved version of TSS the GG20, which was implemented and thoroughly tested by Thorchain.\
It combines the Shamir's Secrect Sharing Scheme and Multi-party-computation (MCP) to securely be able to sign transactions or generate new Vault shares without reveiling anything but the correct Output.

***

## Shamir Secret Sharing

It is based on splitting one secrect into n-amount (total amount) of shares and enabling a subset of all shares, the m-amount (needed amount) of these shares to re-construct the secrect. This is done by utilizing the mathematical concept of polynomial interpolation. Therefore creating a secure and fault-tolerant option for sharing secrects.

<figure><picture><source srcset="../.gitbook/assets/TSS.png" media="(prefers-color-scheme: dark)"><img src="../.gitbook/assets/TSS dark.png" alt="" width="375"></picture><figcaption><p>Example of 3 shares constructing one secret</p></figcaption></figure>

With Shamir secret sharing scheme it is possible to construct a secrect (private key) with a m-amount of shares to sign transactions on the blockchain or having access to assets.

The flaw in Shamir Secrect Sharing lies in that once the secrect is constructed again it is vunerable to exploits or theft and therefore is not secure while using.&#x20;

***

## Multi Party Computation (MPC)

To compensate the flaw in Shamir's Secrect Sharing Scheme, a Multi Party Computation (MCP) is implemented into the Scheme. \
It allows to provide a solution forproof of access to the secret without ever constructing it.

This is archieved through the computation of a function with the Secrect shares of the participants. This function proofs the access to the secrect without the need to be ever created.&#x20;

That enhances the Secrect Sharing where initially a secrect got constructed, seperated and  distributed. Leaving the vunerability that the secrect was constructed once.&#x20;

<figure><picture><source srcset="../.gitbook/assets/MPC white.png" media="(prefers-color-scheme: dark)"><img src="../.gitbook/assets/MPC dark.png" alt="" width="375"></picture><figcaption><p>Scecret Sharing, Key Distribution</p></figcaption></figure>

***

Where with the combination of SSS and MPC essentially building the Threshold Signature Scheme, where the Vault Shares get created on each device and never leave it.&#x20;

Enabled through a key generation function with n-amount of shares, this function will output the same public key to all parties, and a different secret share for each participent. \
Therefore proving the access without revealing the secret shares.

As MPC is an offline computation, that gives it few advantages for the usage in blockchain:

* Enables the Threshold Signature Scheme as with MPC the m-amount can be freely set and later be newly configured through the re-share computation.&#x20;
* The on-chain footprint is similar to a Single Signature, keeping privacy and lowering transaction fees.
* Faster and more efficient signing computations

***

To sign a transaction the before distributed vault shares are used as an input in the MCP to generate a valid and verifiable Signature, which will be used on-chain.

<figure><picture><source srcset="../.gitbook/assets/Tx white.png" media="(prefers-color-scheme: dark)"><img src="../.gitbook/assets/TX black.png" alt="" width="375"></picture><figcaption><p>Signing Transaction</p></figcaption></figure>
