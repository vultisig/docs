# Vultiserver

Vultiserver is an infrastructure feature of Vultisig that enables automatic co-signing of transactions within a user's Vault, but only under [pre-defined parameters](what-can-be-configured.md) set by the Vault owner. It acts as an automatic co-signer that can also engage when only certain user-defined criteria are met.

This functionality allows users to manage a Vultisig "Fast Vault" that inherently functions as a multi-device, multi-factor vault, while providing the user experience of a single-signature wallet when signing transactions. \
This setup is called a 2 of 2 device Fast Vault configuration.

Therefore, the Vultiserver feature ensures a smooth "one-device" experience, which (if desired) undergoes automatic security checks on sending that are consistently met once configured. These features can be modified by obtaining the threshold majority of Vault shares and signing a transaction, providing flexibility in maintaining and updating security protocols.
