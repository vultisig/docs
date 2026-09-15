# Build on Vultisig

SDK, CLI, Extension, and (paused) Marketplace plugins. Same threshold vault the apps use.

## Surfaces

| Surface | Example |
| --------------------------- | ------------------------------------------------------------ |
| **Agents** | Portfolio, send, swap, sign. Authority is the Fast / policy / human choice |
| **Trading scripts** | DCA, rebalance, scheduled send |
| **dApp connect** | Extension `window.vultisig` / EIP-1193 |
| **Treasury** | N-of-M Secure Vault, programmatic send via SDK/CLI |
| **Portfolio** | Balances and txs across the vault's chains |

## Paths

### AI Agents

Agents get a vault. You choose Fast (autonomous), policy-bound (paused), or human-approved Secure Vault.

* **Languages**: TypeScript, shell, or Go

{% content-ref url="ai-agents/" %}
[ai-agents](ai-agents/)
{% endcontent-ref %}

### Marketplace Plugins

{% hint style="warning" %}
**Paused.** The plugin marketplace is not in production. These pages describe the planned product and stay for reference.
{% endhint %}

Planned Go plugins on Vultisig infrastructure. Not in production.

* **Language**: Go
* **Planned revenue**: 70/30 split (developer/treasury)

{% content-ref url="marketplace/" %}
[marketplace](marketplace/)
{% endcontent-ref %}

### SDK

Create and manage vaults from TypeScript.

* **Language**: TypeScript

{% content-ref url="vultisig-sdk/" %}
[vultisig-sdk](vultisig-sdk/)
{% endcontent-ref %}

### Extension Integration

Connect a dApp to the Vultisig Extension.

* **Language**: JavaScript

{% content-ref url="vultisig-extension-integration-guide.md" %}
[vultisig-extension-integration-guide.md](vultisig-extension-integration-guide.md)
{% endcontent-ref %}

## Signing model

* [How TSS Works](../security-and-technology/tss-actions.md) - Threshold signature scheme overview
* [How DKLS23 Works](../security-and-technology/how-dkls23-works.md) - The signing protocol
* [Vultiserver](../vultisig-infrastructure/what-is-vultisigner/) - Fast vault infrastructure

## Help

* **Discord**: [discord.gg/Cugw9T2NrP](https://discord.gg/Cugw9T2NrP) - Dedicated third-party developer section
* **GitHub**: [github.com/vultisig](https://github.com/vultisig) - Source code and issue tracking
