# Hello World Plugin

Building blocks to clone and start creating your own plugin — whether a fixed workflow or an AI agent that triggers specific actions.

## Files

| File | Purpose |
|------|---------|
| [`server/server.go`](server/server.go) | Plugin server entry point — initializes infra and starts HTTP server |
| [`server/spec.go`](server/spec.go) | Plugin specification — recipe schema, validation, supported chains |
| [`worker/worker.go`](worker/worker.go) | Background worker — TSS signing, resharing, task queue processing |
| [`worker/tx.go`](worker/tx.go) | Transaction trigger example — builds and signs transactions |
| [`server.example.json`](server.example.json) | Server configuration template |
| [`worker.example.json`](worker.example.json) | Worker configuration template |

## Quick Start

1. Clone this directory as your plugin base
2. Update `spec.go` with your plugin ID, name, supported chains, and recipe schema
3. Implement your business logic in the worker or a separate trigger service
4. Configure `server.example.json` and `worker.example.json` with your infrastructure details

Each plugin is independent — modify anything to fit your use case.
