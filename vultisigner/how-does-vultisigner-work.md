# How does a Vultisigner work?

A vault share is stored on a dedicated Vultisig server, which holds the vault along with instructions on when that vault share should sign transactions.

When a user wishes to sign a transaction, the request is sent to the Vultisig server. The appropriate vault share will participate in the key signing process only if the specified requirements are met.
