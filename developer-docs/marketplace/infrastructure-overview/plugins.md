# What is a Plugin?

Vultisig plugins are vault automation workflows created by Vultisig or third-party developers that enable automatic on-chain interactions for users. Due to Vultisig's use of MPC technology, these plugins are not limited by any specific chain, action, or protocol interaction. This gives developers the freedom to design secure, self-custodial automations without users giving up private keys or granting full access to wallet funds.

## How Plugins Work

A plugin operates as a trusted automation layer between the user and the blockchain:

1. **User Installs Plugin**: The user selects a plugin from the Marketplace and configures their automation (e.g., "buy $100 of ETH every week")
2. **Policy Creation**: The plugin creates a policy—a signed set of rules defining exactly what transactions are allowed
3. **Verifier Validation**: When the plugin wants to execute a transaction, it submits it to the Verifier, which checks it against the user's policy
4. **TSS Signing**: If compliant, the Verifier participates in a threshold signing ceremony to approve the transaction
5. **Broadcast**: The plugin broadcasts the signed transaction to the blockchain

```
┌──────────────┐     ┌──────────────┐     ┌──────────────┐
│     User     │     │   Plugin     │     │   Verifier   │
│              │     │              │     │              │
│ • Installs   │────►│ • Runs logic │────►│ • Validates  │
│ • Configures │     │ • Builds txs │     │ • Signs      │
│ • Sets rules │     │ • Broadcasts │     │              │
└──────────────┘     └──────────────┘     └──────────────┘
```

## Plugin Independence

Every plugin is a standalone, independent service. There is no shared runtime or plugin framework that constrains what you build. Each plugin:

- **Runs its own infrastructure** — your own HTTP server, workers, database, and task queue
- **Defines its own recipe** — the rules and constraints governing what transactions it can request
- **Sets its own auth and configuration** — each plugin determines what users need to configure and how policies are structured
- **Controls its own business logic** — DCA schedules, AI trading strategies, payroll batches, or anything else you design

The [App Recurring](https://github.com/vultisig/app-recurring) plugin is a reference starting point, not a rigid template. The Vultisig ecosystem provides the security layer (Verifier + TSS signing) and distribution (Marketplace). Everything else is yours to build.

Users must configure policies specific to each installed plugin. A DCA plugin's policies look different from a payroll plugin's policies. Your plugin's `Suggest()` function guides users toward the right rules and constraints for your specific use case.

## Key Security Properties

**Self-Custodial**: Users never give up their private keys. Plugins cannot access funds outside the defined policy rules.

**Policy-Bound**: Every transaction must match the policy the user signed. A plugin configured for "weekly ETH buys" cannot suddenly drain the wallet.

**Multi-Party Signing**: Transactions require TSS coordination—the plugin alone cannot sign transactions without the Verifier's participation.

**Transparent Rules**: Policies use a rules system (MetaRules and Direct Rules) that maps to specific blockchain operations, making permissions auditable.

## Types of Plugins

| Category | Examples | Description |
|----------|----------|-------------|
| **DeFi Automation** | DCA, yield farming, rebalancing | Recurring or conditional trading strategies |
| **AI Agents** | Trading bots, portfolio managers | Autonomous agents with constrained permissions |
| **Treasury Management** | Multi-sig ops, payroll | Organizational fund management |
| **Cross-Chain** | Bridge automation, arbitrage | Operations spanning multiple blockchains |

## Plugin Components

A complete plugin consists of:

| Component | Purpose | Required |
|-----------|---------|----------|
| **HTTP Server** | API endpoints for reshare, sign, and automation management | Yes |
| **Specification** | Defines UI, validation rules, and supported operations | Yes |
| **Worker** | Background processing for TSS signing and transaction monitoring | Yes |
| **Trigger** | Initiates transactions (scheduler, event listener, HTTP endpoint) | Varies |

See [Services Architecture](services.md) for detailed component documentation.

## Reference Implementation

The [App Recurring](https://github.com/vultisig/app-recurring) plugin is the reference implementation for building Vultisig plugins. Use it as a starting point — not a template you must follow. Your plugin defines its own recipe, configuration flow, and business logic.
