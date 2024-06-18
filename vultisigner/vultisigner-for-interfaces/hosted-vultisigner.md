# Hosted Vultisigner

## How the hosted Vultisigner works

In the hosted version of Vultisigner, the integration partner runs their own version of Vultisigner on their infrastructure, giving them more control and customization possibilities. \
This configuration allows the partners to oversee the entire process from key generation to transaction signing while the user interacts with their interface.

***

<figure><img src="../../.gitbook/assets/Hosted (1).png" alt=""><figcaption><p>Flowchart - Hosted</p></figcaption></figure>

### Key generation and Vault setup

In the hosted version, the interface creates a Vault in collaboration with the user to provide the Vault shares. The interface's server initiates the process, creates the vault, and issues the configured vault shares. \
These shares are then securely distributed to the user's devices, ensuring that the user retains control of their Vault shares while the partner facilitates the initial setup.

***

### Share Management

{% hint style="info" %}
In the hosted version the Vault shares are managed at the discretion of the interfaces!
{% endhint %}

They have two primary options:

* **Client-side storage:** Vault shares can be stored in the user's browser, encrypted with the user's password, where the user can download the share and import it into their Vultisig application. Once the shares are securely stored on the client side, they are deleted from the server.\

* **Server-side storage:** Alternatively, Vault shares can be stored on the interface's Vultisig server, encrypted with the user's password. This method ensures secure storage and facilitates easy retrieval when needed.

***

### Keysign process

The interface server registers a session with the Vultisig relay server, encapsulating all the necessary details to coordinate the key signing process.

The Vultisigner instance on the interface's server receives the request to sign a transaction based on user input and initiates the keysign session.

### QR Code Generation

To sign a transaction, the front-end interface generates a QR code containing the transaction payload, session ID, and other pertinent information. The user then scans the QR code with his device.

The interface's Vultisig server monitors for the presence of enough devices of the user to meet the signing threshold.\
Once the required number of devices is detected, the server initiates the keysign ceremony.

#### Transaction broadcast

Upon successful signing, the server broadcasts the signed transaction to the blockchain. The interface then displays the transaction status and details within their application.

***

### Benefits of Hosted Vultisigner

* **Extended control:** Integration partners retain full control over the key generation and transaction signing processes, allowing them to tailor the system to their specific operational and security requirements.
* **Seamless integration:** The Hosted Vultisigner is designed to integrate seamlessly with the partner's existing infrastructure, ensuring a consistent and streamlined user experience.
* **Operational Efficiency:** By managing the entire Vultisigner process internally, partners reduce reliance on external services and increase operational efficiency. This self-sufficient approach enables faster response times and more direct control over the end-user experience.

By using the hosted version of Vultisigner, integration partners can offer their users a flexible and fully controlled environment for interacting with their services.&#x20;

This configuration combines the capabilities of Vultisigner and its security, with the partner's customizable and efficient infrastructure, providing an advanced solution for distributed application interactions.

\
