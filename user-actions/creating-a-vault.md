---
description: How to create a vault.
---

# Creating a Vault

### Setup

You first need to download the Vultisig App to 2 or more devices.

{% hint style="warning" %}
Currently only Macbook (M1), iPad and iPhones are supported. Find the app in the app stores (for Mac, look under "iPad Apps")
{% endhint %}

### Vault Types

The vaults will be m-of-n, where m is at least 2/3rds of n, and no maximum number of n devices. The more devices you use, the longer it will take to process any transactions.

The following are the most common vaults:

1. **2 of 2 vault** - only need two to create a vault and two to sign a transaction. Note, it is not automatically "redundant" so you absolutely should export the vault shares and store them separately and securely. This will be the most popular vault type as this is the most convenient. **For its convenience, this is also the least secured vault type.**
2. **2 of 3 vault** - three devices to create a vault and two to sign a transaction. This is automatically backed (one device is the backup) so you don't need to export vault shares. But you may choose to do this. **This vault type is recommended, as it is more secured than 2 of 2 vault.**
3. **3 of 4 vault** - four devices to create a vault and three to sign a transaction. This is automatically backed (one device is the backup) so you don't need to export vault shares. But you may choose to do this.

### Creating A Vault

{% hint style="success" %}
Remember - all your devices must be open with Vultisig, and must be connected on the (i) same wifi network/ (ii) same hotspot network/ (iii) individual cellular network. If they are not fully connected, they will fail the KeyGen.
{% endhint %}

Get your devices ready and create a vault.

**Main Device:** START -> will show a QR Code to pair with

**Pair Device:** JOIN -> will start the camera to scan the QR code

<figure><img src="../.gitbook/assets/CreateVault-1.png" alt="" width="188"><figcaption></figcaption></figure>

#### Network Type

You can choose Wifi, Hotspot or Cellular.

1. **WiFi**: Simplest and fastest, however may not work on some WiFi networks (since they block mDNS packets)
2. **Hotspot**: Fast and reliable, since devices are connected directly to each other. Some devices may not have Hotspot.
3. **Cellular**: slowest, and can be unreliable. Encrypted packets are routed through a relay server.

<figure><img src="../.gitbook/assets/CreateVault-2.png" alt="" width="188"><figcaption></figcaption></figure>

### Keygen

Once clicking **CONTINUE**, the keygen process will begin. Firstly it will create the pre-params (your vault shares and some other aspects, around 10 seconds), then it will create the ECDSA and EdDSA keys (another 10 seconds). Finally it will be done! Make sure all devices show the done screen.

### Troubleshooting

If Keygen fails it is because you have an unreliable network and the devices dropped connections.

1. Quit the apps and start again
2. Change networks
