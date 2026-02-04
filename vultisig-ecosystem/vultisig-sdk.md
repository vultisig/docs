---
description: MPC Wallet Infrastructure for Developers & AI Agents
---

# Vultisig SDK

The Vultisig SDK gives developers and AI agents direct access to Vultisig's multi-party computation (MPC) technology. Instead of relying on seed phrases that create a single point of failure, the SDK splits wallet keys across multiple devices using threshold signatures. The result: secure, self-custodial wallets that work across 40+ blockchains.

Whether you're building a web app, a backend service, or an AI agent that needs to hold and move crypto — the SDK handles vault creation, balance checks, token swaps, and transaction signing.

## What Can You Build With It?

- **Web3 applications** — Onboard users into seedless, cross-chain wallets with MPC security built in
- **AI agent wallets** — Give autonomous agents their own self-custodial crypto wallets, with optional human co-signing for oversight
- **Backend services** — Programmatic vault creation, balance checks, swaps, and signing from your server
- **Desktop & browser apps** — Works in Electron, any modern browser, and Node.js

## How It Works

The SDK supports two vault types:

- **Fast Vault** — A 2-of-2 setup between your app and VultiServer. Signing is instant. Great for quick setup and individual use.
- **Secure Vault** — A multi-device setup (2-of-3, 3-of-5, etc.) where multiple devices must approve transactions. Other participants join by scanning a QR code with the Vultisig mobile app. Best for teams or high-value wallets.

Both vault types support 40+ chains including Bitcoin, Ethereum, Solana, THORChain, Cosmos, and the broader EVM, UTXO, and Cosmos ecosystems. Cross-chain swaps are powered by THORChain, with 1inch for same-chain DEX trades and LiFi for cross-EVM routes.

There's also a full CLI (`@vultisig/cli`) that mirrors the SDK's capabilities from the command line.

## For AI Agents

AI agents can use the SDK or CLI to create wallets, check balances, send tokens, and execute swaps autonomously. Two trust models are available:

| Mode | How it works | Signing | Use case |
|------|-------------|---------|----------|
| **Fast Vault** | Agent + VultiServer (2-of-2) | Instant, no human needed | Full agent autonomy |
| **Secure Vault** | Agent + human device (2-of-2) | Human approves via QR scan | Human oversight on every transaction |

The CLI makes agent integration simple:

```bash
npm install -g @vultisig/cli

vultisig create
vultisig balance ethereum -o json
vultisig send ethereum 0xRecipient 0.1
vultisig swap ethereum bitcoin 0.1
```

JSON output, silent mode, and environment variable config (`VAULT_PASSWORD`) plug directly into agent pipelines. Stateless usage is also supported — load a vault file, operate, discard.

For the full agent operating procedure, see [vultisig.com/SKILL.md](https://vultisig.com/SKILL.md).

## Getting Started

Install the SDK:

```bash
npm install @vultisig/sdk
```

Create a vault and start using it:

```typescript
import { Vultisig, MemoryStorage } from '@vultisig/sdk'

const sdk = new Vultisig({ storage: new MemoryStorage() })
await sdk.initialize()

const vaultId = await sdk.createFastVault({
  name: 'My Wallet',
  email: 'user@example.com',
  password: 'SecurePassword123!',
})

const vault = await sdk.verifyVault(vaultId, code)

const address = await vault.address('Ethereum')
const balance = await vault.balance('Ethereum')
```

For the full API reference, CLI documentation, and detailed usage guides, head to the developer docs:

{% content-ref url="../developer-docs/vultisig-sdk/" %}
[Developer Docs — Vultisig SDK](../developer-docs/vultisig-sdk/)
{% endcontent-ref %}

## Resources

- [SDK Repository on GitHub](https://github.com/vultisig/vultisig-sdk)
- [AI Agent Operating Procedure (SKILL.md)](https://vultisig.com/SKILL.md)
- [Machine-readable Agent Manifest (agent.json)](https://vultisig.com/.well-known/agent.json)
- [Full SDK & CLI Documentation](../developer-docs/vultisig-sdk/)
