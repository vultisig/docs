# Why is it safe?

Since the Vultisiginer server always holds only one Vault share, it will never be able to sign a transaction on its own without the participation and consent of the user, even in the event that the server is compromised.\
In addition, the server is configured so that it can never initiate a transaction; it can only be a participant in the key signing process.

In the event that the Vultisig server is offline, the user will still have the threshold majority required to perform Keysign ceremonies.\
The hosted Vault share on the Vultisigner server can be requested by email to be sent to the user as indicated [before](how-does-vultisigner-work.md#vault-share-storage), which is mandatory in a 2-of-2 Vault configuration.
