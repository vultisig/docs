---
description: How do Threshold Signatures work ?
---

# How it works

## Concept

The Threshold Signature Scheme (TSS) introduced in 2018 by Gerrano-Goldfeder uses a cryptographically secure zero-knowledge proof to distribute vault shares and sign transactions.\
Vultisig uses the improved version of TSS (GG20), which was implemented and thoroughly tested by THORChain.\
It combines the Shamir's Secret Sharing Scheme and Multi Party Computation (MPC) to securely be able to sign transactions or generate new Vault shares without revealing anything but the correct output.

The Threshold Signature scheme, introduced by Gerrano-Goldfeder in 2018, uses a cryptographically secure zero-knowledge proof to distribute vault shares and sign transactions. Vultisig uses the improved version of TSS, the GG20, which has been implemented and extensively tested by Thorchain. It combines Shamir's Secrect Sharing Scheme and Multi-party-computation (MCP) to securely sign transactions or generate new vault shares without revealing anything but the correct output.

***

## Shamir's Secret Sharing (SSS)

This is based on splitting one secret into `n`-amount (total amount) of shares and enabling a subset of all shares, the `m`-amount (needed amount) of these shares to re-construct the secret. This is done by utilizing the mathematical concept of polynomial interpolation. This creates a secure and fault-tolerant option for sharing secrects.

It is based on splitting a secrect into `n`-amount (total amount) of shares and allowing a subset of all shares, the `m`-amount (needed amount) of these shares, to reconstruct the secrect. This is done using the mathematical concept of polynomial interpolation. This creates a secure and fault-tolerant way to share secrets.

<figure><picture><source srcset="../.gitbook/assets/TSS.png" media="(prefers-color-scheme: dark)"><img src="../.gitbook/assets/TSS dark.png" alt="" width="375"></picture><figcaption><p>Example of 3 shares constructing one secret</p></figcaption></figure>

With the Shamir's Secret Sharing Scheme it is possible to construct a secret (private key) with `m`-of-`n` amount of shares to sign transactions on the blockchain or having access to assets.

The flaw in Shamir's Secrect Sharing Scheme lies in that once the secret is constructed again, it is vunerable to exploits or theft and therefore is not secure while using.

With the Shamir Secrect Sharing scheme, it is possible to construct a secrect (private key) with `m` shares to sign transactions on the blockchain or have access to assets. \
The flaw in Shamir secret sharing is that once the secrect is re-constructed, it is vulnerable to exploits or theft and therefore not secure in use.

***

## Multi Party Computation (MPC)

To compensate for the flaw in Shamir's Secrect Sharing Scheme, a Multi Party Computation (MPC) is implemented into the Scheme.\
This allows a solution for proof of access to the secret without ever actually constructing it.

This is archieved through the computation of a function with the secret shares of the participants. This function proofs the access to the secret without the need for the actual secret to be ever created.

Thus, no parties will ever have access to the actual secret (i.e. private key) itself.

To compensate for the flaw in Shamir's Secret Sharing Scheme, a Multi Party Computation (MPC) is implemented in the scheme. It allows to provide a solution for proving access to the secret without ever constructing it.

This is achieved by calculating a function with the Secrect shares of the participants. This function proves access to the secret without ever having to be created.

This improves secrect sharing, where a secrect was originally constructed, separated, and distributed. Leaving the vulnerability that the secrect (private key) was constructed once and potentionally exposed.

<figure><picture><source srcset="../.gitbook/assets/MPC white.png" media="(prefers-color-scheme: dark)"><img src="../.gitbook/assets/MPC dark.png" alt="" width="375"></picture><figcaption><p>Scecret Sharing, Key Distribution</p></figcaption></figure>

***

The combination of SSS and MPC essentially builds the Threshold Signature Scheme, where the Vault Shares get created on each individual device.

Enabled through a key generation function with `n`-amount of shares, this function will output the same public key to all parties, and a different secret share for each participent.\
Only `m`-amount of shares will be needed to construct a valid key sign.

As MPC is an offline computation, that gives it a few advantages for its usage in blockchains:

Where the combination of SSS and MPC essentially builds the Threshold Signature Scheme, where the Vault Shares are created on each device and never leave it.

Enabled by an `n`-share key generation function, this function outputs the same public key to all parties and a different secret share for each participant. This allows access to be proven without revealing the secret shares.

Since MPC is an offline computation, it has few advantages for use in blockchain:

* The `m`-amount can be freely set and later be re-configured through the re-share computation.
* The on-chain footprint is similar to a Single Signature, keeping privacy and lowering transaction fees.
* Faster and more efficient signing computations.

***

To sign a transaction, the previously distributed vault shares are used as the inputs in the MPC to generate a valid and verifiable signature, which will be used on-chain.

To sign a transaction, the previously distributed vault shares are used as input in the MPC to generate a valid and verifiable signature that is used on-chain.

<figure><picture><source srcset="../.gitbook/assets/Tx white.png" media="(prefers-color-scheme: dark)"><img src="../.gitbook/assets/TX black.png" alt="" width="375"></picture><figcaption><p>Signing Transaction</p></figcaption></figure>
