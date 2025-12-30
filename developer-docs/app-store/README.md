# App Store Apps

Build automation apps that run on Vultisig's infrastructure and distribute them through the App Store marketplace. Apps leverage Vultisig's MPC technology to execute transactions on behalf of users while maintaining self-custody—users never give up their private keys.

## Getting Started

1. **Understand the architecture**: Learn [what an app is](infrastructure-overview/apps.md) and how the [services](infrastructure-overview/services.md) work together
2. **Build your app**: Follow the [quick start guide](create-an-app/basics-quick-start.md) and reference the [DCA app](https://github.com/vultisig/dca)
3. **Submit for review**: Complete the [submission process](create-an-app/submission-process.md) to list on the App Store

## Documentation

### Architecture
- [What is an App](infrastructure-overview/apps.md) - Core concepts and security model
- [Services Architecture](infrastructure-overview/services.md) - HTTP Server, Worker, Scheduler, TX Indexer
- [Policy Rules](infrastructure-overview/metarules.md) - MetaRules and Direct Rules for transaction validation
- [Infrastructure Overview](infrastructure-overview/infrastructure.md) - System components and interactions

### Building
- [Quick Start](create-an-app/basics-quick-start.md) - Step-by-step guide to your first app
- [Build Your App](create-an-app/build-your-app/) - Detailed developer guide
- [Adding a New Chain](create-an-app/build-your-app/adding-a-new-chain-to-the-vultisig-app-ecosystem.md) - Extend chain support

### Publishing
- [Submission Process](create-an-app/submission-process.md) - Review and approval workflow
- [Revenue](infrastructure-overview/revenue.md) - Fee structures and 70/30 revenue split

## Reference Implementation

The [DCA (Dollar Cost Averaging)](https://github.com/vultisig/dca) app is the official reference implementation. Clone it as a starting point for your own development.

## Support

- **Discord**: [discord.gg/vultisig](https://discord.gg/vultisig) - Join the third-party developer section
- **GitHub**: [github.com/vultisig](https://github.com/vultisig) - Source code and issues
