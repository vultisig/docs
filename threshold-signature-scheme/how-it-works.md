---
description: How do Threshold Signatures work ?
---

# How it works

## Concept

The Threshold Signature Scheme (TSS) was originally introduced by Adi Shamir and Yael Tauman in 1998. However, due to its complex nature and resource-intensive computation, it was impractical for widespread adoption. With the introduction of the "GG18" paper by Gerrano-Goldfeder in 2018 and its breakthrough in simplifying and improving the efficiency of the TSS scheme, it became practical for implementation and adoption. \
Vultisig uses the improved version of this TSS (GG20), which has been implemented and thoroughly tested by THORChain.\
\
It mainly combines Homomorphic Secret Sharing, Zero Knowledge Proofs and Multi Party Computation (MPC) to securely sign transactions or generate new Vault shares without revealing anything but the correct output.

***

## Homomorphic Secret Sharing

The homomorphic secret sharing used in Vultisig is based on the Paillier encryption scheme. Named after its inventor Pascal Paillier, who introduced it in 1998. This scheme is an additive homomorphic encryption scheme, which means that ciphertexts can be added homomorphically by mathematical addition.

{% hint style="info" %}
Homomorphic means that the structure of the objects is preserved even when mathematical operations are applied to the objects.
{% endhint %}

These properties allow the TSS to enable secret sharing and apply manipulation, such as mathematics, without the need to decrypt the shares. Essentially, this ensures that the secret shares are kept secure and private, while providing an efficient way to compute them.

***

## Zero Knowledge Proof

Zero Knowledge Proof (ZKP) is a cryptographic method that allows a proving entity to convince a verifying entity that a given fact is true without revealing any additional information that could be used to compromise privacy or security.\
\
Important properties of ZKPS are:

1. **Zero-knowledge**: The verifier learns nothing about the proving entity’s private information, except for the fact that the statement is true.
2. **Soundness**: It is computationally infeasible for an adversary to convince the verifier that a false statement is true.
3. **Completeness**: If the statement is indeed true, the proving entity can convince the verifier with high probability.

Therefore, ZKPs are a perfect fit for threshold signature schemes, providing a powerful tool for ensuring the privacy and security of sensitive information.\
\
The ZKP used in Vultisig's TSS are "zk-SNARKs" (Zero-Knowledge Succinct Non-Interactive Argument of Knowledge), a type of ZKP that allows the verification of complex statements without revealing the underlying data.

## Multi Party Computation (MPC)

Vultisig's TSS uses Multi Party Computation (MPC) to achieve secure computation with a possible dishonest majority. This allows a solution for proving access to the secret without ever actually constructing it.

This allows functions to be computed using the secret shares of the participants without revealing them. This function proves access to the secret without ever creating the actual secret.

{% hint style="success" %}
The private key is never actually constructed in Vultisig
{% endhint %}

Thus, no party will ever have access to the actual secret (i.e. private key) of the other parties.

***

This MPC allows key generation functions with `n`-amount of shares, this function will output the same public key to all parties, and a different secret share for each participant.\
Only `m`-number of shares will be needed to construct a valid key sign.

Since MPC is an offline calculation, it has several advantages for use in blockchains:

* The `m`-amount can be freely set and later be reconfigured through the re-share computation.
* The on-chain footprint is similar to a single signature, keeping privacy and lowering transaction fees.
* Faster and more efficient signing computations.

***

To sign a transaction, the vault shares are used as the inputs in the MPC to generate a valid and verifiable signature, which will be used on-chain.

<figure><picture><source srcset="../.gitbook/assets/Tx white.png" media="(prefers-color-scheme: dark)"><img src="../.gitbook/assets/TX black.png" alt="" width="375"></picture><figcaption><p>Signing Transaction</p></figcaption></figure>

***

## Conclusion

The combination of the above concepts essentially builds the Threshold Signature Scheme used by Vultisig, where the Vault Shares are created on each individual device with the possibility to apply functions without the need to reveal the shares.&#x20;

The privacy and security of the Vault Shares are guaranteed at all times.
