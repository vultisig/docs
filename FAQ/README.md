---
description: Vultisig Frequently Asked Questions
---

# Frequently Asked Questions

### What are the biggest differences between Vultisig and all the other cold and hot wallets out there?

All the other cold and hot wallets are single-sig, and they rely on one set of seed phrase and one piece of private key; while Vultisig is a Multi-Sig Multi-Device wallet that does not rely on seed phrases and no single-point-of-failure to enhance wallet security. Vultisig utilizes Threshold Signature Scheme (TSS): [Reference](https://medium.com/zengo/threshold-signatures-private-key-the-next-generation-f27b30793b).

### What chains can/will Vultisig support?

Any chains on [Trust Wallet Core](https://github.com/trustwallet/wallet-core/tree/master/src) can be supported.

### Will it be possible to add tokens like ERC-20s on EVM blockchains, or SPL tokens on Solana, etc?

Yes.

### Will Vultisig support web apps or browser extension?

Yes, with partners, utilizing Vultisigner.

### Will it be possible to integrate Vultisig with Li-Fi (JumperExchange), and different platforms like Uniswap and others?

Yes, via Vultisigner.&#x20;

### What versions of iOS / iPhones are supported?

At least iOS 17; iPhone XS (2018) or newer.

### What versions of Android are supported?

At least API Level 26; Android 8.0 (Oreo) or newer.

### Are the Vault Shares automatically uploaded into iCloud?

No. Users need to [manually backup the Vault Shares](https://docs.vultisig.com/user-actions/managing-your-vault). Users can then choose to keep the Vault Share files on iCloud.

{% hint style="warning" %}
**Do not store multiple Vault Share files in a single location. Anyone who have m Vault Shares (of a m-of-n setup) will have full access to the Vault**
{% endhint %}

### Can users mix-and-match iOS and Android?

Yes! Vault shares & signers are device agnostic. Users can mix-and-match any devices (iOS, Android, browser apps, etc.) to be used as signers.

### Can users import their existing address into Vultisig?

The single-sig private key and seed phrase of existing addresses are single-sig wallets, and cannot be imported into a Vultisig - which is multi-sig wallet.&#x20;

### Can we bond RUNE/ provide liquidity on THORChain by using Vultisig and will it be counted for airdrop?

Yes, Vultisig will support this. It will be counted to airdrop value.&#x20;

### What is the best practice in the event of losing 67% of vault shares at the same time? For example, a lady has her handbag snatched with two mobile phones inside the handbag and she has Vultisig 2of2 or 2of3 vault setup with both mobile phones.

*Scenario 1 :*
With mobile remote data wipe feature setup done and with at least 67% of vault shares backup done properly

1. Erase mobile phones data remotely. 
2. Buy new devices and import vault shares backup to new devices to regain access to the vault.

*Scenario 2 :*
WITHOUT mobile remote data wipe feature setup done and with at least 67% of vault shares backup done properly

1. Buy new devices and import vault shares backup to new devices to regain access to the vault.
2. Immediately move your funds away from this existing vault.

{% hint style="warning" %}
**You would still loss access to your funds if you do not have at least 67% of vault shares backup done properly.**
{% endhint %}

### What is the key utility of $VULT token to the ordinary users?

All the platform incomes (affiliate, router and bridge fees) will be used to buy $VULT and burn to reduce $VULT supply, and create an deflationary effect.

The larger amount and the longer time a user uses the platform, the more $VULT airdrops the user would receive.
