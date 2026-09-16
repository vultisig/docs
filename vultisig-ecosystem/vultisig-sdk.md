---
description: >-
  Vultisig TypeScript SDK and CLI. Create vaults, check balances, send, swap,
  and sign across 36+ chains. No seed phrase.
---

# Vultisig SDK

TypeScript SDK for vault create, verify, balance, send, swap, and sign. Threshold signatures. No seed phrase. 36+ chains.

Use it from a web app, a backend, Electron, or an agent. CLI (`@vultisig/cli`) mirrors the same surface.

Full reference: [SDK docs](../developer-docs/vultisig-sdk/).

## Vault types

- **Fast Vault** — 2-of-2 with VultiServer. VultiServer co-signs and can never sign alone.
- **Secure Vault** — N-of-M on devices you control. Other shares join by QR from the Vultisig app.

Both vault types cover 36+ chains. The wallet picks the best swap quote (THORChain, Maya, 1inch, Kyber, LiFi, Jupiter, and others).

## For AI Agents

SDK or CLI. Two trust models:

| Mode | How it works | Signing | Use case |
|------|-------------|---------|----------|
| **Fast Vault** | Agent + VultiServer (2-of-2) | Instant, no human needed | Full agent autonomy |
| **Secure Vault** | Agent + human device (2-of-2) | Human approves via QR scan | Human oversight on every transaction |

CLI for unattended agents:

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
import { Vultisig } from '@vultisig/sdk'

const sdk = new Vultisig()
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
