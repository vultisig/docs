---
description: >-
  Vultisig TypeScript SDK and CLI. Create Fast or Secure vaults, send, swap, and
  sign across 36+ chains. No seed phrase.
---

# Vultisig SDK

The self-custodial multi-chain wallet SDK for developers and AI agents.

* **Send and swap** across 36+ blockchains with human-readable amounts (`vault.send({ amount: "0.1" })`)
* **MPC** — keys are split across parties. No seed phrase. No single point of failure
* **Cross-chain swaps** via THORChain, Maya, 1inch, Kyber, LiFi, Jupiter, and others. The SDK picks the quote
* **Portfolio** with balances and fiat prices
* **Dry-run** on send and swap — preview fees and output before signing
* **Agent-ready** — JSON output, programmatic API

Install `@vultisig/sdk` for apps and bots. Install `@vultisig/cli` for shells and coding agents.

Source: [vultisig/vultisig-sdk](https://github.com/vultisig/vultisig-sdk)

## Which interface

|                    | Programmatic SDK (`@vultisig/sdk`)           | CLI (`@vultisig/cli`)                       |
| ------------------ | -------------------------------------------- | ------------------------------------------- |
| **Install**        | `npm install @vultisig/sdk`                  | `npm install -g @vultisig/cli`              |
| **Best for**       | Apps, bots, long-running services            | Coding agents, shell scripts                |
| **Language**       | TypeScript                                   | Shell, JSON out                             |
| **Vault creation** | `sdk.createFastVault()` + email verification | `vultisig create fast --two-step`           |
| **Agent mode**     | `vault.send()`, `vault.swap()`               | `vultisig agent ask "check my ETH balance"` |
| **Long-running**   | Initialize WASM once                         | Each command reloads WASM                   |
| **Signing**        | `onPasswordRequired` or `vault.unlock()`     | `--password` or `VAULT_PASSWORD`            |

Full walkthrough: [SDK implementation guide](sdk-users-guide.md). Command reference: [CLI](cli.md).

## Vault types

* **Fast Vault** — 2-of-2 with Vultiserver. Instant signing.
* **Secure Vault** — N-of-M on devices you control. Other devices join by QR with the Vultisig app.

Both types derive addresses and sign on the same 36+ chains as the apps.

## Installation

```bash
npm install @vultisig/sdk
```

```bash
npm install -g @vultisig/cli
# or
npx @vultisig/cli --help
```

Related: [`@vultisig/rujira`](https://github.com/vultisig/vultisig-sdk/tree/main/packages/rujira) for Rujira (FIN) helpers on THORChain.

## Quick start

### Fast Vault

```typescript
import { Vultisig, Chain } from '@vultisig/sdk'

const sdk = new Vultisig()
await sdk.initialize()

const vaultId = await sdk.createFastVault({
  name: 'My Wallet',
  email: 'user@example.com',
  password: 'secure-password',
})

const vault = await sdk.verifyVault(vaultId, '1234')

const ethAddress = await vault.address(Chain.Ethereum)

const result = await vault.send({ chain: Chain.Ethereum, to: '0x...', amount: '0.1' })

const { signature } = await vault.signMessage('Login to MyDapp')

await vault.swap({
  fromChain: Chain.Ethereum,
  fromSymbol: 'ETH',
  toChain: Chain.Bitcoin,
  toSymbol: 'BTC',
  amount: '0.5',
})

const portfolio = await vault.portfolio('usd')
```

Agents and CI can skip the email step:

```typescript
const vault = await sdk.createFastVault({
  name: 'My Wallet',
  email: 'user@example.com',
  password: 'secure-password',
  skipVerification: true,
})
```

### Secure Vault

```typescript
const { vault } = await sdk.createSecureVault({
  name: 'Team Wallet',
  devices: 3,
  onQRCodeReady: (qrPayload) => {
    displayQRCode(qrPayload)
  },
  onDeviceJoined: (deviceId, total, required) => {
    console.log(`Device joined: ${total}/${required}`)
  },
})

await vault.sign(payload, {
  onQRCodeReady: (qr) => displayQRCode(qr),
  onDeviceJoined: (id, total, required) => {
    console.log(`Signing: ${total}/${required} devices ready`)
  },
})
```

{% hint style="danger" %}
Do not use `MemoryStorage` in production. It is not persistent. Vault shares are gone when the process exits. The SDK picks FileStorage (Node/Electron) or BrowserStorage (IndexedDB) by default. Always `vault.export()` a backup.
{% endhint %}

## Send, swap, portfolio

```typescript
await vault.send({ chain: Chain.Ethereum, to: '0x...', amount: '0.1' })
await vault.send({ chain: Chain.Ethereum, to: '0x...', amount: '0.1', dryRun: true })

await vault.swap({
  fromChain: Chain.Ethereum,
  fromSymbol: 'ETH',
  toChain: Chain.Ethereum,
  toSymbol: 'USDC',
  amount: '0.5',
  dryRun: true,
})

await vault.portfolio('usd')
```

| Need                   | Method                                   |
| ---------------------- | ---------------------------------------- |
| One chain, native coin | `vault.balance(chain)`                   |
| One chain, token       | `vault.balance(chain, tokenId)`          |
| Several chains         | `vault.balances(chains, includeTokens?)` |
| All configured chains  | `vault.allBalances(includeTokens?)`      |
| Fiat portfolio         | `vault.portfolio('usd')`                 |

Omit `symbol` on send for the native coin. Set it for tokens (`"USDC"`).

## Platforms

Node.js 20+, Electron 20+, and browsers with WebAssembly (Chrome, Firefox, Safari, Edge). React Native is built as `@vultisig/sdk/react-native`.

WASM (DKLS, Schnorr, Wallet Core) loads from the npm package. For Vite, use `@vultisig/sdk/vite`. Details are in the [implementation guide](sdk-users-guide.md#browser-setup-vite-and-wasm).

## Next

* [SDK implementation guide](sdk-users-guide.md) — passwords, vault lifecycle, swaps, events, platform notes
* [CLI](cli.md) — commands, JSON, `agent ask`, exit codes
* [Examples](https://github.com/vultisig/vultisig-sdk/tree/main/examples/browser) — browser app
* [MIGRATING.md](https://github.com/vultisig/vultisig-sdk/blob/main/MIGRATING.md) — breaking changes
* [Issues](https://github.com/vultisig/vultisig-sdk/issues)
* [Discord](https://discord.gg/Cugw9T2NrP)
