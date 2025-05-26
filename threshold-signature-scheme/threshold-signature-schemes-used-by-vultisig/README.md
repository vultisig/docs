---
description: 'Understanding Vultisig’s TSS Protocols: GG20 and DKLS23'
---

# Threshold Signature Schemes used by Vultisig

What is TSS and Why Does It Matter?
-----------------------------------

The Threshold Signature Scheme (TSS) protects your digital assets by splitting control among multiple devices or parties using multi-party computation (MPC) technology. Rather than using one private key, which could be stolen, TSS divides control so that a minimum number of shares `t`("threshold") must work together to authorize transactions.

It's like a bank vault that requires three out of five keys to open rather than one master key. If one key is lost or stolen, your assets remain safe because the thief would need additional keys.

{% hint style="warning" %}
Note: If you lose access to enough shares to meet the threshold, you will permanently lose access to your assets. **Backing up each device individually is essential.**
{% endhint %}

For your wallet security, TSS means:

* No single point of failure
* Protection against device loss or theft
* The ability to customize your security setup

***

## The Two Protocols used by Vultisig: GG20 and DKLS23

Vultisig utilizes two TSS protocols: the original GG20 and the newer DKLS23. \
Both protocols secure your assets, but they do so in different technical ways.

### &#xD;[GG20](how-it-works.md): The Protocol Vultisig started with

GG20 (developed in 2020) was Vultisig’s first TSS protocol. It’s like a secure but somewhat&#x20;complex postal system:

* **How it works:** Uses a special type of encryption called “Paillier encryption” that allows  &#x20;mathematical operations on encrypted data
* **Communication style:** Requires 6 rounds of back-and-forth messages between your  &#x20;devices to complete a signature
* **Security approach:** Includes “identifiable abort” - the ability to detect which device might be  &#x20;causing problems



***

[DKLS23](how-dkls23-works.md): The Enhanced Protocol


DKLS23 (developed in 2023) is Vultisig’s current, newer protocol. I\
t’s like a streamlined, modern courier&#x20;service:

* **How it works:** Uses a technique called “oblivious transfer” that’s more efficient than the  &#x20;previous method
* **Communication style:** Requires only 3 rounds of messages (half as many as GG20)
* **Security approach:** Maintains the same strong security but with simpler mechanisms

***

Why Vultisig Upgraded to DKLS23
-------------------------------

In early 2025, Vultisig transitioned from GG20 to DKLS23 to provide users with a better experience while maintaining the highest security standards. This was accomplished in close cooperation with [Silence Laboratories](https://x.com/silencelabs_sl), which developed a customized version of DKLS23 for Vultisig. The open-source protocol can be found on the Silence Laboratories [GitHub](https://github.com/silence-laboratories/dkls23).

The upgrade offers:

1. Faster transactions: Signing operations complete in milliseconds rather than seconds
2. Better reliability: Fewer communication rounds mean less chance of failure when internet   \
   connections are slow
3. Improved device compatibility: Works efficiently even on smartphones with limited   \
   processing power
4. Enhanced battery life: Less computational work means less drain on your device’s battery
5. More responsive experience: Security operations happen in the background without   \
   noticeable delays
6. WASM Compatibility: Provides the Vultisig Extension with the same experience as the mobile apps.

This upgrade represents Vultisig’s commitment to implementing the latest advancements in\
cryptographic security while improving the user experience.

***

Comparing the Protocols


| Feature                          | GG20                               | DKLS23       |    What This Means For The User                     |
| -------------------------------- | ---------------------------------- | ------------ | --------------------------------------------------- |
| Signing Speed                    | Slower                             | 5-10x faster | Quicker transaction approvals                       |
| <p>Communication<br>Rounds</p> | 6 round                            | 3 rounds     | <p>Works better on<br>spotty connections</p>       |
| Security Level                   | Very High                          | Very High    | Both protocols provide&#xD; excellent security&#xD; |
| Network Reliability              | More sensitive to&#xD; delays&#xD; | More robust  | Fewer failed transactions                           |

***

## What This Means For Your Wallet

As a Vultisig user, the transition to DKLS23 happens either with creating a fresh vault or migrating your already existing GG20 vault in the settings. You’ll notice:

* Faster transaction signing
* More reliable operation, especially on mobile networks
* Better performance on all your devices
* The same high level of security you expect
* Possibility to creat Plugins and AI Agent marketplace

Do I need to migrate to DKL23?


While the possible improvements for end users of Vultisig are clear, upgrading the existing GG20 vaults to DKLS23 vaults is unnecessary. Vultisig will continue to support both protocols up to this date.

## Learn More

If you’re interested in the technical details of these protocols, you can explore the more&#x20;comprehensive explanation of [GG20](how-it-works.md) or [DKLS23](how-dkls23-works.md)\


{% hint style="info" %}
This document provides a simplified overview of complex cryptographic protocols. While this is a simplified the concepts for clarity, both protocols provide state-of-the-art security for your digital&#x20;assets.
{% endhint %}
