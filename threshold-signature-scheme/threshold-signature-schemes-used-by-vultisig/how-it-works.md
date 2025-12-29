---
description: >-
  GG20 threshold signature protocol explained. How Vultisig's legacy TSS
  works: secret sharing, MPC signing rounds, and security model.
---

# How GG20 works

## Concept

The Threshold Signature Scheme (TSS) is part of the broader field of Multi-Party-Computation (MPC) and has evolved significantly over the past few decades. The concept of secret sharing was originally introduced by Adi Shamir in 1979, with threshold signatures developing as an extension of this foundational work.

In 2018, [the "GG18" paper by Gerrano-Goldfeder](https://eprint.iacr.org/2019/114.pdf) introduced significant advancements, simplifying and enhancing the efficiency of TSS, thereby making it feasible for practical implementation. This was further improved in 2020 with the ["GG20" paper](https://eprint.iacr.org/2020/540), which introduced the concept of "identifiable abort" - the ability to identify which party caused a protocol failure.

Vultisig utilizes this improved version of TSS, known as GG20, which has been implemented in a comprehensive [library from Binance](https://github.com/bnb-chain/tss-lib) and thoroughly tested in production by THORChain.

GG20 TSS integrates Homomorphic Secret Sharing, Zero Knowledge Proofs, and Multi-Party Computation (MPC) to securely sign transactions or generate new Vault shares without revealing any sensitive information. This process ensures that only the correct output is exposed, preserving the security and privacy of the underlying data.

By leveraging these advanced cryptographic techniques, Vultisig provides a highly secure and cutting-edge solution for managing digital assets.

{% hint style="info" %}
For an excellent summary of the history of MPC protocols, start here[https://www.cryptologie.net/article/605/whats-out-there-for-ecdsa-threshold-signatures/](https://www.cryptologie.net/article/605/whats-out-there-for-ecdsa-threshold-signatures/)
{% endhint %}

***

## Homomorphic Secret Sharing

The homomorphic secret sharing employed by Vultisig is founded on the Paillier encryption scheme, introduced by Pascal Paillier in 1999. This additive homomorphic encryption method allows ciphertexts to be combined through mathematical addition, enabling secure computations on encrypted data without revealing the underlying information.

{% hint style="info" %}
Homomorphic encryption preserves the structure of data, allowing mathematical operations to be performed on encrypted objects without altering their underlying integrity.
{% endhint %}

In the GG20 protocol, Paillier encryption plays a crucial role in the Multiplicative-to-Additive (MtA) conversion process, which is essential for threshold ECDSA signatures. This process allows parties to securely compute the product of their secret values without revealing those values to each other.

These properties enable the TSS to perform secret sharing and mathematical operations directly on encrypted shares without requiring decryption. This ensures that the secret shares remain secure and private while allowing efficient computation.

***

## Zero Knowledge Proof

Zero Knowledge Proof (ZKP) is a cryptographic technique that enables a proving entity to convince a verifying entity that a specific statement is true without disclosing any additional information that could compromise privacy or security.

The key properties of ZKPs are:

* **Zero-Knowledge**: The verifier learns nothing about the proving entity’s private information, other than the fact that the statement is true.
* **Soundness**: It is computationally infeasible for an adversary to convince the verifier of a false statement.
* **Completeness**: If the statement is true, the proving entity can convince the verifier with high probability.

These properties make ZKPs an ideal fit for threshold signature schemes, offering robust tools for ensuring the privacy and security of sensitive information.

In the GG20 protocol, multiple zero-knowledge proofs are used to verify that participants are following the protocol correctly. These include proofs of knowledge for discrete logarithms and range proofs to ensure that values are within appropriate bounds. These proofs are essential for preventing malicious behavior that could compromise the security of the system.

The ZKP techniques used in Vultisig's TSS implementation include various forms of non-interactive zero-knowledge proofs that allow verification without additional communication rounds, enhancing both security and efficiency.

## Multi Party Computation (MPC)

Multi Party Computation includes a broad field of cryptography where Vultisig's Threshold Signature Scheme (TSS) falls into. It enables secure computation even in the presence of a potentially dishonest majority. This approach allows participants to prove access to a secret without ever reconstructing it.

Using MPC, functions can be computed on the secret shares held by participants without revealing those shares. This method ensures that the access to the secret is verified and proven without ever generating or exposing the actual secret, thereby maintaining high security and privacy standards.

{% hint style="success" %}
The private key is never actually constructed in Vultisig
{% endhint %}

Thus, no party will ever have access to the actual secret (i.e., private key) held by the other parties.

The GG20 protocol that Vultisig implements requires 6 rounds of communication between participants to generate a valid signature. This multi-round process ensures security but requires participants to exchange multiple messages in a specific sequence. The protocol provides "identifiable abort," meaning if something goes wrong during the signing process, the system can identify which participant caused the failure - a significant security feature for accountability.

***

This MPC TSS allows for key generation functions using an `n`-amount of shares, where the function outputs the same public key to all parties and a unique secret share for each participant. Only an `t`-number of shares are needed to construct a valid key signature.

Since MPC is an offline computation, it offers several advantages for use in blockchains:

* The `n`-amount can be freely set and later reconfigured through the re-share computation.
* The on-chain footprint is similar to a single signature, maintaining privacy and reducing transaction fees.
* It enables faster and more efficient signing computations.

***

To sign a transaction, the vault shares are used as inputs in the MPC process to generate a valid and verifiable signature, which will then be used to sign a transaction on-chain.

<figure><picture><source srcset="../../.gitbook/assets/Tx white.png" media="(prefers-color-scheme: dark)"><img src="../../.gitbook/assets/TX black.png" alt="" width="375"></picture><figcaption><p>Signing Transaction</p></figcaption></figure>

***

## Conclusion

The combination of these concepts essentially forms the foundation of the Threshold Signature Scheme used by Vultisig. Vault Shares are created on each individual device, allowing functions to be applied without revealing the shares.

The GG20 protocol provides a robust security framework with features like identifiable abort, making it suitable for high-security applications. While it requires 6 communication rounds for signature generation and involves computationally intensive operations, it offers strong security guarantees when properly implemented.

This approach ensures that the privacy and security of the Vault Shares are maintained at all times, while using this technology also reduces the on-chain footprint, improving efficiency and safety for the user.

***

## References

1. Gennaro, R., & Goldfeder, S. (2020). ["One Round Threshold ECDSA with Identifiable Abort."](https://eprint.iacr.org/2020/540)
2. Shamir, A. (1979). "How to share a secret." Communications of the ACM, 22(11), 612-613.
3. Paillier, P. (1999). "Public-key cryptosystems based on composite degree residuosity classes." In International Conference on the Theory and Applications of Cryptographic Techniques (pp. 223-238)
