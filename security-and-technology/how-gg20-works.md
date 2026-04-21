---
description: >-
  GG20 threshold signature protocol explained. How Vultisig's legacy TSS
  works: secret sharing, MPC signing rounds, and security model.
---

# How GG20 Works

## Concept

The Threshold Signature Scheme (TSS) is part of the broader field of Multi-Party Computation (MPC) and has evolved significantly over the past few decades. The concept of secret sharing was originally introduced by Adi Shamir in 1979, with threshold signatures developing as an extension of this foundational work.

In 2018, [the "GG18" paper by Gennaro-Goldfeder](https://eprint.iacr.org/2019/114.pdf) introduced significant advancements, simplifying and enhancing the efficiency of TSS. This was further improved in 2020 with the ["GG20" paper](https://eprint.iacr.org/2020/540), which introduced "identifiable abort"—the ability to identify which party caused a protocol failure.

Vultisig utilizes this improved version of TSS, implemented in a comprehensive [library from Binance](https://github.com/bnb-chain/tss-lib) and thoroughly tested in production by THORChain.

GG20 TSS integrates Homomorphic Secret Sharing, Zero Knowledge Proofs, and Multi-Party Computation (MPC) to securely sign transactions without revealing sensitive information.

{% hint style="info" %}
For an excellent summary of MPC protocol history: [https://www.cryptologie.net/article/605/whats-out-there-for-ecdsa-threshold-signatures/](https://www.cryptologie.net/article/605/whats-out-there-for-ecdsa-threshold-signatures/)
{% endhint %}

***

## Homomorphic Secret Sharing

The homomorphic secret sharing employed by Vultisig is founded on the Paillier encryption scheme, introduced by Pascal Paillier in 1999. This additive homomorphic encryption method allows ciphertexts to be combined through mathematical addition, enabling secure computations on encrypted data.

{% hint style="info" %}
Homomorphic encryption preserves the structure of data, allowing mathematical operations on encrypted objects without altering their underlying integrity.
{% endhint %}

In GG20, Paillier encryption plays a crucial role in the Multiplicative-to-Additive (MtA) conversion process, essential for threshold ECDSA signatures. This allows parties to securely compute the product of their secret values without revealing those values.

***

## Zero Knowledge Proof

Zero Knowledge Proof (ZKP) enables a proving entity to convince a verifying entity that a statement is true without disclosing additional information.

Key properties of ZKPs:

* **Zero-Knowledge**: The verifier learns nothing beyond the statement's truth
* **Soundness**: Computationally infeasible to prove false statements
* **Completeness**: True statements can be proven with high probability

In GG20, multiple zero-knowledge proofs verify that participants follow the protocol correctly, including proofs of knowledge for discrete logarithms and range proofs to ensure values are within appropriate bounds.

***

## Multi-Party Computation (MPC)

MPC enables secure computation even with potentially dishonest participants. Functions are computed on secret shares without revealing those shares.

{% hint style="success" %}
The private key is never actually constructed in Vultisig.
{% endhint %}

The GG20 protocol requires 6 rounds of communication between participants to generate a valid signature. This provides "identifiable abort"—if something goes wrong, the system can identify which participant caused the failure.

Since MPC is an offline computation, it offers several advantages:

* The number of shares can be freely set and reconfigured
* On-chain footprint equals a single signature
* Faster and more efficient signing computations

<figure><picture><source srcset="../.gitbook/assets/Tx white.png" media="(prefers-color-scheme: dark)"><img src="../.gitbook/assets/TX black.png" alt="" width="375"></picture><figcaption><p>Signing Transaction</p></figcaption></figure>

***

## Conclusion

GG20 provides a robust security framework with features like identifiable abort. While it requires 6 communication rounds and involves computationally intensive operations, it offers strong security guarantees when properly implemented.

***

## References

1. Gennaro, R., & Goldfeder, S. (2020). ["One Round Threshold ECDSA with Identifiable Abort."](https://eprint.iacr.org/2020/540)
2. Shamir, A. (1979). "How to share a secret." Communications of the ACM, 22(11), 612-613.
3. Paillier, P. (1999). "Public-key cryptosystems based on composite degree residuosity classes."
