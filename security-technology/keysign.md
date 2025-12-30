---
description: >-
  How Vultisig keysigning works. Multi-device threshold signing process for
  Fast Vaults and Secure Vaults. QR pairing, verification, and broadcast.
---

# Keysign

Keysigning is the process by which threshold devices collaborate to sign a transaction. This page explains how signing works for both Fast Vaults and Secure Vaults.

<figure><img src="../.gitbook/assets/image (4).png" alt="" width="375"><figcaption></figcaption></figure>

***

## How It Works

<figure><img src="../.gitbook/assets/How keysign works.png" alt=""><figcaption><p>Transaction signing flowchart</p></figcaption></figure>

### Session Initiation

A device initiates a cryptographic session to sign a transaction. It transforms user input into the transaction payload while acting as the session host. The initiating device sends session metadata (including session ID) to the Vultisig relay server or broadcasts it over the local network. A QR code is generated containing session-specific details for pairing devices.

### Device Pairing

Pairing devices scan the QR code to join the session using the embedded session ID and encrypted hex chain code. The initiating device monitors joining devices and initiates the keysigning ceremony when the required threshold is reached.

### Signing Ceremony

During the keysigning ceremony, participating devices jointly sign the transaction using threshold signature cryptography. No single device ever has access to the complete private key—only the final signature is assembled.

### Broadcast

Upon successful completion, the initiating device propagates the signed transaction to the blockchain and distributes the transaction hash to other participating devices for verification.

***

## Fast Vault Signing

Fast Vaults provide a single-device signing experience with the Vultiserver as the automatic co-signer.

### Steps

1. Prepare the transaction (Send or Swap)
2. Tap **Continue** or **Swap**

<div><figure><img src="../.gitbook/assets/Send.png" alt="" width="188"><figcaption><p>Send screen</p></figcaption></figure> <figure><img src="../.gitbook/assets/Swap-screen.png" alt="" width="188"><figcaption><p>Swap screen</p></figcaption></figure></div>

3. Verify transaction details

<div><figure><img src="../.gitbook/assets/Verify pair.png" alt="" width="188"><figcaption><p>Send verification</p></figcaption></figure> <figure><img src="../.gitbook/assets/verify-swap.png" alt="" width="188"><figcaption><p>Swap verification</p></figcaption></figure></div>

4. Enter your password
5. Wait for signing to complete

<figure><img src="../.gitbook/assets/Succes-swap (1).png" alt="" width="188"><figcaption></figcaption></figure>

***

## Secure Vault Signing

Secure Vaults require threshold devices to physically participate in signing.

### Video Guide

[![Signing Tutorial](../.gitbook/assets/TwitterVideoThumbnail.jpeg)](https://twitter.com/iceman00008/status/1824690627843592573/video/1)

### Preparation

Depending on your vault configuration (2-of-2, 2-of-3, 3-of-4, or m-of-n), have the threshold number of devices ready.

{% hint style="info" %}
One device creates the transaction, other devices verify and co-sign. Both devices will co-sign but only one transaction is sent—no duplicates.
{% endhint %}

### On the Main Device

1. Prepare the transaction (Send, Swap, or DeFi action)
2. Tap **Continue** or **Swap**

<div><figure><img src="../.gitbook/assets/Send.png" alt="" width="188"><figcaption><p>Send screen</p></figcaption></figure> <figure><img src="../.gitbook/assets/Swap-screen.png" alt="" width="188"><figcaption><p>Swap screen</p></figcaption></figure></div>

3. Verify all details on the confirmation screen
4. Check reminder boxes and tap **Sign**

<div><figure><img src="../.gitbook/assets/Verify.png" alt="" width="188"><figcaption></figcaption></figure> <figure><img src="../.gitbook/assets/verify-swap.png" alt="" width="188"><figcaption></figcaption></figure></div>

5. A QR code will appear

{% hint style="info" %}
Select **Local mode** to sign without using the Vultisig relay server (requires same WiFi network).
{% endhint %}

<figure><img src="../.gitbook/assets/swap-keysign.png" alt="" width="188"><figcaption></figcaption></figure>

### On the Pair Device(s)

1. Select the same vault
2. Tap the **Camera** icon
3. Scan the QR code from the main device

<figure><img src="../.gitbook/assets/Main pair.png" alt="" width="188"><figcaption></figcaption></figure>

4. Verify transaction details match
5. Tap **Join key sign**

<div><figure><img src="../.gitbook/assets/Verify pair.png" alt="" width="188"><figcaption></figcaption></figure> <figure><img src="../.gitbook/assets/Swap-verify-pairing.png" alt="" width="188"><figcaption></figcaption></figure></div>

### Completion

Signing starts automatically when the threshold of devices has joined. The transaction signs and broadcasts from one of the devices.

<figure><img src="../.gitbook/assets/Succes-swap.png" alt="" width="188"><figcaption></figcaption></figure>

***

## Network Modes

| Mode | Description | Use Case |
|------|-------------|----------|
| **Internet** | Via Vultisig relay server | Devices on different networks |
| **Local** | Via local WiFi (mDNS) | Same network, maximum privacy |

***

## Related

- [TSS Actions](tss-actions.md) — Key generation, signing, resharing
- [Sending](../app-guide/wallet/sending.md) — How to send transactions
- [Swapping](../app-guide/wallet/swapping.md) — How to swap assets
