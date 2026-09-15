---
description: >-
  Build AI agents with multi-chain wallet capabilities and choose exactly how
  much signing authority they receive.
---

# Build AI agents on Vultisig

Vultisig gives AI agents a programmable, seedless wallet across 36+ blockchains. Agents can inspect portfolios, prepare transactions, send assets, swap across chains, and sign messages without ever holding a complete private key.

The defining choice is not whether an agent can transact. It is **how much authority the agent should have**.

## Signing models

Same vault, three authority levels. Pick the least authority that still lets the workflow run.

| Operating model | Best for | Signing authority |
| --- | --- | --- |
| **Autonomous Fast Vault** | Bots and services that must transact without a person present | The agent and VultiServer each hold one share of a 2-of-2 MPC vault |
| **Policy-bound automation** | Autonomous strategies with explicit user limits | A Marketplace plugin proposes; the Verifier enforces the configured rules. **This path is paused and not in production.** |
| **Human-approved Secure Vault** | Higher-value workflows requiring approval | The configured threshold of human-controlled devices joins every signing session |

Use the least authority that still lets the workflow succeed. See [Authority and security](authority-and-security.md) for the full decision guide.

## What agents can do

### Portfolio intelligence

> "Summarize my holdings across Bitcoin, Ethereum, Solana, and Cosmos."

Query addresses, native and token balances, and fiat values without signing a transaction.

### Preview and execute transactions

> "Prepare a 0.1 ETH transfer, show me the fee and total, then wait for approval."

Prepare and inspect transaction details before entering the selected signing flow.

### Cross-chain swaps

> "Swap 0.1 ETH to BTC using the best available route."

Request quotes and execute swaps through supported providers such as THORChain, 1inch, KyberSwap, and LiFi.

### Policy-bound strategies

> "Rebalance when BTC moves outside 45-55% of the portfolio, but never trade more than $500 per day."

Marketplace agents can monitor conditions and propose transactions while a Verifier enforces the user's rules. This Marketplace path is paused and not in production.

## Start building

* **Give any agent a wallet (start here):** the [CLI](../vultisig-sdk/CLI.md) is the fastest path — structured JSON, stable exit codes, and a natural-language `agent ask` mode built for AI-to-AI use. It runs unattended in non-interactive environments with no extra setup.
* **Embed wallet logic into your own service:** integrate the [TypeScript SDK](../vultisig-sdk/).
* **Publish constrained autonomous strategies:** build a [Marketplace plugin](../marketplace/) (paused, not in production).
* **Choose the right approval model:** read [Authority and security](authority-and-security.md).
* **Compare the available developer surfaces:** read [Integration options](integration-options.md).
