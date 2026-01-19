# Infrastructure

## Services and Modules

![Vultisig ecosystem](<../../../.gitbook/assets/Services and Modules.jpg>)

<table><thead><tr><th align="center">Component</th><th align="center" valign="middle">Role / Description</th></tr></thead><tbody><tr><td align="center"><strong>Recipes</strong></td><td align="center" valign="middle">Shared library and integration framework for standardized blockchain logic and protocol modules. Provides building blocks for apps and services to interact securely and consistently across supported blockchains.</td></tr><tr><td align="center"><strong>Fees</strong></td><td align="center" valign="middle">Fee calculation and treasury logic; enforces fee-related rules and automates processing for transactions and treasury operations. Utilized by apps and Verifier automations.</td></tr><tr><td align="center"><strong>Verifier</strong></td><td align="center" valign="middle">Centralizes automation management. Receives transaction requests from apps, validates them against custom automations, and performs TSS-based signing for compliant transactions. Does NOT initiate or broadcast transactions; strictly enforces rules and maintains security boundaries.</td></tr><tr><td align="center"><strong>App A/B/...</strong></td><td align="center" valign="middle">Applications that implement custom business logic or user-facing features. Compose and initiate transactions and handle broadcasting. Submit transactions to the Verifier for validation and signing, ensuring automation compliance.</td></tr><tr><td align="center"><strong>Blockchains</strong></td><td align="center" valign="middle">Supported external networks such as Bitcoin, Ethereum, Solana, etc. Recipes standardizes their interfaces for system-wide compatibility.</td></tr></tbody></table>

## Key Interactions

* All applications interface with Recipes modules for standardized logic.
* Fee computations and treasury operations are funneled through the Fees module.
* Every transaction is routed to the Verifier for strict automation compliance and TSS signing.
* Blockchains are kept external, with all interaction managed through the Recipes abstraction layer.
