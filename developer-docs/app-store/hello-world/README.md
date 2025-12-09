# Integrating with the Vultisig Ecosystem: A Developer's Guide

This document provides a comprehensive guide for developers on how to integrate their applications into the Vultisig ecosystem. By following these instructions, you can build a plugin that leverages Vultisig's decentralized security infrastructure for transaction signing and management. The guide covers mandatory requirements, optional features, project setup, and implementation details based on a sample application.

## Core Integration Requirements

To ensure seamless integration with the Vultisig ecosystem, your application must meet several architectural requirements.

### Mandatory Components

- **/reshare and /sign Endpoints**: Your application must expose these API endpoints to handle core TSS (Threshold Signature Scheme) operations for key resharing and transaction signing.
- **Automation API Support**: It must adopt the Vultisig API structure for managing and executing automations (policies).


### Optional Components

- **Transaction Indexing**: You can optionally implement a service to index and track the status of transactions your application submits to the blockchain.
- **Scheduler**: A scheduler can be included to trigger time-based or event-driven actions, such as recurring transactions.


## Getting Started

To accelerate development, you can import the `Server` component from the `github.com/vultisig/verifier/plugin` package. This provides a pre-built foundation that includes all the necessary API structures. The main entry point of the service, `server/main.go`, is responsible for initializing and connecting all the essential components.

## Project Structure

The example code is organized into the following key packages:

- `/server`: The primary HTTP server that exposes the plugin's API endpoints.
- `/spec`: Defines the plugin's capabilities, supported blockchains, and policy validation rules.
- `/worker`: Manages asynchronous background jobs, such as transaction signing and resharing, using a task queue.
- `/trigger`: An example service for initiating transactions via an external HTTP request.


## Server Initialization (`server/main.go`)

The `main` function in `server/main.go` orchestrates the setup of your service by initializing and wiring together several key components:

- **Configuration**: Loads application settings from environment variables into a `struct` using the `envconfig` library.
- **Redis**: Establishes a connection to a Redis instance, which serves as both a cache and a message broker for `asynq`.
- **Asynq**: Initializes the `asynq` client and inspector, which are used to enqueue and manage asynchronous background tasks.
- **Block Storage (Vault)**: Connects to an S3-compatible object storage service (like Minio). This storage is used to securely store vault data containing the reshared keys required for signing.
- **PostgreSQL**: Connects to a PostgreSQL database used for storing automation policies and transaction data. Database migrations are applied automatically on startup.
- **Policy Service**: Initializes the service responsible for creating, validating, and managing automation policies.
- **HTTP Server**: Creates and starts the main web server (`server.NewServer`), which bundles all the services and middleware.


## Defining the Plugin Specification (`spec/`)

This part of the application defines your plugin's capabilities and how it is presented within the Vultisig user interface.

- **Constants**: You must define a `PluginID` (a unique ID provided by the Vultisig team), a `PluginName`, and a list of `supportedChains`.
- **GetRecipeSpecification**: This function defines the "recipe" for an automation, specifying the configurable properties that users will see in the UI (e.g., `asset`, `toAddress`, `amount`) using a JSON schema. You should also provide a `ConfigurationExample` to guide users.
- **ValidatePluginPolicy**: This function is a hook for validating the data that a user enters in the UI against your defined schema.
- **Suggest**: Based on a user's initial input, this function can suggest a complete policy configuration, including `Rules` (e.g., `ethereum.send`) and security parameters like rate limits (`MaxTxsPerWindow`).
- **buildSupportedResources**: This function generates a list of all resources and actions supported by the plugin, such as signing a transaction (`chain.send`), and the parameters required for each action.


## Asynchronous Worker (`worker/`)

The worker is a background process that handles time-consuming or asynchronous tasks, ensuring the main server remains responsive.

- **Initialization**: The worker connects to Redis to listen for jobs on the `asynq` queue, initializes a `TxIndexer` for monitoring transaction statuses, and sets up the `VaultService` for handling `/reshare` and `/sign` operations.
- **Task Handling**: It subscribes to task types like `tasks.TypeKeySignDKLS` and `tasks.TypeReshareDKLS`. When a task is received, the corresponding handler in `VaultService` is executed to perform the cryptographic operations.
- **Trigger Service**: The worker also initializes and runs the `trigger` service, which exposes an HTTP server to listen for external requests to create and sign transactions.


## Triggering Transactions (`trigger/`)

This module provides an example of how to programmatically initiate a signing request. It runs an HTTP server on port `:8090` with a `/trigger` endpoint.

### Workflow

1. The `/trigger` endpoint receives a request with the policy `uuid`, `pubkey`, and transaction details like `toAddress` and `amount`.
2. It retrieves the encrypted vault from block storage using the policy information.
3. It constructs an unsigned transaction. In the example, it creates a native token transfer for Ethereum using `recipes/sdk/evm`.
4. It wraps the policy and transaction data into a `PluginKeysignRequest`.
5. It calls `signer.Sign()`, which dispatches the signing job to the Vultisig network and the local worker queue.
6. Once the transaction is signed, it broadcasts the completed transaction to the Ethereum network via an RPC endpoint.

## Build and Troubleshooting

### Dependency Management

To ensure compatibility, you may need to add the following `replace` directives to your `go.mod` file:

```go
replace (
    github.com/agl/ed25519 => github.com/binance-chain/edwards25519 v0.0.0-20200305024217-f36fc4b53d43
    github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
    nhooyr.io/websocket => github.com/coder/websocket v1.8.6
)
```


### macOS Development

- If you encounter a `dyld[...] Library not loaded` error when running the application locally, you must set the `DYLD_LIBRARY_PATH` environment variable to the path of the `libgodkls.dylib` library. You can find the latest libraries at the [vultisig/go-wrappers](https://github.com/vultisig/go-wrappers) GitHub repository.
- **Example**: `export DYLD_LIBRARY_PATH=/path/to/your/libs`


### Docker Build Performance

- Building the application in a Docker container can be very slow on ARM-based Macs (Apple Silicon) because of the compilation process for the underlying `dkls` cryptography library.
- For a more efficient development cycle, it is recommended to run the plugin service directly on your host machine and use Docker only for the other infrastructure components (e.g., PostgreSQL, Redis).

