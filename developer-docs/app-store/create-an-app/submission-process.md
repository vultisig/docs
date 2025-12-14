# Submission Process

### TL;DR

Submit your Vultisig app by joining Discord's third-party developer section, preparing a `plugin-config.yaml` with app details (ID, endpoint, supported chains, fees), and meeting security/performance requirements. The review includes basic sanity checks (documentation, config validity) for marketplace listing, with optional "Vultisig Approved" comprehensive audits available. Developers deploy and maintain their own infrastructure post-approval. Reference apps include DCA for recurring buy automation and other standard/AI-agent categories.

***

## Overview

To publish your app in the Vultisig ecosystem, you must submit your app/agent for review and approval.&#x20;

After approval, developers are responsible for deploying and maintaining their own applications, ensuring uptime, scaling, and operational security outside of Vultisig core infrastructure.

{% hint style="info" %}
**Note:** Verifier and Fees are managed exclusively by Vultisig and run as core infrastructure services. All other apps are deployed, operated, and updated directly by their respective developers.
{% endhint %}

{% hint style="danger" %}
Vultisig reserves the right to delist apps that fail to perform or adhere to Vultisig's quality standards, or that introduce a risk of harming users.
{% endhint %}

***

## Official App Examples

Below are typical app types found in Vultisig’s ecosystem. Refer to them when designing your own submission:

| App ID            | Title         | Description                      | Category | Endpoint                                         |
| ----------------- | ------------- | -------------------------------- | -------- | ------------------------------------------------ |
| vultisig-dca-0000 | Recurring buy | Dollar Cost Averaging automation | app      | https://apps.vultisig.com/apps/vultisig-dca-0000 |

App Store supports both standard apps with business logic or automation, and agent-based apps for advanced trading and verification tasks.

***

## Submission Requirements

Create a detailed `plugin-config.yaml` describing:

* ID, title, description, and endpoint (if applicable)
* Category (app or ai-agent)
* Supported blockchains
* Payment requirements

Follow this example:

```
plugins:
  - id: vultisig-dca-0000
    title: DCA Plugin
    description: Dollar Cost Averaging automation plugin
    server_endpoint: https://dca.vultisigplugin.app
    category: plugin
```

***

## Submission Process

{% stepper %}
{% step %}
### Join the developer discord

Join Discord at [discord.gg/vultisig](https://discord.gg/vultisig), navigate to dedicated section for third-party developers and get real-time support and feedback from the Vultisig team while building your app
{% endstep %}

{% step %}
### Prepare submission

Prepare all necessary code and documentation, like YAML config and documentation.
{% endstep %}

{% step %}
### Complete Checks

Complete the security checklist and performance requirements for APIs and resource use.
{% endstep %}

{% step %}
### <mark style="color:orange;">Review and approval</mark>

The review process includes basic sanity checks covering documentation completeness, configuration validity, and automation structure. \
Apps that pass these checks can be listed in the App Store.&#x20;
{% endstep %}

{% step %}
### <mark style="color:$info;">Optional: "Vultisig Approved"</mark>

Developers seeking a "Vultisig Approved" badge can request a comprehensive audit, which includes thorough security review, code audit, and performance testing. \
This optional deep audit requires additional resources and time. Use the example apps above as references for best practices and compliance.
{% endstep %}

{% step %}
### <mark style="color:$success;">App Store Listing</mark>

Once approved, your app is provisioned and listed in the Vultisig marketplace alongside other available Apps. Those examples illustrate category standards, endpoint conventions, and API scope for production apps.

Revenue sharing models are available for fee, subscription, and premium features, negotiated during approval. Plugin developers receive **70%** of all generated revenue, while **30%** is allocated to the $VULT token treasury to support ecosystem development and maintenance.
{% endstep %}
{% endstepper %}

***

## Support

Questions? Reach out via email, documentation portal, or the developer Discord for guidance on submission and best practices.

***
