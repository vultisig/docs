# Marketplace Plugins (paused)

{% hint style="warning" %}
**Paused.** The plugin marketplace is not in production. These pages describe the planned product and stay for reference.
{% endhint %}

These pages describe the planned plugin product. Plugins would run on Vultisig infrastructure and sign through the vault. Users do not hand over keys.

## Getting Started

1. **Understand the architecture**: Learn [what a plugin is](plugins.md) and how the [services](services.md) work together
2. **Build your plugin**: Follow the [quick start guide](basics-quick-start.md) and reference the [App Recurring](https://github.com/vultisig/app-recurring) plugin
3. **Submit for review**: Complete the [submission process](submission-process.md) to list on the Marketplace

## Documentation

### Architecture

* [What is a Plugin](plugins.md) - Core concepts and security model
* [Services Architecture](services.md) - HTTP Server, Worker, Scheduler, TX Indexer
* [Policy Rules](metarules.md) - MetaRules and Direct Rules for transaction validation
* [Infrastructure Overview](infrastructure.md) - System components and interactions

### Building

* [Quick Start](basics-quick-start.md) - Step-by-step guide to your first plugin
* [Build Your Plugin](build-your-plugin/) - Detailed developer guide
* [Adding a New Chain](build-your-plugin/adding-a-new-chain.md) - Extend chain support

### Publishing

* [Submission Process](submission-process.md) - Review and approval workflow
* [Revenue](infrastructure-overview/revenue.md) - Fee structures and 70/30 revenue split

## Reference Implementation

The [App Recurring](https://github.com/vultisig/app-recurring) plugin is the official reference implementation. Clone it as a starting point for your own development.

Each plugin is an independent service — you define your own recipe (transaction rules), authentication flow, and business logic. The reference plugin shows integration patterns, but your plugin can be structured however you need.

## Support

* **Discord**: [discord.gg/Cugw9T2NrP](https://discord.gg/Cugw9T2NrP) - Join the third-party developer section
* **GitHub**: [github.com/vultisig](https://github.com/vultisig) - Source code and issues
