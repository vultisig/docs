---
description: >-
  Choose between the Vultisig CLI, TypeScript SDK, Marketplace plugins, and MCP
  tooling for AI agent integrations.
---

# Integration options

## CLI (start here)

The CLI is the fastest way to give an AI agent a wallet. Any agent that can run shell commands gets multi-chain balances, sends, swaps, and signing — through structured JSON, stable exit codes, and a natural-language `agent ask` mode built for AI-to-AI use. No SDK build step, no embedding.

```bash
npm install -g @vultisig/cli

# Structured wallet commands
vultisig portfolio --output json
vultisig swap-quote ethereum bitcoin 0.1 --output json
vultisig swap ethereum bitcoin 0.1 --output json

# Natural-language, one-shot — designed for AI agent integration
vultisig agent ask "What is my ETH balance?" --password "$VAULT_PASSWORD" --json
```

The CLI auto-detects non-interactive (non-TTY) environments and skips the prompts that would otherwise hang an agent, so the same commands run unattended in scripts and pipelines. Set `VAULT_PASSWORD` to avoid interactive password entry, and branch on the documented exit codes and error codes for reliable orchestration.

See the [CLI reference](../vultisig-sdk/CLI.md), including the AI agent integration guide.

## TypeScript SDK

Use `@vultisig/sdk` when you are embedding wallet capabilities directly into your own agent, bot, or service in TypeScript. The SDK handles vault creation, balances, portfolio tracking, sends, swaps, message signing, and transaction broadcasting.

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

Use a Marketplace plugin when an agent should act within user-defined rules instead of receiving broad wallet authority. The plugin proposes an unsigned transaction; the Verifier checks the proposal before participating in signing.

Start with the [Marketplace developer docs](../marketplace/).

## Agent discovery resources

| Resource | Purpose |
| --- | --- |
| [SKILL.md](https://vultisig.com/SKILL.md) | Operating instructions for balances, sends, swaps, gas estimation, and vault workflows |
| [llms.txt](https://vultisig.com/llms.txt) | Compact documentation index |
| [llms-full.txt](https://vultisig.com/llms-full.txt) | Full SDK context and examples |
| [agent.json](https://vultisig.com/.well-known/agent.json) | Structured capabilities manifest |
| [MCP server](https://github.com/vultisig/mcp) | MCP tools for multi-chain queries, unsigned transaction building, DeFi interactions, and plugin management |
