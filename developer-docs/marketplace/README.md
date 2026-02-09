# Marketplace Plugins

Build automation plugins that run on Vultisig's infrastructure and distribute them through the Marketplace. Plugins leverage Vultisig's MPC technology to execute transactions on behalf of users while maintaining self-custody—users never give up their private keys.

## Getting Started

1. **Understand the architecture**: Learn [what a plugin is](infrastructure-overview/plugins.md) and how the [services](infrastructure-overview/services.md) work together
2. **Build your plugin**: Follow the [quick start guide](create-a-plugin/basics-quick-start.md) and reference the [App Recurring](https://github.com/vultisig/app-recurring) plugin
3. **Submit for review**: Complete the [submission process](create-a-plugin/submission-process.md) to list on the Marketplace

## Documentation

### Architecture
- [What is a Plugin](infrastructure-overview/plugins.md) - Core concepts and security model
- [Services Architecture](infrastructure-overview/services.md) - HTTP Server, Worker, Scheduler, TX Indexer
- [Policy Rules](infrastructure-overview/metarules.md) - MetaRules and Direct Rules for transaction validation
- [Infrastructure Overview](infrastructure-overview/infrastructure.md) - System components and interactions

### Building
- [Quick Start](create-a-plugin/basics-quick-start.md) - Step-by-step guide to your first plugin
- [Build Your Plugin](create-a-plugin/build-your-plugin/README.md) - Detailed developer guide
- [Adding a New Chain](create-a-plugin/build-your-plugin/adding-a-new-chain.md) - Extend chain support

### Publishing
- [Submission Process](create-a-plugin/submission-process.md) - Review and approval workflow
- [Revenue](infrastructure-overview/revenue.md) - Fee structures and 70/30 revenue split

## Reference Implementation

The [App Recurring](https://github.com/vultisig/app-recurring) plugin is the official reference implementation. Clone it as a starting point for your own development.

Each plugin is an independent service — you define your own recipe (transaction rules), authentication flow, and business logic. The reference plugin shows integration patterns, but your plugin can be structured however you need.

## Support

- **Discord**: [discord.gg/vultisig](https://discord.gg/vultisig) - Join the third-party developer section
- **GitHub**: [github.com/vultisig](https://github.com/vultisig) - Source code and issues
