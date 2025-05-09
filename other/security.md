---
description: Everything related to security what affects the Vultisig Project
---

# Security

## Audits

### Kudelski Audit of `mobile-tss-lib`

<figure><img src="../.gitbook/assets/TSS-Security.png" alt=""><figcaption><p>Kudelski Findings</p></figcaption></figure>

{% file src="../.gitbook/assets/Kudelski_Security_Vultisig_Mob_CR_v1.1_Public.pdf" %}
Audit Summery
{% endfile %}

> 1.5 Follow-up
>
> After the draft report (v1.0) was delivered, the client addressed all findings in the following
>
> PRs:
>
> • Audit 1 #17 (commit 06fc76f4d6d34f21fa5d1cafd1eb594d8ac4fdd7)
>
> • Audit 2 #18 (commit 2577eb3b00d4d58a7318fa0ada726ba7965579ab)

### **$VULT Contract Audit**

The audit was done by [Code4rena ](https://x.com/code4rena)for  the $VULT contract and can be found [here](https://code4rena.com/reports/2024-06-vultisig).

### Staking Contract Audit

The audit was done by [Zenith](https://x.com/zenith256) and can be found [here](https://github.com/zenith-security/reports/blob/main/reports/Vultisig%20-%20Zenith%20Audit%20Report.pdf).&#x20;

***

## Threshold Signature Security (TSS and DKLS)

Vultisig currently supports two Threshold Signature Schemes (TSS): GG20 and DKLS. As of March 2025, DKLS is the preferred cryptographic standard for all new vaults on iOS, Android, and macOS. Windows also fully supports DKLS, including vault participation and optional initiation via an advanced toggle. Default behavior may evolve in future releases.

**GG20** is a well-established threshold signing protocol that has been battle-tested in production through projects like [THORChain](https://thorchain.org), which uses it in open-source, adversarial environments. In Vultisig, GG20 remains available primarily for legacy vaults and advanced configuration cases.

**DKLS** (Distributed Keygen and Local Signing) is a modern threshold signature scheme developed by [Silence Laboratories](https://github.com/silence-laboratories/dkls23). It enables multiple trusted devices to sign transactions collaboratively, without ever reconstructing the full private key.

> 🔒 Both protocols eliminate the need for a seed phrase or centralized private key, offering a seedless, self-custodial experience.

In practice:
- Your devices connect only during deliberate signing sessions.
- Private key shares are never combined or stored.
- An attacker would need access to all your devices simultaneously to forge a signature.

Vultisig will continue to evolve with the latest advancements in TSS protocols to provide secure, resilient self-custody for all users.

## Research

The Threshold Signature Scheme is a relatively new area within the field of Multi-Party Computation. As such, advancements in security and efficiency are ongoing and continuously evolving.

Vultisig closely monitors these developments and is committed to adopting more secure and efficient versions as they become available. Additionally, we are actively researching new possibilities within this space.

## DKLS Threshold Signing

Vultisig leverages [DKLS threshold signatures](https://github.com/silence-laboratories/dkls23), a cryptographic scheme developed by Silence Laboratories. This protocol enables multiple trusted devices to sign a transaction together without ever reconstructing or exposing the full private key.

This technique enhances security by eliminating single points of failure—no seed phrase or complete key is ever stored on any single device. Instead, each device holds a secure share of the key, and must participate in the signing process.

> 🔒 This is how Vultisig enables secure, seedless, self-custody for your crypto.

We will link to a verified DKLS audit from Silence Labs once available.

