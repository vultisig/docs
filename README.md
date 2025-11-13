---
description: >-
  Redefining crypto wallet security for the AI Agent era: Multi-chain,
  multi-factor, and the end of seed phrases and infrastructure for the bots
cover: .gitbook/assets/Background.png
coverY: 12
---

# Overview

## Vultisig Project

🛠 Vultisig is creating and maintaining an open source cryptocurrency wallet that takes the novel approach of providing a seedless, multi-factor, multi-chain wallet infrastructure based on MPC (Multi-Party Computation) technology.

Distributed across all major platforms like iOS, Android, Windows, and Linux, users get a wallet with the enhanced security that comes from a multi-sig signing approach without the inconvenience of traditional multi-sig operations.

Vultisig operates several different solutions and is constantly expanding, such as:

* A Wallet (Vultisig Vaults)
* An Extension (Vultisig Extension)
* Webapp (View only and Airdrop)
* Marketplace (for Plugins and AI Agents)
* An SDK for builders

***

## What is Vultisig Wallet

🙋‍♀️ Vultisig's wallet solution is a multi-chain, multi-factor, multi-platform Threshold Signature Scheme (TSS) vault that does not require any specialised hardware.

Vultisig supports most UTXO, EVM, BFT and EdDSA chains (e.g., Bitcoin, Ethereum, Cosmos, Solana, etc.) by generating both ECDSA and EdDSA signature keys for every vault. This means you can hold and sign for BTC, ETH, SOL, and more from the same vault setup.

Vultisig is powered by the state-of-the-art TSS protocol, [DKLS23](https://github.com/silence-laboratories/dkls23), which is developed and maintained by Silence Laboratories.

Vultisig is built by the founders of [THORChain](https://thorchain.org), creators of the largest and longest-running multi-chain DEX powered by Threshold Signature Schemes (TSS).

## How can it be used?

Vultisig works much like a traditional multisig wallet but offers flexible setup options—including multiple devices and advanced security configurations.

#### Ease of use options:

* Fast Vaults (Single device solution, "hot wallet")

#### Self Custodial options, Secure Vaults:

* 2 of 2 (fastest setup, no fallback device\*)
* 2 of 3 (recommended, with fallback device\*)
* 3 of 4 (most secure, with fallback device\*)

{% hint style="info" %}
_These are the most widely used vault configurations, but Vultisig supports a wide range of threshold setups for advanced users._
{% endhint %}

## Why

🔮 Vultisigs are better than Single-sig or Multi-sig Wallets:

* Vultisig vault shares do not store funds - thus no funds at risk if you lose a key.
* There is no on-chain registration of signing-keys.
* Vultisigs are compatible with all chains and fully DeFi compatible (Multi-sigs are not).

📱 Vultisig uses your favourite laptops, tablets, and phones.

* Rich, large screens with familiar UI, Secure Enclaves to store vault shares and with biometric locking.
* Does not look like a hardware device and therefore grants higher security.

🌈 Vultisig is open-source, based on secure technology and is audited (see in [Security](other/security.md), _with further audits to come_).

{% hint style="info" %}
**No registration required:**

* Vultisig never permanently stores your email or personal info.
* _Fast Vaults_ briefly require your email to send the `.vult` backup file — it's used just once and not retained.
{% endhint %}

## Who should use it

* Everyone should use it!
* Vultisig is far safer than anything else (mobile, desktop, hardware, etc).

## Download & Contribute

Vultisig is an ever-expanding ecosystem, with ongoing work on new platforms, products, and features.

👩‍💻 Read the docs!

🍿 Download here:

* [iOS AppStore](https://apps.apple.com/us/app/vultisig/id6503023896)
* [iOS Testflight](https://testflight.apple.com/join/dxKWaCNK)
* [Google Play Store](https://play.google.com/store/apps/details?id=com.vultisig.wallet)
* [Windows](https://github.com/vultisig/vultisig-windows/releases)
* [Linux](https://github.com/vultisig/vultisig-windows/releases)
* [Vultisig Chrome Extension](https://chromewebstore.google.com/detail/vulticonnect/ggafhcdaplkhmmnlbfjpnnkepdfjaelb?authuser=0\&hl=en-GB)
* [Android source](https://github.com/vultisig/vultisig-android)
* [macOS source](https://github.com/vultisig/vultisig-ios)

💬 Join the community: Connect with the team and other users on [Discord](https://discord.com/invite/9wrAfNSRpS).

{% hint style="info" %}
**Fallback device:** A fallback device means that a device can be lost and access to funds is still ensured. Deleting the application is equivalent to losing the device. To properly back up the vault shares, please use the [following steps](vultisig-app-actions/managing-your-vault/vault-backup.md).
{% endhint %}
