---
description: What happens if the TSS software is unavailable
---

# Emergency Recovery

It is possible to recombine the vault shares of a vault and generate the private key to extract it and use it in other wallets. Scripts are created to allow users to do this if the Vultisig app ever goes offline:

[https://github.com/vultisig/mobile-tss-lib/tree/main/cmd/recovery-web](https://github.com/vultisig/mobile-tss-lib/tree/main/cmd/recovery-web)

{% hint style="danger" %}
Before the user does this, the private key never existed. This is a one-way function; once the private key is created, the vault is no longer a TSS vault, but a single-signature wallet.

Never use a single-signature wallet again.
{% endhint %}
