# Why is it safe?

Since the Vultisig server always holds only one Vault share, it can never sign a transaction on its own without the user's participation, even if the server is compromised. \
In addition, the server is configured to never initiate a transaction; it can only participate in the key signing process.

In the event that the Vultisig server is offline, the user still retains the threshold majority required to perform key generation and key signing ceremonies. \
In addition, the Vault share hosted on the Vultisig server can be requested by email to be sent to the user, which is mandatory in a 2-of-2 Vault configuration.
