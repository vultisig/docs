---
description: >-
  Vultiserver security model: Single share can't sign alone, server can't
  initiate transactions, secure even if compromised. Backup procedures.
---

# Why is it safe?

Since the Vultiserver always holds only one Vault Share, it will never be able to sign a transaction on its own without the participation and consent of the user, even in the event that the server is compromised.\
\
In addition, the server is configured so that it can never initiate a transaction; it can only be a participant in the key signing process.\\

In the event that the Vultisig server is offline, the user will still need the threshold majority required to perform Keysign ceremonies.

The Vultisigner's password-encrypted Vault Share which was emailed to the user during setup must be stored securely by the user (at a separate location from the user's own device's Vault Share) for this reason.
