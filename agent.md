---
description: >-
  Give AI agents multi-chain wallet capabilities with a choice of autonomous,
  policy-bound, or human-approved signing.
---

# Agent capabilities with Vultisig

Vultisig gives agents a programmable, seedless wallet across 40+ blockchains. Agents can inspect a portfolio, prepare transactions, send assets, swap across chains, and sign messages without ever holding a complete private key.

The important choice is not whether an agent can transact. It is **how much authority the agent should have**. Vultisig supports three operating models, from full runtime autonomy to human approval on every signature.

## Choose the right operating model

| Operating model | Best for | How signing works | Start with |
| --- | --- | --- | --- |
| **Autonomous Fast Vault** | Bots, services, and agents that must transact without a person present | The agent's device and VultiServer form a 2-of-2 MPC vault. Signing is immediate. | [SDK](developer-docs/vultisig-sdk/) or [CLI](developer-docs/vultisig-sdk/CLI.md) |
| **Policy-bound automation** | Recurring workflows and autonomous strategies with defined limits | A Marketplace plugin proposes transactions. A separate Verifier signs only when the proposal matches the user's configured rules. | [Plugin Marketplace](vultisig-ecosystem/marketplace.md) |
| **Human-approved Secure Vault** | Higher-value or sensitive workflows where a person should approve every transaction | The agent prepares the action, then the configured threshold of devices joins the MPC signing session. | [SDK implementation guide](developer-docs/vultisig-sdk/SDK-USERS-GUIDE.md) |

These models can support the same agent experience while placing the final signing authority in different places. Use the least authority that still lets the workflow succeed.

## Pick the integration surface

### CLI: give an existing agent a wallet

Use the Vultisig CLI when an agent can run shell commands. It exposes structured JSON output, non-interactive options, quote commands, and specific exit codes for reliable agent workflows.

```bash
npm install -g @vultisig/cli

# Read-only portfolio snapshot
vultisig portfolio --output json

# Preview a cross-chain swap quote
vultisig swap-quote ethereum bitcoin 0.1 --output json

# Execute a cross-chain swap
vultisig swap ethereum bitcoin 0.1 --output json
```

See the [CLI reference](developer-docs/vultisig-sdk/CLI.md) for vault creation, environment variables, commands, and JSON output.

### SDK: build wallet capabilities into an agent

Use `@vultisig/sdk` when you are building a TypeScript agent, bot, or application. The SDK handles vault creation, balances, portfolio tracking, sends, swaps, message signing, and transaction broadcasting.

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

Start with the [SDK guide](developer-docs/vultisig-sdk/) and [implementation guide](developer-docs/vultisig-sdk/SDK-USERS-GUIDE.md).

### Marketplace: publish policy-bound automation

Use a Marketplace plugin when an agent should act within user-defined rules rather than receive broad wallet authority. The plugin proposes an unsigned transaction; an independent Verifier checks the destination, amount, timing, and other configured rules before participating in MPC signing.

This model is designed for workflows such as recurring swaps, payments, portfolio rebalancing, and condition-based strategies. Learn about the [Plugin Marketplace](vultisig-ecosystem/marketplace.md) or [build a plugin](developer-docs/marketplace/).

## What agents can do

### Read-only portfolio intelligence

> "Summarize my portfolio across Bitcoin, Ethereum, Solana, and Cosmos."

The agent queries addresses, native and token balances, and fiat values without signing a transaction.

### Preview, then execute

> "Prepare a 0.1 ETH transfer, show me the fee and total, then wait for approval."

The agent uses a dry run to surface the expected cost and transaction details before entering the selected signing flow.

### Natural-language cross-chain swaps

> "Swap 0.1 ETH to BTC using the best available route."

The agent can request a quote, explain estimated output and fees, then execute through THORChain, 1inch, KyberSwap, LiFi, or another supported route.

### Policy-bound automation

> "Rebalance when BTC moves outside 45-55% of the portfolio, but never trade more than $500 per day."

A Marketplace agent can monitor conditions and propose transactions while the Verifier enforces the rules configured by the user.

## The security boundary

Vultisig removes the complete private key as a single point of failure, but MPC does not make an agent's decisions correct. Your application is still responsible for prompts, strategy logic, token and contract selection, transaction parameters, credential storage, and economic outcomes.

Before allowing an agent to sign:

* Use dry runs and surface the destination, assets, fees, and expected output.
* Prefer a Secure Vault when every transaction needs human approval.
* Prefer Marketplace policies when autonomy should be limited by explicit rules.
* Protect vault passwords and backups; do not embed them in prompts, source code, or logs.
* Treat arbitrary message and byte signing as high-risk capabilities.

{% hint style="warning" %}
Fast Vaults enable unattended signing. Only fund an autonomous vault with an amount appropriate for the workflow's risk, and keep a tested vault backup.
{% endhint %}

## Agent discovery

Vultisig publishes machine-readable resources for agents and coding assistants:

| Resource | Purpose |
| --- | --- |
| [SKILL.md](https://vultisig.com/SKILL.md) | Operating instructions for balances, sends, swaps, gas estimation, and vault workflows |
| [llms.txt](https://vultisig.com/llms.txt) | Compact documentation index |
| [llms-full.txt](https://vultisig.com/llms-full.txt) | Full SDK context and examples |
| [agent.json](https://vultisig.com/.well-known/agent.json) | Structured capabilities manifest |
| [MCP server](https://github.com/vultisig/mcp) | Experimental MCP tools for EVM reads and transaction building |

{% hint style="info" %}
The MCP server is a work in progress. For production wallet operations, use the SDK, CLI, or Marketplace plugin infrastructure.
{% endhint %}

## Get started

* **Give an existing coding agent a wallet:** install the [CLI](developer-docs/vultisig-sdk/CLI.md).
* **Build a long-running agent or bot:** integrate the [TypeScript SDK](developer-docs/vultisig-sdk/).
* **Build constrained autonomous strategies:** start with the [Marketplace developer docs](developer-docs/marketplace/).
* **Require a person to approve every signature:** use a [Secure Vault](app-guide/creating-a-vault/secure-vault.md).
