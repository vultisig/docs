---
description: >-
  Choose between autonomous, policy-bound, and human-approved agent signing
  models with Vultisig.
---

# Authority and security

MPC removes the complete private key as a single point of failure. It does not make an agent's decisions correct. Choose a signing model based on the value at risk, the workflow's speed requirements, and the controls that must remain outside the agent.

## Compare signing models

### Autonomous Fast Vault

A Fast Vault combines the agent-controlled device share with a VultiServer share. The agent can sign without a person present.

**Use when:** the workflow requires unattended execution and the vault can be funded according to a defined risk budget.

**Primary controls:** isolated vaults, limited funding, protected credentials, transaction previews, monitoring, and tested backups.

### Policy-bound Marketplace automation

A Marketplace plugin proposes unsigned transactions. A separate Verifier checks each proposal against the user's configured policy before participating in MPC signing.

**Use when:** an autonomous strategy should operate inside explicit limits, such as allowed assets, destinations, amounts, or timing.

**Primary controls:** narrowly scoped rules, independent validation, plugin review, monitoring, and the ability to disable the automation.

### Human-approved Secure Vault

A Secure Vault requires the configured threshold of devices to join each signing session. The agent can research and prepare an action, but cannot complete the signature alone.

**Use when:** every transaction requires human review or multiple participants must approve.

**Primary controls:** transaction review, device separation, threshold configuration, backups, and clear approval procedures.

## Responsibilities that remain with the developer

Your application remains responsible for:

* Prompt and tool-call security
* Strategy logic and economic outcomes
* Token, contract, route, and destination selection
* Transaction parameters and fee limits
* Vault passwords, backups, logs, and runtime credentials
* Monitoring, incident response, and disabling compromised agents

## Recommended controls

* Start with read-only portfolio access.
* Require previews before enabling signing.
* Use separate vaults for separate agents or strategies.
* Fund autonomous vaults only to an explicit risk budget.
* Prefer policy-bound automation when rules can constrain the workflow.
* Prefer Secure Vaults for high-value or irreversible actions.
* Treat arbitrary message and byte signing as high-risk capabilities.
* Never embed vault passwords or backups in prompts, source code, or logs.

{% hint style="warning" %}
Fast Vaults enable unattended signing. Only grant an agent the assets and capabilities its workflow requires.
{% endhint %}

