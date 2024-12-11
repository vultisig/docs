---
description: How do Threshold Signatures Work ?
---

# How it works

## Concept

The Threshold Signature Scheme (TSS) is part of the broader field of Multi-Party-Computation (MPC) and was originally introduced by Adi Shamir and Yael Tauman in 1998. However, due to its complex nature and resource-intensive computation, it was impractical for widespread adoption.

In 2018, [the "GG18" paper by Gerrano-Goldfeder](https://eprint.iacr.org/2019/114.pdf) introduced significant advancements, simplifying and enhancing the efficiency of TSS, thereby making it feasible for practical implementation.

Vultisig utilizes an [improved version of this TSS, known as GG20,](https://github.com/bnb-chain/tss-lib) which has been rigorously implemented and thoroughly tested by THORChain.

[GG20 TSS integrates Homomorphic Secret Sharing, Zero Knowledge Proofs, and Multi-Party Computation (MPC)](https://eprint.iacr.org/2020/540) to securely sign transactions or generate new Vault shares without revealing any sensitive information. This process ensures that only the correct output is exposed, preserving the security and privacy of the underlying data.

By leveraging these advanced cryptographic techniques, Vultisig provides a highly secure and cutting-edge solution for managing digital assets.

{% hint style="info" %}
For an excellent summary of the history of MPC protocols, start here\
[https://www.cryptologie.net/article/605/whats-out-there-for-ecdsa-threshold-signatures/](https://www.cryptologie.net/article/605/whats-out-there-for-ecdsa-threshold-signatures/)
{% endhint %}

***

## Homomorphic Secret Sharing

The homomorphic secret sharing employed by Vultisig is founded on the Paillier encryption scheme, introduced by Pascal Paillier in 1998. This additive homomorphic encryption method allows ciphertexts to be combined through mathematical addition, enabling secure computations on encrypted data without revealing the underlying information.

{% hint style="info" %}
Homomorphic encryption preserves the structure of data, allowing mathematical operations to be performed on encrypted objects without altering their underlying integrity.
{% endhint %}

These properties enable the TSS to perform secret sharing and mathematical operations directly on encrypted shares without requiring decryption. This ensures that the secret shares remain secure and private while allowing efficient computation.

***

## Zero Knowledge Proof

Zero Knowledge Proof (ZKP) is a cryptographic technique that enables a proving entity to convince a verifying entity that a specific statement is true without disclosing any additional information that could compromise privacy or security.

The key properties of ZKPs are:

* **Zero-Knowledge**: The verifier learns nothing about the proving entity’s private information, other than the fact that the statement is true.
* **Soundness**: It is computationally infeasible for an adversary to convince the verifier of a false statement.
* **Completeness**: If the statement is true, the proving entity can convince the verifier with high probability.

These properties make ZKPs an ideal fit for threshold signature schemes, offering robust tools for ensuring the privacy and security of sensitive information.

The ZKP used in Vultisig's TSS is "zk-SNARKs" (Zero-Knowledge Succinct Non-Interactive Argument of Knowledge). This advanced type of ZKP allows the verification of complex statements without revealing the underlying data, enhancing both security and efficiency.

## Multi Party Computation (MPC)

Multi Party Computation includes a broad field of cryptography where Vultisig's Threshold Signature Scheme (TSS) falls into. It enables secure computation even in the presence of a potentially dishonest majority. This approach allows participants to prove access to a secret without ever reconstructing it.

Using MPC, functions can be computed on the secret shares held by participants without revealing those shares. This method ensures that the access to the secret is verified and proven without ever generating or exposing the actual secret, thereby maintaining high security and privacy standards.

{% hint style="success" %}
The private key is never actually constructed in Vultisig
{% endhint %}

Thus, no party will ever have access to the actual secret (i.e., private key) held by the other parties.

***

This MPC TSS allows for key generation functions using an `n`-amount of shares, where the function outputs the same public key to all parties and a unique secret share for each participant. Only an `m`-number of shares are needed to construct a valid key signature.

Since MPC is an offline computation, it offers several advantages for use in blockchains:

* The `m`-amount can be freely set and later reconfigured through the re-share computation.
* The on-chain footprint is similar to a single signature, maintaining privacy and reducing transaction fees.
* It enables faster and more efficient signing computations.

***

To sign a transaction, the vault shares are used as inputs in the MPC process to generate a valid and verifiable signature, which will then be used to sign a transaction on-chain.

<figure><picture><source srcset="../.gitbook/assets/Tx white.png" media="(prefers-color-scheme: dark)"><img src="../.gitbook/assets/TX black.png" alt="" width="375"></picture><figcaption><p>Signing Transaction</p></figcaption></figure>

***

## Conclusion

The combination of these concepts essentially forms the foundation of the Threshold Signature Scheme used by Vultisig. Vault Shares are created on each individual device, allowing functions to be applied without revealing the shares.\
\
This ge approach ensures that the privacy and security of the Vault Shares are maintained at all times.\
\
While using this technology also reduces the on-chain footprint, improving efficiency and safety for the user.

