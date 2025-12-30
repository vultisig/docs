# What is an App?

Vultisig apps are vault automation workflows created by Vultisig or third-party developers that enable automatic on-chain interactions for users. Due to Vultisig's use of MPC technology, these apps are not limited by any specific chain, action, or protocol interaction. This gives developers the freedom to design secure, self-custodial automations without users giving up private keys or granting full access to wallet funds.

## How Apps Work

An app operates as a trusted automation layer between the user and the blockchain:

1. **User Installs App**: The user selects an app from the App Store and configures their automation (e.g., "buy $100 of ETH every week")
2. **Policy Creation**: The app creates a policy—a signed set of rules defining exactly what transactions are allowed
3. **Verifier Validation**: When the app wants to execute a transaction, it submits it to the Verifier, which checks it against the user's policy
4. **TSS Signing**: If compliant, the Verifier participates in a threshold signing ceremony to approve the transaction
5. **Broadcast**: The app broadcasts the signed transaction to the blockchain

```
┌──────────────┐     ┌──────────────┐     ┌──────────────┐
│     User     │     │     App      │     │   Verifier   │
│              │     │              │     │              │
│ • Installs   │────►│ • Runs logic │────►│ • Validates  │
│ • Configures │     │ • Builds txs │     │ • Signs      │
│ • Sets rules │     │ • Broadcasts │     │              │
└──────────────┘     └──────────────┘     └──────────────┘
```

## Key Security Properties

**Self-Custodial**: Users never give up their private keys. Apps cannot access funds outside the defined policy rules.

**Policy-Bound**: Every transaction must match the policy the user signed. An app configured for "weekly ETH buys" cannot suddenly drain the wallet.

**Multi-Party Signing**: Transactions require TSS coordination—the app alone cannot sign transactions without the Verifier's participation.

**Transparent Rules**: Policies use a rules system (MetaRules and Direct Rules) that maps to specific blockchain operations, making permissions auditable.

## Types of Apps

| Category | Examples | Description |
|----------|----------|-------------|
| **DeFi Automation** | DCA, yield farming, rebalancing | Recurring or conditional trading strategies |
| **AI Agents** | Trading bots, portfolio managers | Autonomous agents with constrained permissions |
| **Treasury Management** | Multi-sig ops, payroll | Organizational fund management |
| **Cross-Chain** | Bridge automation, arbitrage | Operations spanning multiple blockchains |

## App Components

A complete app consists of:

| Component | Purpose | Required |
|-----------|---------|----------|
| **HTTP Server** | API endpoints for reshare, sign, and automation management | Yes |
| **Specification** | Defines UI, validation rules, and supported operations | Yes |
| **Worker** | Background processing for TSS signing and transaction monitoring | Yes |
| **Trigger** | Initiates transactions (scheduler, event listener, HTTP endpoint) | Varies |

See [Services Architecture](services.md) for detailed component documentation.

## Reference Implementation

The [DCA (Dollar Cost Averaging)](https://github.com/vultisig/dca) app is the reference implementation for building Vultisig apps. It demonstrates:

- Complete server setup with all required endpoints
- Policy specification and validation
- Worker implementation for async signing
- Scheduler-based transaction triggers

Use it as a starting point for your own app development.
