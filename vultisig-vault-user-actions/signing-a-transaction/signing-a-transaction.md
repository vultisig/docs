---
description: How to sign a transaction with your Vault
---

# Signing a Transaction

## Fast Vault

#### Main Device

Follow the Sending or Signing flow, prepare the transaction with input all the needed values and press **Continue** or **Swap**

<div><figure><img src="../../.gitbook/assets/Send.png" alt="" width="188"><figcaption><p>Send screen</p></figcaption></figure> <figure><img src="../../.gitbook/assets/Swap-screen.png" alt="" width="188"><figcaption><p>Swap screen</p></figcaption></figure></div>

#### Verify your input



<div><figure><img src="../../.gitbook/assets/Verify pair.png" alt="" width="188"><figcaption><p>Send</p></figcaption></figure> <figure><img src="../../.gitbook/assets/verify-swap.png" alt="" width="188"><figcaption><p>Swap</p></figcaption></figure></div>

Put in your password and wait until the signing is finished

<figure><img src="../../.gitbook/assets/Succes-swap (1).png" alt="" width="188"><figcaption></figcaption></figure>

## Secure Vault

### Get your device's ready

Depending on your Vault type (2-of-2, 2-of-3, 3-of-4, or m-of-n), you need to have the threshold number of devices ready for the signing process.

{% hint style="info" %}
Remember: One device is creating the transaction, and the other devices are your "verification" devices. You need to "pair" the devices to each other to exchange the transaction information.

Make sure you confirm the transaction on your pair device.

Both devices will co-sign and send the final transaction. Only one transaction will be sent (you won't have duplicate transactions).
{% endhint %}

[![](../../.gitbook/assets/TwitterVideoThumbnail.jpeg)](https://twitter.com/iceman00008/status/1824690627843592573/video/1)

_Click on the above image to watch an explanation video on Twitter_

### Main Device

On your Main Device, prepare the transaction (e.g. Send or Swap). When ready, click on **Continue** or **Swap**.

<div><figure><img src="../../.gitbook/assets/Send (1).png" alt="" width="188"><figcaption><p>Send Screen</p></figcaption></figure> <figure><img src="../../.gitbook/assets/Swap-screen.png" alt="" width="188"><figcaption><p>Swap screen</p></figcaption></figure></div>

Confirm all the details on the Verify screen, check all boxes for the reminder messages, and click **Sign**.

<div><figure><img src="../../.gitbook/assets/Verify.png" alt="" width="188"><figcaption></figcaption></figure> <figure><img src="../../.gitbook/assets/verify-swap.png" alt="" width="188"><figcaption></figcaption></figure></div>

On the initiating device a QR code will be generated.

{% hint style="info" %}
If you want to sign locally without using the Vultisig Relay server, select &#x4C;_&#x6F;cal mode_
{% endhint %}

<figure><img src="../../.gitbook/assets/swap-keysign.png" alt="" width="188"><figcaption></figcaption></figure>

### Pair Device

On your other device, after selecting the same Vault, click the **Camera** icon and scan the QR code shown on the first device.

<figure><img src="../../.gitbook/assets/Main pair.png" alt="" width="188"><figcaption></figcaption></figure>

The transaction details will load, verify them and click **Join key sign** to proceed.

<div><figure><img src="../../.gitbook/assets/Verify pair (1).png" alt="" width="188"><figcaption></figcaption></figure> <figure><img src="../../.gitbook/assets/Swap-verify-pairing.png" alt="" width="188"><figcaption></figcaption></figure></div>

The signing will automatically start when the Threshold of devices joined.

That's it! The transaction should sign and be sent by one of the devices.

<figure><img src="../../.gitbook/assets/Succes-swap.png" alt="" width="188"><figcaption></figcaption></figure>
