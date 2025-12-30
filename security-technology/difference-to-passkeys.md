---
description: >-
  TSS vs Passkeys comparison. Why threshold signatures offer better crypto
  security: multi-device, no single point of failure, true self-custody.
---

# Difference to Passkeys

## What is a Passkey?

A passkey is an advanced method for securely storing sensitive data, developed under the FIDO (Fast IDentity Online) Alliance. Initially adopted by Apple, Microsoft, and Google, passkeys are gaining traction as an alternative to traditional passwords.

***

## How Passkeys Work

Passkeys utilize asymmetric cryptography—a private key and public key pair—similar to cryptocurrencies. Keys are generated locally on the user's device using secure hardware modules (TPM or Secure Enclave).

The public key is stored by the application server, while the private key remains on the user's device, protected by biometric authentication.

During authentication:
1. The server generates a cryptographic challenge
2. The device signs the challenge with the private key
3. The server verifies the signature with the stored public key

The private key is never transmitted, maintaining security and privacy.

***

## Why Vultisig Doesn't Use Passkeys

While passkeys are secure for general authentication, they have critical limitations for cryptocurrency custody:

### Centralization Concerns

Although passkey technology is open source, it relies on centralized authentication platforms operated by large corporations. This raises concerns about data collection and transparency.

### Single Point of Failure

The authentication process constitutes a single point of failure. A physical attack on the device holding the private key can lead to complete security compromise.

### Crypto-Specific Limitations

1. **Lack of Multi-Chain Support**: Passkeys aren't designed to operate across multiple blockchain networks
2. **Single Signature Mechanism**: Insufficient security for high-stakes transactions that benefit from multi-factor authentication

***

## Vultisig's Approach

Recognizing these shortcomings, Vultisig developed a solution tailored for cryptocurrency:

1. **Open Source Everything**: Transparency through open-source protocols
2. **Multi-Chain Compatibility**: Seamless interoperability across blockchain networks
3. **Multi-Factor Authentication**: Multiple devices eliminate single points of failure

***

## Comparison

| Factor | Passkeys | Vultisig TSS |
|--------|----------|--------------|
| **Single Point of Failure** | Yes (one device) | No (distributed across devices) |
| **Multi-Chain** | No | Yes (30+ chains) |
| **Self-Custody** | Partial (corporate infrastructure) | Full (no third parties) |
| **Device Loss Recovery** | Dependent on platform | Threshold redundancy |
| **Open Source** | Partial | Fully open source |

***

## Conclusion

While passkeys represent progress in general authentication, they don't meet the security requirements for cryptocurrency custody. Vultisig's threshold signature approach provides the multi-factor, multi-chain, fully self-custodial solution that digital assets require.

***

## Related

- [Security & Technology Overview](README.md)
- [Difference to Multi-Signatures](difference-to-multi-sig.md)
