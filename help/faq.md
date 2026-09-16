---
description: >-
  Vultisig FAQ: vaults, backups, fees, signing, and recovery.
---

# Vultisig FAQ

## General

### What is Vultisig?

Vultisig is a vault for many chains with no seed phrase. Your key is split into shares on your devices. They sign together. One share cannot move funds.

### Which platforms support Vultisig?

Vultisig is available on:

* iOS (iPhone, iPad)
* Android
* macOS
* Windows
* Linux
* Browser Extension (Chrome, Brave, Edge, Firefox). Full vault.

### Which blockchains does Vultisig support?

Vultisig supports 36+ chains. The named list is on the [Wallet tab](../app-guide/wallet/) page.

***

## Vaults

### What is the difference between Fast Vault and Secure Vault?

**Fast Vault:** One device. VultiServer holds the second share and signs with you. It can never sign alone.

**Secure Vault:** Two or more devices you hold. No VultiServer. 2-of-3 is the usual setup — you can lose one device and still sign.

### Can I convert a Fast Vault to a Secure Vault?

No. Create a new Secure Vault and move the funds across. The Fast Vault stays as it is.

### What happens if I lose a device?

For **Fast Vault**: restore from your device backup. VultiServer is still the second share.

For **Secure Vault**: on 2-of-3 you can still sign with the remaining two, then reshare to add a replacement device.

***

## Backup & Recovery

### How do I backup my vault?

Export a backup from each device in the vault. Keep those files apart — anyone with enough shares can move funds.

On Fast Vault, the device backup matters most. You can re-request the server share with your password.

See [Backup & Recovery](../getting-started/backup-recovery.md).

### What if I lose my backup?

If you still have enough devices, export new backups from them. If the devices and the backups are both gone, the funds cannot be recovered. Vultisig cannot reconstruct them.

### Can Vultisig recover my funds?

No. Vultisig has no access to your vault shares and cannot reconstruct them.

***

## Transactions

### Why is my transaction taking so long?

Transaction times depend on:

* Blockchain network congestion
* Gas/fee settings
* Number of confirmations required

For TSS signing delays, ensure all participating devices have stable network connections.

### What fees does Vultisig charge?

Sends charge network fees only. Vultisig does not add a send fee.

Swaps charge a 50 bps Vultisig fee on top of the route's protocol and network fees. Hold $VULT in the vault to reduce that fee. Tiers are on [$VULT in-app fee discounts](../vultisig-token/vult/in-app-utility.md).

### Can I cancel a pending transaction?

Once broadcast to the blockchain, transactions cannot be cancelled. Some networks support "replace-by-fee" to override pending transactions with higher fees.

***

## Security

### Is Vultisig open source?

Yes. All Vultisig code is open source and available on [GitHub](https://github.com/vultisig).

### Has Vultisig been audited?

The DKLS23 signing library was audited by Trail of Bits. See [Audits](security.md). Application audit reports will be linked there when published.

### What if Vultisig company disappears?

Your funds remain accessible. Vultisig provides [emergency recovery tools](../security-and-technology/emergency-recovery.md) to reconstruct private keys from vault shares if the software becomes unavailable.

***

## Troubleshooting

### Devices not connecting for signing

1. Ensure all devices are on the same network (or using relay server)
2. Check that vault names match exactly
3. Verify devices have the latest app version
4. Try restarting the signing process

### QR code not scanning

1. Ensure adequate lighting
2. Clean camera lens
3. Hold device steady at appropriate distance
4. Try adjusting screen brightness on displaying device

### App crashes or freezes

1. Force close and restart the app
2. Check for app updates
3. Restart device if issue persists
4. Report persistent issues on [GitHub](https://github.com/vultisig/vultisig-ios/issues) or [Discord](https://discord.gg/Cugw9T2NrP)

***

## Related

* [Getting Started](../getting-started/overview.md)
* [Backup & Recovery](../getting-started/backup-recovery.md)
* [Security](security.md)
