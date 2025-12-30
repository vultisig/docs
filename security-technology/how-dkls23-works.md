---
description: >-
  DKLS23 protocol deep-dive. Vultisig's modern TSS: faster signing, fewer
  rounds, improved security. Developed with Silence Laboratories.
---

# How DKLS23 Works

## Evolution and Innovation

The [DKLS23 protocol](https://eprint.iacr.org/2023/765), introduced in 2023 by Doerner, Kondi, Lee, and Shelat, represents a significant advancement in threshold ECDSA. Building upon GG18 and GG20, DKLS23 reduces communication rounds from 6 to 3 while maintaining robust security.

This protocol shifts from homomorphic encryption toward more efficient oblivious transfer techniques—simpler to implement correctly with fewer potential vulnerabilities.

DKLS23 integrates Oblivious Transfer, Zero Knowledge Proofs, and Multi-Party Computation (MPC) to create a threshold signature system with exceptional performance.

***

## Three-Round Architecture

The most distinctive feature of DKLS23 is its streamlined three-round signing process—a 50% reduction from GG20.

### Round 1: Commitment Phase

* Participants generate and exchange commitments to secret values
* Nonce generation and sharing are combined (unlike GG20's separate rounds)
* Uses an intermediate representation of ECDSA signatures

### Round 2: Multiplication Phase

* Secure two-party multiplication via oblivious transfer
* Replaces GG20's computationally intensive MtA conversion
* Statistical consistency checks ensure security

### Round 3: Signature Completion

* Final signature components computed and combined
* Result is indistinguishable from standard ECDSA signature

### Performance Implications

* **Reduced Latency**: Fewer communication rounds
* **Improved Reliability**: Fewer failure points
* **Enhanced Scalability**: Maintains efficiency as participants increase

***

## Oblivious Transfer

While GG20 relies on Paillier's homomorphic encryption, DKLS23 uses oblivious transfer (OT)—a cryptographic primitive offering significant efficiency advantages.

### Understanding Oblivious Transfer

In 1-out-of-2 OT, a sender transfers one of two messages to a receiver without learning which was chosen. This enables secure two-party computation without revealing private inputs.

### OT Extensions

OT extensions allow many OT instances from a small number of base OTs. DKLS23 leverages these for a two-round vectorized multiplication protocol, eliminating computationally intensive homomorphic operations.

### Efficiency Advantages

* **Computational Efficiency**: OT operations faster than Paillier encryption
* **Reduced Bandwidth**: Smaller message sizes
* **Better Parallelization**: More opportunities for parallel computation

***

## Statistical Security Measures

DKLS23 uses statistical consistency checks rather than complex zero-knowledge proofs:

* **Commitment Schemes**: Prevent input changes after seeing others' values
* **Statistical Checks**: Verify consistent behavior throughout protocol
* **Simplified ZKPs**: Where used, they are simpler and more efficient

### Security Properties

DKLS23 provides information-theoretic UC-security against malicious adversaries with dishonest majority:

* **Reduced Attack Surface**: Fewer cryptographic primitives
* **No Paillier Vulnerabilities**: Not susceptible to "Alpha-Rays Attack"
* **No Early Nonce Leakage**: Protocol design prevents R leakage
* **Simpler Security Proofs**: Fewer assumptions

***

## Why Vultisig Upgraded to DKLS23

In early 2025, Vultisig transitioned to DKLS23 in cooperation with [Silence Laboratories](https://github.com/silence-laboratories/dkls23):

1. **Faster transactions**: Signing in milliseconds rather than seconds
2. **Better reliability**: Fewer rounds mean less chance of network failures
3. **Improved compatibility**: Works efficiently on resource-constrained devices
4. **Enhanced battery life**: Less computational work
5. **WASM Compatibility**: Enables the Vultisig Extension

***

## Comparing the Protocols

| Feature | GG20 | DKLS23 | User Impact |
|---------|------|--------|-------------|
| Signing Speed | Slower | 5-10x faster | Quicker approvals |
| Communication Rounds | 6 rounds | 3 rounds | Works on spotty connections |
| Security Level | Very High | Very High | Both excellent |
| Network Reliability | More sensitive | More robust | Fewer failed transactions |

***

## References

* Doerner, J., Kondi, Y., Lee, E., & Shelat, A. (2023). ["Threshold ECDSA in Three Rounds."](https://eprint.iacr.org/2023/765)
* Gennaro, R., & Goldfeder, S. (2020). ["One Round Threshold ECDSA with Identifiable Abort."](https://eprint.iacr.org/2020/540)
* Silence Laboratories. ["Silent Shard's threshold ECDSA signatures implementing DKLS23."](https://github.com/silence-laboratories/dkls23)
