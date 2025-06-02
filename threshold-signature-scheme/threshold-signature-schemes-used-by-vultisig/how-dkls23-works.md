---
description: How do Threshold Signatures work with the DKLS23 protocol?
---

# How DKLS23 works


# Evolution and Innovation

The field of Threshold Signature Schemes (TSS) has seen remarkable advancements since Adi Shamir's groundbreaking work on secret sharing in 1979. While early protocols laid important foundations, recent innovations have dramatically improved efficiency and security.

The [DKLS23 protocol](https://eprint.iacr.org/2023/765), introduced in 2023 by Doerner, Kondi, Lee, and Shelat, represents a significant leap forward in threshold ECDSA implementation. Building upon the security framework established by earlier protocols like GG18 and GG20, DKLS23 achieves the remarkable feat of reducing communication rounds from 6 to just 3 while maintaining robust security against malicious adversaries.

This protocol marks a paradigm shift in how threshold signatures are implemented, moving away from homomorphic encryption toward more efficient oblivious transfer techniques. The result is a protocol that is not only more efficient but also simpler to implement correctly, reducing the risk of security vulnerabilities.

DKLS23 integrates Oblivious Transfer, Zero Knowledge Proofs, and Multi-Party Computation (MPC) to create a threshold signature system that offers exceptional performance without compromising security. This approach ensures that sensitive information remains protected throughout the signing process while delivering significantly faster transaction times.

***


# Three-Round Architecture

The most distinctive feature of DKLS23 is its streamlined three-round signing process, which represents a 50% reduction in communication rounds compared to GG20's six-round approach. This architectural innovation has profound implications for both performance and reliability.

### Round Structure

#### Round 1: Commitment Phase

* Participants generate and exchange commitments to their secret values
* Unlike GG20, which requires separate rounds for nonce generation and sharing, DKLS23 combines these steps
* The protocol employs an intermediate representation of ECDSA signatures that enables this consolidation

#### Round 2: Multiplication Phase

* Secure two-party multiplication is performed using oblivious transfer
* This replaces the more computationally intensive Multiplicative-to-Additive (MtA) conversion in GG20
* Statistical consistency checks are employed to ensure security

#### Round 3: Signature Completion

* Final signature components are computed and combined
* The protocol ensures that the resulting signature is valid and indistinguishable from a standard ECDSA signature

This streamlined architecture significantly reduces latency, especially in high-latency or unstable network environments. The fewer communication rounds also improve reliability by reducing the opportunities for network failures to disrupt the signing process.

### Performance Implications

The three-round design delivers substantial performance benefits:

* Reduced Latency: Fewer communication rounds mean less time waiting for network exchanges
* Improved Reliability: Fewer rounds reduce the chance of network failures disrupting the signing process
* Enhanced Scalability: The protocol maintains efficiency even as the number of participants increases

For applications requiring frequent signing operations or operating in challenging network environments, these improvements translate to a significantly better user experience.

***


# Oblivious Transfer

While GG20 relies on Paillier's homomorphic encryption scheme, DKLS23 takes a fundamentally different approach by building upon oblivious transfer (OT), a cryptographic primitive that offers significant efficiency advantages.

### Understanding Oblivious Transfer


Oblivious transfer is a protocol where a sender transfers one of potentially many pieces of information to a receiver, while remaining oblivious as to which piece was transferred. In its simplest form (1-out-of-2 OT), the sender has two messages, and the receiver obtains one of them without the sender learning which one was chosen.

This primitive serves as the foundation for secure two-party computation in DKLS23, enabling parties to perform joint computations without revealing their private inputs.

### OT Extensions


A key innovation in modern OT implementations is the concept of OT extensions, which allow for the efficient implementation of many OT instances from a small number of base OTs. DKLS23 leverages these extensions to achieve remarkable performance improvements.

The protocol employs a two-round vectorized multiplication protocol based on OT that enables secure multiplication of secret values without revealing those values to other participants. This approach eliminates the need for computationally intensive homomorphic encryption operations required in GG20.

### Efficiency Advantages


The OT-based approach offers several significant advantages:

* Computational Efficiency: OT operations are substantially faster than Paillier encryption/decryption
* Reduced Bandwidth: Smaller message sizes compared to Paillier ciphertexts
* Better Parallelization: More opportunities for parallel computation

These efficiency gains are particularly evident in performance benchmarks, where DKLS23 demonstrates signing times that are 5-10x faster than previous protocols, especially on resource constrained devices.

***


# Statistical Security Measures

DKLS23 employs a different approach to security verification compared to GG20, focusing on statistical consistency checks rather than complex zero-knowledge proofs. This approach maintains strong security guarantees while reducing cryptographic overhead.

### Simplified Verification


While GG20 uses multiple zero-knowledge proofs including range proofs to ensure values are within appropriate bounds, DKLS23 employs fewer and simpler verification mechanisms:

* Commitment Schemes: Used to ensure participants cannot change their inputs after seeing others' values
* Statistical Checks: Verify the consistency of participants' behavior throughout the protocol
* Simplified ZKPs: Where zero-knowledge proofs are used, they are simpler and more efficient

This simplified approach contributes to the protocol's overall efficiency while maintaining strong security properties.

### Security Properties


DKLS23 provides information-theoretic UC-security against malicious adversaries with dishonest majority. Similar to GG20's security in the Universal Composability (UC) framework, DKLS23 maintains security even when composed with other protocols.

The protocol offers several security advantages:

* Reduced Attack Surface: Fewer cryptographic primitives mean fewer potential vulnerabilities
* No Paillier Vulnerabilities: Not susceptible to attacks like the "Alpha-Rays Attack" that can affect GG20 implementations without proper range proofs
* No Early Nonce Leakage: The protocol design prevents leakage of the nonce R, closing certain attack vectors
* Simpler Security Proofs: Relies on fewer assumptions, making security analysis more straightforward

These security properties make DKLS23 not only more efficient but also potentially more secure in practice, as simpler systems typically have fewer implementation vulnerabilities.


***

## Multi-Party Computation


DKLS23, like GG20, operates within the broader framework of Multi-Party Computation (MPC), allowing participants to compute functions on shared data without revealing their individual inputs. This approach is fundamental to threshold signature schemes.

### Distributed Key Generation


The key generation process in DKLS23 follows a commit-release-and-complain procedure that is simpler than GG20's approach while maintaining strong security properties:

1. Commitment Phase: Participants generate random values and commit to them
2. Release Phase: Commitments are opened and verified
3. Complaint Phase: Any inconsistencies are identified and addressed

This process results in each participant holding a share of the private key, with no single party having access to the complete key. The corresponding public key is known to all participants and can be used for standard ECDSA verification.

### Threshold Properties


DKLS23 supports flexible threshold configurations, allowing for `t`-of-`n` setups where any `t` participants from a group of `n` can collaborate to create valid signatures. This flexibility is crucial for real-world applications where different security models may be required.

The protocol also supports dynamic participant sets and key resharing, allowing the system to adapt to changing requirements without generating new public keys.

# Implementation Considerations


One of DKLS23's significant advantages is its simplified implementation compared to GG20. This simplification reduces the risk of security vulnerabilities that can arise from implementation errors.

### Reduced Complexity


DKLS23 requires fewer cryptographic primitives and simpler security mechanisms:

* No Paillier Implementation: Eliminates the need for complex homomorphic encryption
* Simpler Proofs: Fewer and less complex zero-knowledge proofs
* Fewer Rounds: Simpler protocol flow with fewer state transitions

This reduced complexity makes the protocol easier to implement correctly and easier to audit for security vulnerabilities.

### Performance Optimization


When implementing DKLS23, several optimizations can further improve performance:

* Parallelization: Many operations can be performed in parallel
* Precomputation: Certain values can be precomputed to reduce online signing time
* Batch Processing: Multiple signatures can be processed efficiently in batch mode

These optimizations can further enhance the already impressive performance of the protocol, making it suitable for even the most demanding applications.

***

# Comparative Advantages


While both GG20 and DKLS23 provide secure threshold signature capabilities, DKLS23 offers several significant advantages that make it the preferred choice for many applications.

### Performance Improvements


DKLS23 delivers substantial performance improvements over GG20:

* Signing Speed: 5-10x faster signing times
* Communication Efficiency: 50% reduction in communication rounds
* Resource Usage: Lower computational and bandwidth requirements
* Scalability: Better performance with increasing numbers of participants

These improvements are particularly significant for mobile applications, IoT devices, and other resource-constrained environments.

### Security Enhancements


DKLS23 also offers several security advantages:

* Simplified Security Model: Fewer cryptographic primitives reduce the attack surface
* Resistance to Known Attacks: Not vulnerable to certain attacks that can affect GG20 implementations
* Easier Verification: Simpler security proofs make verification more straightforward

These security enhancements make DKLS23 not only more efficient but potentially more secure in practice.

***

# Conclusion

The DKLS23 protocol represents a significant advancement in threshold signature technology, offering substantial improvements in efficiency, security, and implementation simplicity compared to earlier protocols like GG20.

By reducing communication rounds from 6 to 3 and replacing homomorphic encryption with more efficient oblivious transfer techniques, DKLS23 delivers signing times that are 5-10x faster than previous protocols. This performance improvement, combined with simplified security mechanisms and reduced implementation complexity, makes DKLS23 the fastest TSS protocol up to this date.

While GG20 remains a secure and well-established protocol, DKLS23's advantages make it the preferred choice for new implementations of threshold ECDSA, particularly in scenarios requiring high performance and robust security.

The advancements in DKLS23 represent a significant step forward in making threshold signatures practical for widespread adoption, bringing the security benefits of distributed key management to a broader range of applications and users.

# References

* Doerner, J., Kondi, Y., Lee, E., & Shelat, A. (2023). ["Threshold ECDSA in Three Rounds."](https://eprint.iacr.org/2023/765) (https://eprint.iacr.org/2023/765)
* Gennaro, R., & Goldfeder, S. (2020). ["One Round Threshold ECDSA with Identifiable Abort."](https://eprint.iacr.org/2020/540) (https://eprint.iacr.org/2020/540)
* Shamir, A. (1979). "How to share a secret." Communications of the ACM, 22(11), 612-613.
* Silence Laboratories. "[Silent Shard's threshold ECDSA signatures implementing DKLs23 protocol.](https://github.com/silence-laboratories/dkls23)"(https://github.com/silence-laboratories/dkls23)
* [DKLs.info](http://dkls.info/). "Threshold ECDSA in Three Rounds."(http://dkls.info/)
