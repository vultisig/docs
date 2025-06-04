---
description: Vultisig Frequently Asked Questions
---

# Frequently Asked Questions

### What are the biggest differences between Vultisig and all the other cold and hot wallets out there?

All other cold and hot wallets are single-sig and rely on one set of seed phrase and one piece of private key, while Vultisig is a Multi-Sig Multi-Device wallet that does not rely on seed phrases and has no single point of failure, enhancing wallet security. Vultisig utilizes the Threshold Signature Scheme (TSS): [Reference](https://medium.com/zengo/threshold-signatures-private-key-the-next-generation-f27b30793b).

### What Platforms are supported and which are planned?

* Available: iOS, Android, Windows and Web
* Planned: Linux

### What versions of iOS/IPadOS/Mac are supported?

**iOS:** Minimum iOS 17 with iPhone XS (2018)

**iPad:** Minimum iPad OS 17.2

* iPad Pro 12.9-inch (2nd generation and later)
* iPad Pro 10.5-inch
* iPad Pro 11-inch (1st generation and later)
* iPad Air (3rd generation and later)
* iPad (6th generation and later)
* iPad mini (5th generation and later)

**Mac:** Minimum MacOS 14.0

### What versions of Android are supported?

Minimum API Level 26; Android 8.0 (Oreo) or newer.

### Can users mix and match iOS and Android devices?

Yes! Vault shares and signers are device agnostic. Users can mix and match any devices (iOS, Android, browser apps, etc.) to be used as signers.

### What chains are supported by Vultisig?

Currently supported are:

* Arbitrum
* Avalanche
* BSC
* Base
* Bitcoin
* Bitcoin-Cash
* Blast
* Cosmos
* CronosChain
* Dash
* Dogecoin
* dydx
* Ethereum
* Kujira
* Litecoin
* Maya Protocol
* Optimisim
* Polkadot
* Polygon
* Solana
* Sui
* THORChain
* Ton
* Zksync

Generally any chains on [Trust Wallet Core](https://github.com/trustwallet/wallet-core/tree/master/src) can be supported.

### Is it possible to add tokens like ERC-20s on EVM blockchains, or SPL tokens on Solana, etc?

Yes, most ERC-20 and SPL tokens are available. More and more tokens on different chains will become available gradually.

### Will Vultisig support web apps and different platforms like Uniswap or Li-Fi(JumperExchange)?

Yes, most web apps can be accessed with the VultiConnect browser extension. Some platforms will need a dedicated integration with Vultisig.

### Are the Vault Shares automatically uploaded into iCloud?

No. Users need to [manually backup the Vault Shares](https://docs.vultisig.com/user-actions/managing-your-vault). Users can then choose to keep the Vault Share files on iCloud.

{% hint style="warning" %}
**Do not store multiple Vault Share files in a single location. Anyone who has m Vault Shares (of a `m`-of-`n` setup) will have full access to the Vault**
{% endhint %}

### Can users import their existing address into Vultisig?

The single-sig private key and seed phrase of existing addresses are single-sig wallets, and cannot be imported into a Vultisig - which is mpc wallet.\
This is to maintain and improve the security of the vaults, as single-sig seed phrases may have been previously exposed or had a bad generation, making them more vulnerable.

### Can we bond RUNE/ provide liquidity on THORChain/MAYA by using Vultisig and will it be counted for airdrop?

Yes, Vultisig supports MsgDeposits for THORChain and MAYA Protocol.\
Saver, Liquidity Positions, Lending and Bonding will be counted for the airdrop

### What is the best practice in the event of losing 67% of vault shares at the same time?

For example, a lady has her handbag snatched with two mobile phones inside the handbag and she has Vultisig 2of2 or 2of3 vault setup with both mobile phones

_Scenario 1 :_ With mobile remote data wipe feature setup done and with at least 67% of vault shares backup done properly

1. Erase mobile phones data remotely.
2. Buy new devices and import vault shares backup to new devices to regain access to the vault.

_Scenario 2 :_ WITHOUT mobile remote data wipe feature setup done and with at least 67% of vault shares backup done properly

1. Buy new devices and import vault shares backup to new devices to regain access to the vault.
2. Immediately move your funds away from this existing vault.

{% hint style="warning" %}
**You would still lose access to your funds if you do not have at least 67% of vault shares backup done properly.**\
**ALWAYS BACK UP EACH DEVICE INDIVIDUALLY**
{% endhint %}

### What is the key utility of $VULT token to the ordinary users?

All the platform incomes (affiliate, router and bridge fees) will be used to buy $VULT and burn to reduce $VULT supply, and create an deflationary effect.

The larger amount and the longer time a user uses the platform, the more $VULT airdrops the user would receive.
