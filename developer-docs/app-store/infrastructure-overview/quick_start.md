# Vultisig App Development: Basics \& Quick Start

This guide provides a step-by-step walkthrough for building your first Vultisig app. By the end of this guide, you'll understand the core architecture, required components, and how to integrate your application with the Vultisig ecosystem.

## Overview

To integrate your application with the Vultisig ecosystem, you need to build an app that supports:

- **Mandatory**: `/reshare` and `/sign` endpoints for TSS operations
- **Mandatory**: API structure for automation management
- **Optional**: Transaction indexing for monitoring blockchain confirmations
- **Optional**: Scheduler for time-based or event-driven automations


## Getting Started with the Server

The fastest way to get started is to import the `Server` component from the Vultisig library. This provides all the necessary API infrastructure, including the required endpoints for signing, resharing, and automation management, so you don't need to implement them from scratch.

### Why use the built-in Server?

- **Handles TSS complexity**: All the cryptographic protocol complexity is abstracted away
- **Standardized endpoints**: Provides the exact API structure that the Vultisig Verifier expects
- **Built-in middleware**: Includes authentication, rate limiting, and error handling out of the box
- **Storage integration**: Automatically connects with Redis, PostgreSQL, and S3-compatible storage
- **Compatibility guarantee**: Ensures your app works seamlessly with the Vultisig mobile app and web interface


### Implementation

The server is your app's main entry point. It initializes all infrastructure components (database, cache, storage), sets up the automation engine, and starts the HTTP server to handle requests from the Vultisig ecosystem.

**Key responsibilities:**

- Set up S3-compatible storage for encrypted vault data (key shares)
- Apply database migrations automatically
- Register your plugin specification (capabilities and UI configuration)
- Create async task for worker(s)

**Full example**: [server/main.go on GitHub](https://github.com/vultisig/docs/blob/main/developer-docs/app-store/create-an-app/build-your-app/hello-world/server/server.go)

## Defining Your App Specification

The specification defines what your app does and how it appears in the Vultisig UI.

### What the Spec Does

- **UI Configuration**: Defines the form fields users see when creating an automation (e.g., asset selector, addresses, amounts)
- **Validation Rules**: Ensures user input meets your requirements before creating a policy
- **Permission Policies**: Translates high-level user intent into specific, security-constrained execution rules
- **Resource Declaration**: Tells the Verifier which blockchain operations your plugin can perform
- **Rate Limiting**: Sets transaction frequency limits to prevent abuse


### Key Components

- `GetRecipeSpecification()`: Returns JSON Schema for the UI form
- `ValidatePluginPolicy()`: Validates user input against your schema
- `Suggest()`: Converts user configuration into specific execution rules with constraints

**Full example**: [spec.go on GitHub](https://github.com/vultisig/docs/blob/main/developer-docs/app-store/create-an-app/build-your-app/hello-world/server/spec.go)

## Setting Up the Worker

The worker process handles background tasks asynchronously, separate from the HTTP server. This architecture ensures that long-running cryptographic operations don't block API requests.

### Worker Responsibilities

- **TSS Signing**: Participates in multi-party signing ceremonies to approve transactions
- **Key Resharing**: Creates distributed key shares for enhanced security
- **Transaction Monitoring**: Tracks submitted transactions and updates their status
- **Task Queue Processing**: Consumes jobs from Redis queue using Asynq


### Architecture Benefits

Running the worker separately from the server provides:

- **Scalability**: Can run multiple worker instances for high throughput
- **Isolation**: Cryptographic operations don't impact API responsiveness
- **Retry Logic**: Failed tasks can be automatically retried
- **Monitoring**: Easy to track task queue depth and processing metrics


### Key Components

- Initializes the Vault Management Service for TSS operations
- Sets up Transaction Indexer to monitor blockchain confirmations
- Registers handlers for `TypeKeySignDKLS` and `TypeReshareDKLS` tasks
- Processes tasks with configurable concurrency (default: 10 parallel tasks)

**Full example**: [worker/main.go on GitHub](https://github.com/vultisig/docs/blob/main/developer-docs/app-store/create-an-app/build-your-app/hello-world/worker/worker.go)

## Example: Transaction Trigger Implementation

The trigger service demonstrates **one possible implementation approach** for initiating transactions programmatically. This is not a required component, but serves as a reference for understanding the complete transaction lifecycle.

### What This Example Shows

- How to retrieve and decrypt vault data to derive blockchain addresses
- How to construct unsigned transactions using the EVM SDK
- How to create a `PluginKeysignRequest` and submit it to the TSS network
- How to broadcast signed transactions to the blockchain


### Important Note

This is **not a prescriptive pattern**—it's one of many ways to trigger transactions. Your plugin might:

- Listen to blockchain events instead of HTTP requests
- Use a scheduler for recurring transactions
- React to price feeds, governance votes, or other external data
- Implement custom business logic specific to your use case

The trigger example is provided to help you understand how the signing flow works end-to-end, not as a template you must follow.

**Full example**: [trigger/main.go on GitHub](https://github.com/vultisig/docs/blob/main/developer-docs/app-store/create-an-app/build-your-app/hello-world/worker/trigger.go)

## Next Steps

Now that you understand the basic structure, you can:

1. **Clone the example repository** and modify it for your use case
2. **Define your plugin specification** with the chains and operations you need
3. **Implement your business logic** as part of the worker or create separate service
4. **Test locally** with local verifier and App store UI