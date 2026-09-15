---
description: >-
  Vultisig docs: seedless multi-chain vault, Fast Vault and Secure Vault, SDK
  and CLI. No seed phrase.
cover: .gitbook/assets/Banner-new.png
coverY: 0
---

# Vultisig Documentation

## What is Vultisig?

Vultisig is an open-source seedless vault. Threshold signatures (MPC) give multi-sig security without on-chain multi-sig, and without specialized hardware.

One vault generates both ECDSA and EdDSA keys, so you can hold and sign Bitcoin, Ethereum, Cosmos, Solana, and 36+ other UTXO, EVM, BFT, and EdDSA chains.

Signing uses [DKLS23](https://github.com/silence-laboratories/dkls23) from Silence Laboratories.

### Key Features

* **No seed phrases**: Vault shares replace vulnerable 12/24 word seeds
* **Multi-chain**: 36+ blockchains from one vault (Bitcoin, Ethereum, Solana, Cosmos, and more)
* **Multi-factor**: Multiple devices required to sign transactions
* **Multi-platform**: iOS, Android, macOS, Windows, Linux, Browser Extension
* **Open source**: Fully auditable code on [GitHub](https://github.com/vultisig)

***

## Why Vultisig?

**No complete key.** Shares sign together. One lost or stolen share cannot move funds.

**A normal signature on-chain.** No published signer set, no extra contract. Works on Bitcoin and in DeFi where multisig often does not.

**Devices you already own.** Shares sit in the device Secure Enclave, behind biometrics. Nothing that looks like a hardware wallet.

**Open source and audited.** Code on [GitHub](https://github.com/vultisig). [Audit reports](help/security.md).

{% hint style="info" %}
**No registration required.** Vultisig never permanently stores your email or personal info. Fast Vaults use email once to send a backup file, then discard it.
{% endhint %}

***

## How It Works

Vultisig works like a traditional multisig wallet but with flexible setup options and modern convenience and rich UI.

**Fast Vault** — Single device + Vultisig server. Quick setup, instant signing.

**Secure Vault** — Multiple devices you control:

* 2-of-2: Fastest setup, no fallback
* 2-of-3: Recommended, with fallback device
* 3-of-4: Maximum security, with fallback device

{% hint style="info" %}
These are the most common configurations, but Vultisig supports a wide range of threshold setups for advanced users.
{% endhint %}

***

## Quick Navigation

<table data-card-size="large" data-view="cards"><thead><tr><th></th><th></th><th data-hidden data-card-target data-type="content-ref"></th></tr></thead><tbody><tr><td><strong>Getting Started</strong></td><td>New to Vultisig? Download, create your first vault, and make your first transaction.</td><td><a href="getting-started/overview.md">overview.md</a></td></tr><tr><td><strong>App Guide</strong></td><td>Complete guides for Wallet, DeFi, and Vault Management features.</td><td><a href="app-guide/overview.md">overview.md</a></td></tr><tr><td><strong>Security &#x26; Technology</strong></td><td>How threshold signatures work, protocol deep-dives, and security architecture.</td><td><a href="security-and-technology/overview.md">overview.md</a></td></tr><tr><td><strong>Developer Docs</strong></td><td>Build on Vultisig: SDK, CLI, and Extension integration.</td><td><a href="developer-docs/">README.md</a></td></tr></tbody></table>

***

## Vultisig Ecosystem

<table data-card-size="large" data-view="cards"><thead><tr><th></th><th></th><th data-hidden data-card-target data-type="content-ref"></th></tr></thead><tbody><tr><td><strong>Vultisig Wallet</strong></td><td>Core vault application for iOS, Android, macOS, Windows, Linux</td><td><a href="getting-started/download-install.md">download-install.md</a></td></tr><tr><td><strong>Vultisig Extension</strong></td><td>Browser extension for Web3 dApp connections</td><td><a href="vultisig-ecosystem/vultisig-extension/">vultisig-extension</a></td></tr><tr><td><strong>Web App</strong></td><td>Public web view (airdrop importer; host not currently serving)</td><td><a href="vultisig-ecosystem/web-app.md">web-app.md</a></td></tr><tr><td><strong>Marketplace</strong></td><td>Plugin marketplace (paused, not in production)</td><td><a href="vultisig-ecosystem/marketplace.md">marketplace.md</a></td></tr></tbody></table>

***

## Download

<table data-view="cards"><thead><tr><th></th><th data-hidden data-card-target data-type="content-ref"></th></tr></thead><tbody><tr><td><strong>iOS</strong></td><td><a href="https://apps.apple.com/us/app/vultisig/id6503023896">https://apps.apple.com/us/app/vultisig/id6503023896</a></td></tr><tr><td><strong>Android</strong></td><td><a href="https://play.google.com/store/apps/details?id=com.vultisig.wallet">https://play.google.com/store/apps/details?id=com.vultisig.wallet</a></td></tr><tr><td><strong>macOS</strong></td><td><a href="https://github.com/vultisig/vultisig-ios">https://github.com/vultisig/vultisig-ios</a></td></tr><tr><td><strong>Windows/Linux</strong></td><td><a href="https://github.com/vultisig/vultisig-windows/releases">https://github.com/vultisig/vultisig-windows/releases</a></td></tr><tr><td><strong>Browser Extension</strong></td><td><a href="https://chromewebstore.google.com/detail/vulticonnect/ggafhcdaplkhmmnlbfjpnnkepdfjaelb">https://chromewebstore.google.com/detail/vulticonnect/ggafhcdaplkhmmnlbfjpnnkepdfjaelb</a></td></tr></tbody></table>

***

## Community

* **Discord**: [discord.gg/Cugw9T2NrP](https://discord.gg/Cugw9T2NrP)
* **Twitter/X**: [@vultisig](https://x.com/vultisig)
* **GitHub**: [github.com/vultisig](https://github.com/vultisig)
