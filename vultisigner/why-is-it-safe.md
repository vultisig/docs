# Why is it safe?

Since the Vultisig server will always have only one Vault share, it can never sign a transaction on its own without the user, even if the Vultisig server is compromised. Furthermore, the server is configured to never initiate a transaction and can only sign one.

In the event that the Vultisig server is offline, the user will still have the threshold majority to perform keygen and keysign ceremonies.
