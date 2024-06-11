---
description: Actions
---

# TSS Actions

1. Key-Generation
2. Key-Signing
3. Re-Sharing

## Key Generation

Key generation (Key-gen) is a process by which devices prove access to a secret and generate a shared public key. This requires 100% of the devices to be online. The shared public key generated through this process can then be used to create on-chain addresses, referred to as the "vault," for receiving funds.

These vault addresses resemble simple wallet addresses, known as "externally-owned accounts" (EOA), and do not appear as special contracts or scripts.

{% hint style="success" %}
**Importantly:** The individual secrets (vault shares) do not contain funds. In the normal vault user flow, the actual private key to the vault never exists. Consequently, individual vault shares can be safely emailed, saved, uploaded to websites, and more.

Furthermore, no part of a vault share indicates the location of the final vault. There is no linkage between a vault share and the final vault, ensuring that someone who discovers a single vault share will not know its purpose or where to look. This design guarantees that the privacy and security of the vault are maintained even if a vault share is exposed.
{% endhint %}

{% hint style="danger" %}
It is crucial that individual Vault shares are never stored together, as a malicious party could potentially recombine the Vault shares to access the Vault:

1. Do not back up more than one Vault share to the same device, email, Google Drive, or iCloud.
2. Do not upload more than one Vault share to the same website.

Following these guidelines ensures the security and integrity of the Vault shares and prevents unauthorized access to the Vault.
{% endhint %}

## Key Signing

Key Signing is a process in which a threshold of parties, specifically 67% as chosen by Vultisig TSS, must collaborate in a coordinated manner to prove access to a secret and generate a signed transaction object. This transaction can then be submitted to a crypto network.

For instance, in a 2-of-3 TSS vault, only 2 of the 3 devices need to participate to sign an outgoing transaction. This ensures that the signing process remains secure and efficient while maintaining the integrity and security of the vault.

## Re-share

Re-share is a process that allows the signing threshold of a vault to adapt by replacing a non-responsive or lost devices with a new one or by modifying the vault's configuration.

This can include upgrading from a 2-of-2 configuration to a 3-of-4, or even downgrading as needed. This flexibility ensures that the vault remains functional and secure despite changes in device availability or requirements.
