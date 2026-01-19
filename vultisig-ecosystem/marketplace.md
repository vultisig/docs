---
description: >-
  Self-custodial automation marketplace. Install DCA, recurring payments, and
  trading plugins without giving up custody. MPC-secured, multi-chain.
cover: ../.gitbook/assets/image (3).png
coverY: 0
---

# Marketplace

{% hint style="warning" %}
The Marketplace is in early access. Use for testing only—do not use with production funds.
{% endhint %}

## Overview

The Vultisig Marketplace is a marketplace for self-custodial automation. Install plugins to automate your crypto management—like recurring investments (DCA) and scheduled payments—without giving up custody of your funds.

Unlike traditional automation that requires handing over your keys, Vultisig plugins use MPC technology to keep you in control. Your vault executes automations securely—no third parties, no smart contract risks, no seed phrase exposure.

Browse plugins at [apps.vultisig.com](https://apps.vultisig.com).

***

## How It Works

Vultisig plugins use a **proposing/validating architecture** that keeps you in control through every step.

### Setup & Installation

Installing a plugin creates a secure connection between your vault and the plugin's automation service. Here's what happens:

**Step 1: Browse & Select**

Visit [apps.vultisig.com](https://apps.vultisig.com), connect your Vultisig vault, and choose a plugin.

**Step 2: Install the Plugin**

Approve the installation to connect your vault to the plugin. This creates a secure link between your vault and the plugin's Verifier service:

| What Changes | Before Installation | After Installation |
|--------------|--------------------|--------------------|
| **Signing parties** | Your device + Vultisigner | Your device + Plugin Verifier |
| **Who can propose** | Only you | Plugin (within your rules) |
| **Who validates** | — | Verifier checks every transaction |

{% hint style="info" %}
**Technical note:** Under the hood, installation performs a "reshare"—your vault's signing arrangement is reconfigured so the plugin's Verifier can participate in signing, but ONLY for transactions you authorize.
{% endhint %}

**Step 3: Configure Your Rules**

Now set your automation parameters:
- Which assets to use
- How much per transaction (amount limits)
- How often (schedule/frequency)
- Where to send (recipient addresses)

**Step 4: Automation Goes Live**

Your automation policy is now active. The plugin can propose transactions within your configured limits, and the Verifier ensures nothing outside your rules ever gets signed.

### Execution Flow

Once installed, your automation runs automatically:

| Step | What Happens |
|------|--------------|
| **1. Trigger** | Scheduler triggers based on your rules (e.g., "every Monday at 9am") |
| **2. Propose** | Plugin creates an unsigned transaction matching your policy |
| **3. Validate** | Verifier checks the transaction against your exact rules |
| **4. Sign** | MPC signs only if validation passes—keys never leave your devices |
| **5. Broadcast** | Signed transaction is sent to the blockchain |

{% hint style="info" %}
**Security guarantee:** The Verifier rejects any transaction that doesn't match your configured rules exactly. If a plugin tries to send more than you authorized, to a different address, or at the wrong time—it gets blocked.
{% endhint %}

***

## Available Plugins

Current plugins available at [apps.vultisig.com](https://apps.vultisig.com):

### Recurring Swaps (DCA)

Automate dollar-cost averaging into any asset. Convert one token to another on a schedule you define.

**Use cases:**
- Weekly Bitcoin accumulation
- Monthly portfolio rebalancing
- Gradual position building

### Recurring Sends

Schedule automatic payments to any address. Set up payroll, subscriptions, or regular transfers.

**Use cases:**
- Payroll automation
- Subscription payments
- Regular savings transfers

***

## For Users

### Getting Started

1. Open [apps.vultisig.com](https://apps.vultisig.com)
2. Connect your Vultisig vault
3. Browse and select a plugin
4. Approve the installation
5. Configure your automation rules

Each plugin shows its risk level and fee structure before installation.

### Security Features

| Feature | What It Means |
|---------|---------------|
| **Risk ratings** | Plugins are reviewed and assigned risk levels (Low/Medium/High) |
| **Rule-based execution** | Plugins can only do what you've configured—nothing more |
| **No key exposure** | Your private keys never leave your devices |
| **Transparent fees** | All costs shown upfront before you install |

***

## For Developers

Build plugins for Vultisig users and earn 70% of the revenue. Plugins are reviewed for quality and security before listing.

**What you can build:**
- Trading bots and DCA strategies
- Portfolio management tools
- Payment automation
- DeFi integrations
- Cross-chain workflows

{% hint style="info" %}
Ready to build? Check the [Developer Documentation](../developer-docs/marketplace/) for Plugins and join [Discord](https://discord.gg/vultisig) for support.
{% endhint %}

***

## Fee Structure

Plugin revenue is split:

| Recipient | Share |
|-----------|-------|
| Plugin developer | 70% |
| $VULT token | 30% |

<figure><img src="../.gitbook/assets/Group 1000004758.png" alt="70/30 revenue split visualization"><figcaption><p>Revenue split: 70% to developers, 30% to $VULT</p></figcaption></figure>

**Fee models available to developers:**
- **Per-transaction** — charged each time the plugin executes
- **Subscription** — monthly or yearly recurring fee
- **Per-installation** — one-time setup fee

***

## Vultisig vs Alternatives

How does Vultisig Marketplace compare to other automation options?

| Feature | Vultisig Marketplace | CEX Recurring Buy | Smart Contract Bots | Custodial Services |
|---------|-------------------|-------------------|---------------------|-------------------|
| **Self-custody** | ✅ Yes | ❌ No | ⚠️ Partial | ❌ No |
| **Seed phrase exposure** | ✅ None | N/A | ⚠️ Often required | ❌ Full access |
| **Multi-chain** | ✅ 10+ chains | ❌ Single platform | ⚠️ Per-chain deployment | ⚠️ Varies |
| **Smart contract risk** | ✅ None | ✅ None | ❌ Yes | ✅ None |
| **Rule-based limits** | ✅ Enforced by MPC | ❌ No limits | ⚠️ Contract-dependent | ❌ No limits |
| **Transparent fees** | ✅ Upfront | ⚠️ Hidden spreads | ⚠️ Gas + protocol fees | ⚠️ Varies |
| **Works with DeFi** | ✅ Yes | ❌ No | ✅ Yes | ⚠️ Limited |

**Key differentiators:**
- **True self-custody:** Unlike CEX recurring buys, your funds stay in your vault
- **No smart contract risk:** Unlike DeFi bots, there's no contract to exploit
- **Cross-chain native:** One setup works across Bitcoin, Ethereum, Solana, and more
- **Enforced limits:** MPC verification ensures apps can't exceed your rules

***

## FAQ

### General

<details>
<summary><strong>What happens if I run out of funds?</strong></summary>

The automation simply skips that execution. No failed transactions, no wasted gas. The scheduler will try again at the next scheduled time. You'll see a notification in the app.

</details>

<details>
<summary><strong>Can I cancel or pause an automation?</strong></summary>

Yes. You can pause, modify, or delete any automation at any time from apps.vultisig.com. Changes take effect immediately.

</details>

<details>
<summary><strong>What chains are supported?</strong></summary>

The Marketplace supports all chains that Vultisig supports, including:
- EVM chains (Ethereum, Polygon, Arbitrum, Base, Optimism, Avalanche, BSC)
- Bitcoin
- Solana
- THORChain
- And more

Each plugin specifies which chains it supports.

</details>

<details>
<summary><strong>Do I need to keep my devices online?</strong></summary>

No. Once you've set up an automation, it runs on Vultisig's infrastructure. Your devices only need to be online for the initial setup and any changes.

</details>

### Security

<details>
<summary><strong>Can a plugin drain my wallet?</strong></summary>

No. Plugins can only execute transactions that match your configured rules exactly. The Verifier service validates every transaction before MPC signing. If a plugin tries to send more than you authorized or to a different address, the transaction is rejected.

</details>

<details>
<summary><strong>What if Vultisig's servers go down?</strong></summary>

Your funds remain safe in your vault. Automations would pause until service is restored. You can always access your funds directly through the Vultisig app—the Marketplace is an optional feature.

</details>

<details>
<summary><strong>How are plugins reviewed?</strong></summary>

All plugins go through a security review before listing. We check:
- Code quality and security practices
- Transaction logic matches declared functionality
- No hidden fees or unexpected behaviors
- Proper error handling

Plugins are assigned risk ratings (Low/Medium/High) based on their complexity and permissions.

</details>

<details>
<summary><strong>Is my seed phrase ever exposed?</strong></summary>

No. Vultisig is seedless by design. There's no seed phrase to expose. The MPC architecture means your private key is never reconstructed—not even during signing.

</details>

### Fees

<details>
<summary><strong>What fees do I pay?</strong></summary>

You pay two types of fees:
1. **Plugin fees** — Set by the developer (shown before installation)
2. **Network fees** — Standard blockchain gas fees for each transaction

There are no hidden Vultisig platform fees.

</details>

<details>
<summary><strong>How do subscription fees work?</strong></summary>

If a plugin uses subscription pricing, the fee is collected automatically at the start of each billing period. You can cancel anytime—you'll retain access until the current period ends.

</details>

***

## What's Next

The Marketplace is actively growing. Upcoming features include:

- More automation plugins from third-party developers
- Advanced DeFi integrations (yield farming, liquidity provision)
- Cross-chain yield strategies
- Community-built tools
- Enhanced analytics and reporting

Want to request a plugin or feature? Join the [Discord](https://discord.gg/vultisig) and share your ideas.
