---
description: >-
  Send or receive your first transaction in Vultisig. Check the address
  before you sign.
---

# Your first Vultisig transaction

Once the vault exists and you have a backup, you can receive and send.

***

## Receive

1. Open your vault
2. Tap the asset you want to receive
3. Tap **Receive** or the QR icon
4. Copy the address, or show the QR to the sender

{% hint style="info" %}
Each chain has its own address. A Bitcoin address will not receive Ethereum.
{% endhint %}

You can deposit from another wallet, an exchange, or a friend.

***

## Send

1. Select the asset you want to send
2. Tap **Send**
3. Paste the recipient, scan a QR, or pick a saved address
4. Enter the amount, or tap **Max**
5. Check destination, amount, and fee
6. Sign with your devices
7. The app broadcasts after a successful signature

Check the address before you sign. Once it is on-chain, you cannot undo it. For a large first send to a new address, send a small test first. Leave a little native token in the vault for the network fee (ETH, BTC, SOL, and so on).

Full send path: [Send crypto from Vultisig](../app-guide/wallet/sending.md).

***

## Signing

On a Fast Vault, you approve on your device and VultiServer co-signs. On a Secure Vault, your other device has to join.

1. Start the transaction on one device
2. Pair the other device(s) by QR or relay
3. Every device shows the same destination, amount, and (on EVM) decoded function
4. Each device approves on its own
5. The shares combine into one signature
6. The app broadcasts

No device holds the complete key. One compromised device cannot move funds.

[How Vultisig signs a transaction](../security-and-technology/keysign.md)

***

## Next

| Action | Guide |
| -------- | ------------------------------------------------- |
| **Send** | [Sending](../app-guide/wallet/sending.md) |
| **Swap** | [Swapping](../app-guide/wallet/swapping.md) |
| **DeFi** | [DeFi](../app-guide/defi/) |

***

## If something goes wrong

**Stuck after broadcast.** Check the chain explorer. Busy networks and low fees delay confirmation.

**Signing failed.** Same Wi-Fi or the relay, matching vault name, latest app. Try the sign again.

- [FAQ](../help/faq.md)
- [Discord](https://discord.gg/Cugw9T2NrP)
