# How does a Vultisigner work?

## Set Up

### 2-of-2 Vault

In this setup, the Vault is created together with a user's device and the Vultisigner server. \
The user registers a keygen request with the Vultisigner server, which creates a new Vault that records the connection to the user's device. A keygen ceremony is then executed, which creates the Vault shares on the user's device and on the Vultisigner server. \
For security reasons, these Vault shares are password encrypted and the user must provide an email address to which the Vault share will be sent by the Vultisigner for backup and independent access.

### 2-of-3 Vault

In this setup, the Vault can be created initially together with the Vultisigner, or later added to an existing Vault with the[ Reshare](../user-actions/managing-your-vault/vault-reshare.md) feature. \
The rest of the process is similar to that described for setting up a [2-of-2 Vault](how-does-vultisigner-work.md#id-2-of-2-vault).

## Vault Share storage

The Vault Shares are stored on a dedicated Vultisigner server, which contains the Vault Shares along with configured instructions for signing transactions.\
Those Vault shares can be requested via email.

{% hint style="info" %}
Requesting the Vault Share stored on the Vultisigner server via email is always advised but optional in a 2 of 3 vault. \
**It is mandatory in a 2 of 2 setup.**
{% endhint %}

## Transaction Signing

When a user wishes to sign a transaction, the request is sent to the Vultisigner server. \
The appropriate Vultisigner Vault will verify that the configured transaction policies are met before participating in the Keysign process. \
If they are met, the Vultisigner joins the Keysign ceremony and the transaction is broadcasted to the blockchain.

<figure><img src="../.gitbook/assets/default vultisigner.png" alt=""><figcaption><p>Vultisigner Flowchart</p></figcaption></figure>
