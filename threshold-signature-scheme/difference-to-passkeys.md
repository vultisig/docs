---
description: What is the difference between TSS and Passkeys??
---

# Difference to Passkeys

## What is a passkey?

A passkey is a new way of storing sensitive data to access applications and websites, developed by an initiative of the FIDO (Fast IDentity Online) Alliance to create a new sign-in standard. \
Initially used by Apple, Microsoft and Google, it is gaining popularity as a secure replacement for passwords.

## How do passkeys work?

Passkeys consist of a private key and a public key, as we already know from the crypto industry. \
They are generated on your device and the public key is stored on the application or website. \
The private key is stored on the device and is used to authenticate the private key, usually created with biometric authentication tools. When a user wants to log into a website or authenticate themselves, the website/app sends a challenge to the user's device, which uses the private key to digitally sign the challenge with the public key credentials. \
The response proves that the user has the private key without exposing it. \
Much like cryptocurrency transactions.

## Why does Vultisig not use passkeys?

The devil is in the details, because in general passkeys are very secure and should really be adapted to replace passwords.

There are some major flaws when it comes to securing assets and wealth, which should never be compromised:

* The passkey technology itself is open source, but relies on centralized platform authentications from large cooperations.
* It is not completely sure how and what data is collected by large cooperations.
* Authentication is a single point of failure, where a physical attack can easily compromise security.

These are general weaknesses of passkeys, but when used specifically for cryptocurrencies, there are two more points:

* It is not multi-chain
* It is just another single signature technology

This is why Vultisig decided to develop our own solution as we want to bring the cryptocurrency space forward.\
We want to set new standards, which are:&#x20;

* Open source everything
* Multi-Chain&#x20;
* Multi-Factor Authentication
