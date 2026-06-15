---
description: >-
  Choose between the Vultisig CLI, TypeScript SDK, Marketplace plugins, and MCP
  tooling for AI agent integrations.
---

# Integration options

## CLI

Use the Vultisig CLI when an agent can run shell commands. It exposes structured JSON output, non-interactive options, quote commands, and specific exit codes for reliable automation.

```bash
npm install -g @vultisig/cli

vultisig portfolio --output json
vultisig swap-quote ethereum bitcoin 0.1 --output json
vultisig swap ethereum bitcoin 0.1 --output json
```

See the [CLI reference](../vultisig-sdk/CLI.md).

## TypeScript SDK

Use `@vultisig/sdk` when building wallet capabilities directly into an agent, bot, or application. The SDK handles vault creation, balances, portfolio tracking, sends, swaps, message signing, and transaction broadcasting.

```typescript
import { Chain, Vultisig } from '@vultisig/sdk'

const sdk = new Vultisig()
await sdk.initialize()

const vault = await sdk.getActiveVault()
if (!vault) throw new Error('No active vault')

const portfolio = await vault.portfolio('usd')
const preview = await vault.send({
  chain: Chain.Ethereum,
  to: '0xRecipient',
  amount: '0.1',
  dryRun: true,
})
```

Start with the [SDK guide](../vultisig-sdk/) and [implementation guide](../vultisig-sdk/SDK-USERS-GUIDE.md).

## Marketplace plugins

Use a Marketplace plugin when an agent should act within user-defined rules instead of receiving broad wallet authority. The plugin proposes an unsigned transaction; the Vultisig-managed Verifier checks the proposal before participating in signing.

Start with the [Marketplace developer docs](../marketplace/).

## Agent discovery resources

| Resource | Purpose |
| --- | --- |
| [SKILL.md](https://vultisig.com/SKILL.md) | Operating instructions for balances, sends, swaps, gas estimation, and vault workflows |
| [llms.txt](https://vultisig.com/llms.txt) | Compact documentation index |
| [llms-full.txt](https://vultisig.com/llms-full.txt) | Full SDK context and examples |
| [agent.json](https://vultisig.com/.well-known/agent.json) | Structured capabilities manifest |
| [MCP server](https://github.com/vultisig/mcp) | MCP tools for multi-chain queries, unsigned transaction building, DeFi interactions, and plugin management |
