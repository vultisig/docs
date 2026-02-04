---
description: Structured index of Vultisig documentation for AI agents and language models. SDK and CLI focus.
---

# For AI Agents

> Vultisig: Seedless, multi-chain, self-custody wallet using TSS/MPC. TypeScript SDK and CLI for building AI agents, automated plugins, and crypto applications across 40+ blockchains.

## SDK (`@vultisig/sdk`)

TypeScript SDK for MPC vault creation, signing, balance checks, swaps, and transaction management.

**Install:** `npm install @vultisig/sdk`

**Core flow:** Initialize → Create vault → Verify → Use

```typescript
import { Vultisig, MemoryStorage } from '@vultisig/sdk'

const sdk = new Vultisig({ storage: new MemoryStorage() })
await sdk.initialize()

const vaultId = await sdk.createFastVault({ name: 'Agent Wallet', email: 'agent@example.com', password: 'pass' })
const vault = await sdk.verifyVault(vaultId, code)

const address = await vault.address('Ethereum')
const balance = await vault.balance('Ethereum')
```

**Key methods:**

| Method | What it does |
|--------|-------------|
| `sdk.createFastVault(opts)` | Create 2-of-2 vault with VultiServer (instant signing) |
| `sdk.createSecureVault(opts)` | Create N-of-M multi-device vault (human co-signing) |
| `sdk.verifyVault(vaultId, code)` | Verify vault via email code, returns vault |
| `vault.address(chain)` | Derive address for a chain |
| `vault.balance(chain)` | Get native balance |
| `vault.balances(chains, includeTokens)` | Get balances across chains |
| `vault.prepareSendTx(params)` | Prepare a send transaction |
| `vault.sign(payload)` | Sign a transaction (MPC) |
| `vault.broadcastTx(params)` | Broadcast signed transaction |
| `vault.gas(chain)` | Get gas/fee estimate |
| `vault.getSwapQuote(params)` | Get swap quote (THORChain, 1inch, LiFi) |
| `vault.prepareSwapTx(params)` | Prepare swap transaction (handles approval) |
| `vault.signBytes(opts)` | Sign arbitrary pre-hashed bytes |
| `vault.broadcastRawTx(params)` | Broadcast pre-signed raw transaction |
| `sdk.importVault(content, password)` | Import vault from `.vult` file |
| `vault.export(password)` | Export vault to backup |

**Vault types:**

| | Fast Vault | Secure Vault |
|---|-----------|-------------|
| **Threshold** | 2-of-2 (with VultiServer) | N-of-M (configurable) |
| **Signing** | Instant, no human needed | Requires device coordination via QR |
| **Agent use case** | Full autonomy | Human oversight on every tx |

**Supported chains (36+):** Bitcoin, Ethereum, Solana, THORChain, Polygon, Arbitrum, Optimism, Base, BSC, Avalanche, Cosmos, Litecoin, Dogecoin, Sui, TON, Ripple, Tron, Polkadot, Cardano, and more.

**Full docs:**
- [SDK README](developer-docs/vultisig-sdk/README.md): Installation, quick start, API reference, vault types, error handling
- [SDK Implementation Guide](developer-docs/vultisig-sdk/SDK-USERS-GUIDE.md): Complete usage guide — password management, vault lifecycle, transactions, swaps, events, caching, platform notes

---

## CLI (`@vultisig/cli`)

Command-line wallet for scripting, automation, and agent pipelines. Mirrors the SDK's full capabilities.

**Install:** `npm install -g @vultisig/cli`

**Key commands:**

```bash
# Vault management
vultisig create fast --name "Wallet" --email user@example.com --password pass
vultisig create secure --name "Team Wallet" --shares 3
vultisig import /path/to/vault.vult
vultisig vaults
vultisig export

# Balances & addresses
vultisig balance                      # All chains
vultisig balance ethereum --tokens    # Specific chain + tokens
vultisig addresses
vultisig portfolio

# Transactions
vultisig send ethereum 0xRecipient 0.1
vultisig send ethereum 0xRecipient 100 --token 0xTokenAddress

# Swaps
vultisig swap-quote ethereum bitcoin 0.1
vultisig swap ethereum bitcoin 0.1

# Advanced: sign arbitrary bytes, broadcast raw tx
vultisig sign --chain ethereum --bytes "base64hash" -o json
vultisig broadcast --chain ethereum --raw-tx "0x02f8..."

# Seedphrase import
vultisig create-from-seedphrase fast --name "Imported" --email user@example.com --discover-chains
```

**Agent-friendly features:**
- `--output json` (or `-o json`) — structured JSON for all commands
- `--silent` — suppress spinners and progress messages
- `--password` flag — avoid interactive prompts
- `VAULT_PASSWORD` env var — for automation pipelines
- `VULTISIG_VAULT` env var — pre-select vault by name or ID
- Exit codes: 0 success, 1-7 for specific error types
- `vsig` shorthand alias for `vultisig`

**Full docs:** [CLI Documentation](developer-docs/vultisig-sdk/CLI.md): All commands, options, environment variables, JSON output examples, exit codes, interactive shell

---

## Agent Resources

Files on [vultisig.com](https://vultisig.com) for agent discovery and integration:

| File | What it is |
|------|-----------|
| [SKILL.md](https://vultisig.com/SKILL.md) | Full operating procedure — 14 steps covering vault creation, sends, swaps, balances, gas estimation |
| [llms.txt](https://vultisig.com/llms.txt) | Spec-compliant link index (llmstxt.org format) |
| [llms-full.txt](https://vultisig.com/llms-full.txt) | Full SDK context with verified code examples and source references |
| [agent.json](https://vultisig.com/.well-known/agent.json) | Structured capabilities manifest — chains, operations, SDK info |

---

## Documentation Index

### Getting Started
- [Overview](README.md): What Vultisig is, key features
- [Download & Install](getting-started/download-install.md): iOS, Android, macOS, Windows, Linux, browser extension
- [Create a Vault](getting-started/create-vault.md): Fast Vault and Secure Vault creation
- [Backup & Recovery](getting-started/backup-recovery.md): Vault backup and restore
- [Your First Transaction](getting-started/first-transaction.md): Send your first transaction

### App Guide
- [Vault Creation](app-guide/creating-a-vault/README.md): Vault types and flows
  - [Fast Vault](app-guide/creating-a-vault/fast-vault.md): 2-of-2 with VultiServer
  - [Secure Vault](app-guide/creating-a-vault/secure-vault.md): Multi-device with QR pairing
- [Sending](app-guide/wallet/sending.md): How to send tokens
- [Swapping](app-guide/wallet/swapping.md): Cross-chain and same-chain swaps
- [DeFi](app-guide/defi/README.md): Circle Protocol, THORChain, MayaChain, staking
- [Vault Management](app-guide/vault-management/README.md): Details, backups, reshare, rename, upgrade

### Security & Technology
- [Overview](security-technology/README.md): Security architecture
- [Keysign](security-technology/keysign.md): Transaction signing
- [TSS Actions](security-technology/tss-actions.md): Threshold signature operations
- [How GG20 Works](security-technology/how-gg20-works.md): GG20 protocol
- [How DKLS23 Works](security-technology/how-dkls23-works.md): DKLS23 protocol
- [Difference to Multi-Signatures](security-technology/difference-to-multi-sig.md): TSS vs multisig
- [Emergency Recovery](security-technology/emergency-recovery.md): Recovery procedures

### Ecosystem
- [Vultisig Extension](vultisig-ecosystem/vultisig-extension/README.md): Browser extension for dApps
- [Plugin Marketplace](vultisig-ecosystem/marketplace.md): Self-custodial automation — plugins and AI agents
- [Web App](vultisig-ecosystem/web-app.md): Browser-based vault access
- [Vultisig SDK](vultisig-ecosystem/vultisig-sdk.md): SDK overview with agent section
- [Community Tools](vultisig-ecosystem/community-tools.md): Third-party tools

### VULT Token
- [The $VULT Token](vultisig-token/vult/README.md): Tokenomics and utility
- [In-App Utility](vultisig-token/vult/in-app-utility.md): Fee discounts, staking tiers
- [Marketplace Utility](vultisig-token/vult/marketplace-utility.md): Revenue distribution
- [Governance](vultisig-token/vult/governance-utility.md): Voting rights

### Infrastructure
- [Overview](vultisig-infrastructure/overview.md): Architecture
- [Vultiserver](vultisig-infrastructure/what-is-vultisigner/README.md): Co-signing server for Fast Vaults
- [Transaction Policies](vultisig-infrastructure/what-is-vultisigner/what-can-be-configured.md): Spending limits, whitelists, time delays
- [Relay Server](vultisig-infrastructure/relay-server.md): Device session coordination

### Developer Docs
- [Developer Home](developer-docs/README.md): Entry point
- [Marketplace Plugins](developer-docs/marketplace/README.md): Build and publish plugins
  - [What is a Plugin](developer-docs/marketplace/infrastructure-overview/plugins.md): Architecture and scope
  - [Services Architecture](developer-docs/marketplace/infrastructure-overview/services.md): Service components
  - [Policy Rules](developer-docs/marketplace/infrastructure-overview/metarules.md): Transaction validation rules
  - [Infrastructure](developer-docs/marketplace/infrastructure-overview/infrastructure.md): Plugin infrastructure
  - [Quick Start](developer-docs/marketplace/create-a-plugin/basics-quick-start.md): Scaffold your first plugin
  - [Build Your Plugin](developer-docs/marketplace/create-a-plugin/build-your-plugin/README.md): Full development guide
  - [Submission & Revenue](developer-docs/marketplace/create-a-plugin/submission-process.md): Review process, 70/30 split
- [Extension Integration](developer-docs/vultisig-extension-integration-guide.md): window.vultisig API, code examples
- [SDK](developer-docs/vultisig-sdk/README.md): Full SDK docs
  - [SDK Implementation Guide](developer-docs/vultisig-sdk/SDK-USERS-GUIDE.md): Detailed usage guide
  - [SDK CLI](developer-docs/vultisig-sdk/CLI.md): CLI reference

### Help & Legal
- [FAQ](help/faq.md): Common questions
- [Security](help/security.md): Security policy
- [Privacy](help/privacy.md): Privacy policy
- [Terms of Use](help/terms.md): Terms
