---
description: >-
  Frequently asked questions about Vultisig. Common questions about vaults,
  backups, transactions, and troubleshooting.
---

# FAQ

## General

### What is Vultisig?

Vultisig is a multi-chain, multi-platform crypto vault that uses threshold signatures (TSS) to secure digital assets. Unlike traditional wallets with single private keys, Vultisig distributes security across multiple devices—no single point of failure.

### Which platforms support Vultisig?

Vultisig is available on:
- iOS (iPhone, iPad)
- Android
- macOS
- Windows
- Linux
- Browser Extension (Chrome, Brave, Edge)

### Which blockchains does Vultisig support?

Vultisig supports 30+ chains including Bitcoin, Ethereum, THORChain, Solana, Cosmos ecosystem chains, and many EVM-compatible networks. See [Supported Chains](../app-guide/wallet/README.md) for the complete list.

***

## Vaults

### What is the difference between Fast Vault and Secure Vault?

**Fast Vault**: Uses your device plus Vultisig's server as the second signer. Quick setup, instant transactions, single-device convenience. Best for everyday use.

**Secure Vault**: Requires multiple physical devices you control. Higher security with no server dependency. Best for significant holdings.

### Can I convert a Fast Vault to a Secure Vault?

Not directly. You would need to create a new Secure Vault and transfer your assets. Both vault types use the same underlying TSS technology.

### What happens if I lose a device?

For **Fast Vault**: Restore from your device backup using the Vultiserver as the second share.

For **Secure Vault**: With a 2-of-3 configuration, you can still sign transactions with the remaining two devices and reshare to add a replacement device.

***

## Backup & Recovery

### How do I backup my vault?

Export vault shares from each device in your vault configuration. Store backups securely—anyone with threshold shares can access your funds.

For Fast Vault, device backup is the priority since server backup can be re-requested with your password.

See [Backup & Recovery](../getting-started/backup-recovery.md) for detailed instructions.

### What if I lose my backup?

If you still have access to threshold devices, you can re-export backups. If you've lost access to threshold devices AND backups, funds cannot be recovered—this is the security tradeoff of true self-custody.

### Can Vultisig recover my funds?

No. Vultisig has no access to your vault shares. This is a feature, not a limitation—it ensures true self-custody.

***

## Transactions

### Why is my transaction taking so long?

Transaction times depend on:
- Blockchain network congestion
- Gas/fee settings
- Number of confirmations required

For TSS signing delays, ensure all participating devices have stable network connections.

### What fees does Vultisig charge?

Vultisig charges no fees for basic operations. You only pay standard blockchain network fees for transactions.

Swap operations through integrated DEXs (THORChain, Maya) include their standard protocol fees.

### Can I cancel a pending transaction?

Once broadcast to the blockchain, transactions cannot be cancelled. Some networks support "replace-by-fee" to override pending transactions with higher fees.

***

## Security

### Is Vultisig open source?

Yes. All Vultisig code is open source and available on [GitHub](https://github.com/vultisig).

### Has Vultisig been audited?

Yes. See [Security](security.md) for audit reports and security practices.

### What if Vultisig company disappears?

Your funds remain accessible. Vultisig provides [emergency recovery tools](../security-technology/emergency-recovery.md) to reconstruct private keys from vault shares if the software becomes unavailable.

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
4. Report persistent issues on [GitHub](https://github.com/vultisig/vultisig-ios/issues) or [Discord](https://discord.vultisig.com)

***

## Related

- [Getting Started](../getting-started/README.md)
- [Backup & Recovery](../getting-started/backup-recovery.md)
- [Security](security.md)
