# How Keysigning works

<figure><img src="../../.gitbook/assets/How keysign works.png" alt=""><figcaption><p>Flowchart - TX signing</p></figcaption></figure>

A device initiates a cryptographic session to sign a transaction. It takes the user input and transforms it into the transaction payload, while also acting as the host for that session. The initiating device sends the session metadata, including the session ID, to the Vultisig relay server or broadcasts it over the local network. At the same time, a QR code is generated that embeds the session-specific details required for the pairing devices to join the session.

The pairing devices scan the QR code to join the session using the embedded session ID, hex chain code for pairing, which happens over encrypted messaging. The initiating device monitors the joining devices and initiates the keysigning ceremony when it detects the required number of devices.

During the keysigning ceremony, the participating devices jointly sign the transaction, ensuring compliance with the threshold signature scheme. Upon successful completion of the keysigning process, the initiating device propagates the signed transaction to the blockchain and distributes it to the other participating devices. These devices then display the transaction hash, verifying the successful execution and broadcast of the transaction.

This technical workflow provides a secure, coordinated process for session initiation, device pairing and multi-signature transaction execution.
