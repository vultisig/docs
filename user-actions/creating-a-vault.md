---
description: How to create a vault.
---

# Creating a Vault

### Setup

You first need to download the Voltix App to 2 or more devices.

{% hint style="warning" %}
Currently only Macbook (M1), iPad and iPhones are supported. Find the app in the app stores (for Mac, look under "iPad Apps")
{% endhint %}

### Vault Types

The vaults will be m-of-n, where m is at least 2/3rds of n, and no maximum number of n devices. The more devices you use, the longer it will take to process any transactions.

The following are the most common vaults:

1. **2 of 2 vault** - only need two to create a vault and two to sign a transaction. Note, it is not automatically "redundant" so you absolutely should export the vault shares and store them separately and securely. This will be the most popular vault type since it is the most convenient.
2. **2 of 3 vault** - three devices to create a vault and two to sign a transaction. This is automatically backed (one device is the backup) so you don't need to export vault shares. But you may choose to do this.
3. **3 of 4 vault** - four devices to create a vault and three to sign a transaction. This is automatically backed (one device is the backup) so you don't need to export vault shares. But you may choose to do this.

### Steps

1. On all your devices, click **Create a New Vault**. Note: if your devices already have existing vaults, first click **Add New Vault** to access this page. ![IMG\_1298](https://github.com/SamYap0/Voltix-docs/assets/96066776/860ff2e0-5b46-417f-9e71-e04d1fbaa88e)
2. On one of your device, click **Start**. ![IMG\_1299](https://github.com/SamYap0/Voltix-docs/assets/96066776/e27cb5a7-b282-4dd3-a0a2-8787da1e1ae3)
3. On the next screen, select **WiFi**, **Hotspot** or **Cellular**. If using WiFi or Hotspot, all devices must be on the same network. ![IMG\_1300](https://github.com/SamYap0/Voltix-docs/assets/96066776/2315f8f3-c9f5-4eb0-8f22-ed539a88a7f1)
4. _On the other devices_, click **Join**. ![IMG\_1299](https://github.com/SamYap0/Voltix-docs/assets/96066776/edadcd17-576e-4c83-adb6-42516a2c4ecf)
5. The camera function will be enabled, and use it to scan the QR Code displayed on the _first device_. Once joined, it will show the "Waiting for key generation to start" message. ![IMG\_1301](https://github.com/SamYap0/Voltix-docs/assets/96066776/7088ed64-a8a0-4965-bd1e-7630bf857707)
6. _On the first device_, select the other devices that showed up, then click **Continue**. ![IMG\_6256](https://github.com/SamYap0/Voltix-docs/assets/96066776/f56779ad-fb6d-495e-95ae-4668f7576559)
7. On the Summary page, check that the details are correct, then click **Continue**. ![IMG\_6257](https://github.com/SamYap0/Voltix-docs/assets/96066776/022713bd-c503-41db-859c-a2dd0b2ac52d)
8. The vault key generation process will start and once completed, your vault/address will be ready for use.

Note: The above screenshots illustrate a 2-of-2 vault creation, but the process is similar for any m-of-n vaults.

### Creating A Vault

Get your devices ready and create a vault.&#x20;

**Main Device:** START -> will show a QR Code to pair with

**Pair Device:** JOIN -> will start the camera to scan the QR code

<figure><img src="../.gitbook/assets/image (3).png" alt="" width="188"><figcaption></figcaption></figure>

#### Network Type

You can choose Wifi, Hotspot or Cellular.&#x20;

1. **WiFi**: Simplest and fastest, however may not work on some WiFi networks (since they block mDNS packets)
2. **Hotspot**: Fast and reliable, since devices are connected directly to each other. Some devices may not have Hotspot.
3. **Cellular**: slowest, and can be unreliable. Encrypted packets are routed through a relay server.&#x20;

<figure><img src="../.gitbook/assets/image (4).png" alt="" width="188"><figcaption></figcaption></figure>

### Troubleshooting

If Keygen fails it is because you have an unreliable network and the devices dropped connections.&#x20;

1. Quit the apps and start again
2. Change networks
