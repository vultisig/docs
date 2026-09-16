---
description: >-
  Vultisig Docs: seedless MPC vault. ECDSA and EdDSA from one vault. Fast Vault,
  Secure Vault, SDK, and CLI. No seed phrase.
cover: .gitbook/assets/Banner-new.png
coverY: 0
---

# Vultisig Documentation

## What is Vultisig?

Vultisig is an open-source vault with no seed phrase. It uses threshold signatures (MPC / TSS). Your key is split into shares. They sign together. One share cannot move funds.

One vault generates both ECDSA and EdDSA keys. You hold and sign Bitcoin, Ethereum, Cosmos, Solana, and 36+ other UTXO, EVM, BFT, and EdDSA chains. Named list: [Wallet tab](app-guide/wallet/).

Signing uses [DKLS23](https://github.com/silence-laboratories/dkls23) from Silence Laboratories.

* **No seed phrase.** Vault shares replace 12/24-word seeds.
* **MPC / TSS.** Threshold signatures. No complete private key in normal operation.
* **ECDSA and EdDSA.** Both key types from one vault.
* **36+ chains** from one vault.
* **Platforms:** iOS, Android, macOS, Windows, Linux, Chrome, Firefox.
* **Open source** on [GitHub](https://github.com/vultisig).

***

## How it works

* **No complete key.** Shares sign together. One lost or stolen share cannot move funds.
* **A normal signature on-chain.** No published signer set, no extra contract. That is why it works on Bitcoin and in DeFi where on-chain multisig often does not.
* **Devices you already own.** Shares sit in the device Secure Enclave, behind biometrics.
* **Open source and audited.** Code on [GitHub](https://github.com/vultisig). [Audit reports](help/security.md).

{% hint style="info" %}
**No registration.** Vultisig does not keep your email. Fast Vaults use it once to send a backup file, then discard it.
{% endhint %}

***

## Vault types

Think of the vault as a safe that needs a threshold of keys to open.

**Fast Vault** — Your device plus VultiServer (2-of-2). About a minute to set up. VultiServer signs with you and can never sign alone.

**Secure Vault** — Two or more devices you control:

* 2-of-2: both devices required. No fallback if one is gone.
* 2-of-3: the usual setup. You can lose one device and still sign.
* 3-of-4: three of four devices have to agree.

{% hint style="info" %}
Those are the usual setups. Other thresholds are available if you need them.
{% endhint %}

***

## Quick Navigation

<table data-card-size="large" data-view="cards"><thead><tr><th></th><th></th><th data-hidden data-card-target data-type="content-ref"></th></tr></thead><tbody><tr><td><strong>Getting Started</strong></td><td>New here? Download, create a vault, back it up, then send.</td><td><a href="getting-started/overview.md">overview.md</a></td></tr><tr><td><strong>App Guide</strong></td><td>Wallet, DeFi, vault management, Settings.</td><td><a href="app-guide/overview.md">overview.md</a></td></tr><tr><td><strong>Security &#x26; Technology</strong></td><td>How threshold signatures work, and the signing protocols.</td><td><a href="security-and-technology/overview.md">overview.md</a></td></tr><tr><td><strong>Developer Docs</strong></td><td>SDK, CLI, and Extension integration.</td><td><a href="developer-docs/">developer-docs</a></td></tr></tbody></table>

***

## Vultisig Ecosystem

<table data-card-size="large" data-view="cards"><thead><tr><th></th><th></th><th data-hidden data-card-target data-type="content-ref"></th></tr></thead><tbody><tr><td><strong>Vultisig Wallet</strong></td><td>iOS, Android, macOS, Windows, Linux</td><td><a href="getting-started/download-install.md">download-install.md</a></td></tr><tr><td><strong>Vultisig Extension</strong></td><td>Full vault in Chrome and Firefox, plus dApp connect</td><td><a href="vultisig-ecosystem/vultisig-extension/">vultisig-extension</a></td></tr><tr><td>Vultisig SDK</td><td>The modular SDK for developers and agents</td><td></td></tr></tbody></table>

***

## Download

<table data-view="cards"><thead><tr><th></th><th data-hidden data-card-target data-type="content-ref"></th></tr></thead><tbody><tr><td><strong>iOS</strong></td><td><a href="https://apps.apple.com/us/app/vultisig/id6503023896">https://apps.apple.com/us/app/vultisig/id6503023896</a></td></tr><tr><td><strong>Android</strong></td><td><a href="https://play.google.com/store/apps/details?id=com.vultisig.wallet">https://play.google.com/store/apps/details?id=com.vultisig.wallet</a></td></tr><tr><td><strong>macOS</strong></td><td><a href="https://github.com/vultisig/vultisig-ios">https://github.com/vultisig/vultisig-ios</a></td></tr><tr><td><strong>Windows/Linux</strong></td><td><a href="https://github.com/vultisig/vultisig-windows/releases">https://github.com/vultisig/vultisig-windows/releases</a></td></tr><tr><td><strong>Chrome Extension</strong></td><td><a href="https://chromewebstore.google.com/detail/vulticonnect/ggafhcdaplkhmmnlbfjpnnkepdfjaelb">https://chromewebstore.google.com/detail/vulticonnect/ggafhcdaplkhmmnlbfjpnnkepdfjaelb</a></td></tr><tr><td><strong>Firefox Extension</strong></td><td><a href="https://addons.mozilla.org/en-US/firefox/addon/vultisig-extension/">https://addons.mozilla.org/en-US/firefox/addon/vultisig-extension/</a></td></tr></tbody></table>

***

## Community

* **Discord**: [discord.gg/Cugw9T2NrP](https://discord.gg/Cugw9T2NrP)
* **Twitter/X**: [@vultisig](https://x.com/vultisig)
* **GitHub**: [github.com/vultisig](https://github.com/vultisig)
