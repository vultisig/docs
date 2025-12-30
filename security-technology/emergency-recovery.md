---
description: >-
  Emergency recovery tool for Vultisig vaults. Reconstruct keys if TSS
  software becomes unavailable. Last-resort fund recovery process.
---

# Emergency Recovery

{% hint style="danger" %}
**Last Resort Only**

Emergency recovery permanently converts a TSS vault to a single-signature wallet. Only use if Vultisig software is completely unavailable.
{% endhint %}

***

## What is the Emergency Recovery Tool?

In Vultisig, the private key never exists during normal operation—this significantly improves security since there's no private key to be extracted or stolen.

However, it is possible to recombine vault shares and generate the private key for use in other wallets. Scripts are provided for this purpose if Vultisig software ever becomes unavailable.

***

## Recommended: Vault Share Decoder

For most users, the community-built **Vault Share Decoder** is the easiest way to recover:

* **Website**: [https://share-decoder.vultisig.com](https://share-decoder.vultisig.com)
* **Repository**: [GitHub](https://github.com/vultisig/community-tools/tree/main/recovery-tools/vultisig-share-decoder)

This tool provides a user-friendly interface to decode vault shares and recover private keys without requiring developer tools.

{% hint style="info" %}
The Vault Share Decoder is the preferred method for emergency recovery. Only use the CLI/Web options below if the community tool is unavailable.
{% endhint %}

***

## Prerequisites

* [Go (Golang)](https://golang.org/dl/) installed on your system
* Access to threshold number of vault shares

***

## Option 1: Web UI Version

1. **Clone the Repository:**

```sh
git clone https://github.com/vultisig/mobile-tss-lib
```

2. **Navigate to the Recovery Web Directory:**

```sh
cd mobile-tss-lib/cmd/recovery-web
```

3. **Run the Web Server:**

```sh
make
```

Access the recovery UI via your web browser.

***

## Option 2: CLI Version

1. **Clone the Repository:**

```sh
git clone https://github.com/vultisig/mobile-tss-lib
```

2. **Navigate to the Recovery CLI Directory:**

```sh
cd mobile-tss-lib/cmd/recovery-cli
```

3. **Run the Recovery CLI Tool:**

```sh
go run main.go
```

Follow the terminal instructions to proceed with recovery.

***

## Supported Assets

{% hint style="danger" %}
Only the following assets are supported: Bitcoin, Bitcoin Cash, Litecoin, Dogecoin, Ethereum, THORChain, MayaChain.

**Warning:** Before using this tool, the private key never existed. This is a one-way function—once the private key is created, the vault is no longer a TSS vault but a single-signature wallet.

Never use a single-signature wallet for significant value again.
{% endhint %}

***

## Resources

* Web UI: [https://github.com/vultisig/mobile-tss-lib/tree/main/cmd/recovery-web](https://github.com/vultisig/mobile-tss-lib/tree/main/cmd/recovery-web)
* CLI: [https://github.com/vultisig/mobile-tss-lib/tree/main/cmd/cli](https://github.com/vultisig/mobile-tss-lib/tree/main/cmd/cli)

***

## Related

- [Backup & Recovery](../getting-started/backup-recovery.md) — Standard backup procedures
- [TSS Actions](tss-actions.md) — How threshold signatures work
