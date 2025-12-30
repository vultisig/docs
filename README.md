---
description: >-
  Crypto security reimagined for the AI Agent era: Self-custodial automation,
  multi-chain, multi-factor, and the end of seed phrases. Your ultimate wallet.
cover: .gitbook/assets/Banner-new.png
coverY: 0
---

# Overview

{% hint style="danger" %}
**Backup Your Vault Shares**

Your vault shares are the ONLY way to recover your funds. Vultisig cannot help you recover lost shares. Export and securely store backups from each device in your vault configuration.

[Learn how to backup →](getting-started/backup-recovery.md)
{% endhint %}

***

## Quick Navigation

<table data-card-size="large" data-view="cards">
<thead><tr><th></th><th></th><th data-hidden data-card-target data-type="content-ref"></th></tr></thead>
<tbody>
<tr><td><strong>Getting Started</strong></td><td>New to Vultisig? Start here. Download, create your first vault, and secure your backup.</td><td><a href="getting-started/README.md">getting-started/README.md</a></td></tr>
<tr><td><strong>App Guide</strong></td><td>Step-by-step guides for Wallet, DeFi, and Vault Management features.</td><td><a href="app-guide/README.md">app-guide/README.md</a></td></tr>
<tr><td><strong>Security & Technology</strong></td><td>How threshold signatures work, protocol details, and security architecture.</td><td><a href="security-technology/README.md">security-technology/README.md</a></td></tr>
<tr><td><strong>FAQ</strong></td><td>Common questions about vaults, security, and troubleshooting.</td><td><a href="help/faq.md">help/faq.md</a></td></tr>
</tbody>
</table>

***

## What is Vultisig?

Vultisig is a multi-chain, multi-factor, multi-platform Threshold Signature Scheme (TSS) vault that does not require any specialized hardware.

Unlike traditional wallets with single private keys, Vultisig distributes security across multiple devices. No single point of failure—compromise one device, funds remain safe.

### Key Features

- **No seed phrases**: Vault shares replace vulnerable 12/24 word seeds
- **Multi-chain**: 30+ blockchains from one vault (Bitcoin, Ethereum, Solana, Cosmos, and more)
- **Multi-factor**: Multiple devices required to sign transactions
- **Multi-platform**: iOS, Android, macOS, Windows, Linux, Browser Extension
- **Open source**: Fully auditable code on [GitHub](https://github.com/vultisig)

***

## Vault Types

### Fast Vault

Single device paired with Vultisig's server (Vultiserver). Quick setup for everyday use.

- Instant transactions
- Single device convenience
- Server provides second signature

### Secure Vault

Multiple physical devices you control. Maximum security for significant holdings.

| Configuration | Devices | Redundancy |
|--------------|---------|------------|
| 2 of 2 | 2 required | No fallback |
| 2 of 3 | 2 of 3 required | 1 device backup |
| 3 of 4 | 3 of 4 required | 1 device backup |

{% hint style="info" %}
**Fallback device**: With 2-of-3 or 3-of-4 configurations, you can lose one device and still access funds. Always maintain backups regardless of configuration.
{% endhint %}

***

## Why Vultisig?

**Better than single-signature wallets:**
- No single private key to steal
- No seed phrase to compromise
- Multi-device authentication

**Better than traditional multi-sig:**
- Works on all chains (multi-sig often doesn't)
- Single signature on-chain (lower fees)
- Flexible device management

**Uses your existing devices:**
- Phones, tablets, laptops you already own
- Secure Enclaves protect vault shares
- Biometric authentication

***

## Vultisig Ecosystem

Vultisig operates several interconnected products:

- **Vultisig Wallet** — Core vault application
- **Vultisig Extension** — Browser extension for Web3 dApps
- **Web App** — View-only access and airdrop tracking
- **App Store** — Self-custodial automation apps
- **SDK** — Build on Vultisig infrastructure

***

## Download

<table data-view="cards">
<thead><tr><th></th><th data-hidden data-card-target data-type="content-ref"></th></tr></thead>
<tbody>
<tr><td><strong>iOS</strong></td><td><a href="https://apps.apple.com/us/app/vultisig/id6503023896">App Store</a></td></tr>
<tr><td><strong>Android</strong></td><td><a href="https://play.google.com/store/apps/details?id=com.vultisig.wallet">Google Play</a></td></tr>
<tr><td><strong>macOS</strong></td><td><a href="https://github.com/vultisig/vultisig-ios">GitHub</a></td></tr>
<tr><td><strong>Windows/Linux</strong></td><td><a href="https://github.com/vultisig/vultisig-windows/releases">GitHub</a></td></tr>
<tr><td><strong>Browser Extension</strong></td><td><a href="https://chromewebstore.google.com/detail/vulticonnect/ggafhcdaplkhmmnlbfjpnnkepdfjaelb">Chrome Web Store</a></td></tr>
</tbody>
</table>

***

## Community

- **Discord**: [discord.vultisig.com](https://discord.vultisig.com)
- **Twitter/X**: [@vaboraxi](https://x.com/vaboraxi)
- **GitHub**: [github.com/vultisig](https://github.com/vultisig)

***

## Built By

Vultisig is built by the founders of [THORChain](https://thorchain.org), creators of the largest and longest-running multi-chain DEX powered by Threshold Signature Schemes.

Powered by [DKLS23](https://github.com/silence-laboratories/dkls23), a state-of-the-art TSS protocol developed by Silence Laboratories.

{% hint style="info" %}
**No registration required**: Vultisig never permanently stores your email or personal info. Fast Vaults briefly require email to send backup files—used once and not retained.
{% endhint %}
