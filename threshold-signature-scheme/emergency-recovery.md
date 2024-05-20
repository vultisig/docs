---
description: What happens if the TSS software is unavailable
---

# Emergency Recovery

It is possible to re-combine the vault shares of a vault and produce the private key to extract it and use it in other wallets.\
Scripts are produced to enable users to do this should the Vultisig App ever go offline:

[https://github.com/vultisig/mobile-tss-lib/tree/main/cmd/recovery-web](https://github.com/vultisig/mobile-tss-lib/tree/main/cmd/recovery-web)

{% hint style="danger" %}
Before the user does this, the private key has never existed. \
This is a one-way function, once the private key is produced, the vault is no longer a TSS vault, and is now a single-signature wallet.

Do not ever use a single-signature wallet again.
{% endhint %}
