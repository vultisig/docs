---
description: >-
  Emergency recovery for Vultisig vaults with the Vault Share Decoder.
  Reconstruct private keys from a threshold of vault shares (DKLS and GG20)
  if Vultisig software becomes unavailable. Last-resort fund recovery.
---

# Emergency recovery in Vultisig

{% hint style="danger" %}
**Last Resort Only**

Emergency recovery permanently converts a TSS vault to a single-signature wallet. Only use if Vultisig software is completely unavailable.
{% endhint %}

***

## What is emergency recovery?

In Vultisig, the private key never exists during normal operation—this significantly improves security since there's no private key to be extracted or stolen.

However, it is possible to recombine vault shares and generate the private key for use in other wallets. The **Vault Share Decoder** does this if Vultisig software ever becomes unavailable.

***

## Vault Share Decoder

* **Website**: [https://share-decoder.vultisig.com](https://share-decoder.vultisig.com)
* **Repository**: [GitHub](https://github.com/vultisig/community-tools/tree/main/recovery-tools/vultisig-share-decoder)

It supports both vault types: **DKLS** (current vaults) and **GG20** (older vaults). The scheme is detected automatically.

{% hint style="info" %}
Shares are decoded in the browser and never leave your device. For real funds, run the decoder locally (see below) on an offline machine rather than using the hosted website.
{% endhint %}

***

## What you need

* **A threshold of vault shares** — for example, 2 shares for a 2-of-3 vault. Each share is the `.vult` (or older `.bak`) backup file from a different device.
* **The backup password for each share**, if the backups are encrypted. Each backup may have its own password.
* **Your vault's addresses**, noted from the Vultisig app while it still works. Use them to confirm the recovered keys are correct.

***

## Recover with the website

1. Open [share-decoder.vultisig.com](https://share-decoder.vultisig.com), or your local copy.
2. Upload a threshold of vault shares and enter the password for any encrypted share.
3. Run the recovery and note the private keys and addresses for each chain.
4. **Check that every recovered address matches the address in your Vultisig app.** If an address does not match, the recovery is wrong. Do not send funds to it. Retry with the correct shares.

{% hint style="warning" %}
With fewer shares than the threshold, older (GG20) vaults can still show keys and addresses without an error. They are wrong. Always upload a full threshold of shares and compare the addresses.
{% endhint %}

***

## Run the decoder locally

Build from source. Requires [Go 1.25+](https://go.dev/doc/install) with CGo.

```sh
git clone https://github.com/vultisig/community-tools.git
cd community-tools/recovery-tools/vultisig-share-decoder
```

**Web interface:**

```sh
make webserver
./dist/webserver
```

Then open [http://localhost:5000](http://localhost:5000). Start the server from the `vultisig-share-decoder` folder.

**Command line:**

```sh
make cli
./dist/cli recover --files share1.vult --files share2.vult --password <password1> --password <password2>
```

Omit `--password` for unencrypted shares. Add `--format json` for machine-readable output.

***

## Supported chains

* **UTXO:** Bitcoin, Bitcoin Cash, Dash, Dogecoin, Litecoin
* **Cosmos/BFT:** THORChain, MayaChain, Cosmos Hub, Kujira, dYdX, Terra Classic, Terra
* **EVM:** Ethereum
* **Tron:** TRX
* **XRP Ledger:** XRP
* **EdDSA:** Solana, Sui, TON

Tokens are recovered with their chain's key. For example, ERC-20 tokens use the Ethereum key. Other EVM chains, such as Arbitrum, Base, BNB Chain and Polygon, also use the Ethereum key and address. Osmosis, Noble and Akash use the Cosmos Hub key.

{% hint style="warning" %}
Chains not listed above, such as Cardano, Polkadot, Bittensor, and Zcash, are not yet supported by the decoder.
{% endhint %}

How to import each key type into another wallet is described in the [repository README](https://github.com/vultisig/community-tools/tree/main/recovery-tools/vultisig-share-decoder#importing-results).

{% hint style="danger" %}
**Warning:** Before using this tool, the private key never existed. This is a one-way function—once the private key is created, the vault is no longer a TSS vault but a single-signature wallet.

Move funds to a new vault after recovery. Never use a single-signature wallet for significant value again.
{% endhint %}

***

## Legacy recovery tools

The older recovery tools in [mobile-tss-lib](https://github.com/vultisig/mobile-tss-lib) only read GG20 vaults and cover fewer chains. Use the Vault Share Decoder instead.

***

## Related

- [Backup & Recovery](../getting-started/backup-recovery.md) — Standard backup procedures
- [TSS Actions](tss-actions.md) — How threshold signatures work
