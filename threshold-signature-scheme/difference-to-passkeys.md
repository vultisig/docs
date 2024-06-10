---
description: What is the difference between TSS and Passkeys??
---

# Difference to Passkeys

## What is a passkey?

A passkey is an advanced method for securely storing sensitive data, designed to facilitate access to applications and websites. \
This technology was developed under the auspices of the FIDO (Fast IDentity Online) Alliance as part of an initiative to establish a new authentication standard. \
Initially adopted by major tech companies such as Apple, Microsoft, and Google, passkeys are rapidly gaining traction as a robust alternative to traditional passwords.

## How do passkeys work?

Passkeys utilize an asymmetric cryptographic pair, consisting of a private key and a public key, similar to the principles employed in the cryptocurrency domain. These keys are generated locally on the user's device using secure hardware modules, such as the Trusted Platform Module (TPM) or Secure Enclave.

The public key is transmitted to and stored by the application or website's server, while the private key remains securely on the user's device, protected by advanced biometric authentication mechanisms like fingerprint or facial recognition.

During the authentication process, the website or application generates a unique cryptographic challenge, which is sent to the user's device. The device responds by using the private key to create a digital signature of the challenge. This signature is then verified by the server using the stored public key.

The cryptographic process ensures that the user possesses the private key without ever transmitting it, thereby maintaining a high level of security and privacy. \
This mechanism is analogous to the verification process in cryptocurrency transactions, ensuring both integrity and authenticity of the authentication process.

## Why does Vultisig not use passkeys?

While passkeys are generally considered highly secure and hold great promise as a replacement for traditional passwords, the implementation details reveal some critical vulnerabilities that must be addressed, especially when securing high-value assets and wealth.

Firstly, although passkey technology is open source, it relies on centralized authentication platforms operated by large corporations. This dependency raises concerns about data collection practices and the extent to which user data might be harvested by these entities. The lack of transparency in data handling by these corporations is a significant risk factor.

Secondly, the authentication process itself constitutes a single point of failure. A physical attack on the device holding the private key can lead to a complete security compromise, undermining the reliability of the passkey system. \
Additionally, if the authentication relies on cookies for session management, attackers could potentially circumvent the system by stealing the authentication cookie, thereby gaining unauthorized access without needing the private key.

Specific to the cryptocurrency industry, passkeys exhibit additional limitations:

1. **Lack of Multi-Chain Support**: Passkey technology is not inherently designed to operate across multiple blockchain networks, which is a critical requirement for the dynamic and diverse cryptocurrency ecosystem.
2. **Single Signature Mechanism**: Passkeys employ a single signature for authentication, which may not provide sufficient security for complex, high-stakes transactions that benefit from multi-signature or multi-factor authentication.

Recognizing these shortcomings, Vultisig has embarked on developing an advanced solution tailored for the cryptocurrency space. Our objectives are to set new standards by ensuring:

1. **Open Source Everything**: Transparency and community trust through open-source protocols and implementations.
2. **Multi-Chain Compatibility**: Seamless interoperability across various blockchain networks.
3. **Multi-Factor Authentication**: Enhanced security through multiple layers of authentication, mitigating the risks associated with single points of failure.

By addressing these critical issues, Vultisig aims to advance the security and functionality of the cryptocurrency industry.
