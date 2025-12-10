package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hibiken/asynq"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"

	"github.com/vultisig/verifier/plugin/config"
	"github.com/vultisig/verifier/plugin/keysign"
	"github.com/vultisig/verifier/plugin/tasks"
	"github.com/vultisig/verifier/plugin/tx_indexer"
	"github.com/vultisig/verifier/plugin/tx_indexer/pkg/storage"
	"github.com/vultisig/verifier/vault"
	"github.com/vultisig/verifier/vault_config"
	"github.com/vultisig/vultisig-go/relay"
)

// main initializes the worker process.
// Unlike the 'server' process which handles HTTP requests, the 'worker' focuses on
// executing long-running or resource-intensive background tasks, specifically:
// 1. TSS Key Resharing (refreshing key shares for security).
// 2. TSS Transaction Signing (participating in the multi-party signing ceremony).
// 3. Listening for external triggers.
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := newConfig()
	if err != nil {
		logrus.Fatalf("failed to load config: %v", err)
	}

	logger := logrus.New()

	// Initialize Block Storage (S3/Minio).
	// The worker needs access to the vault data to retrieve key shares during
	// signing and resharing operations.
	vaultStorage, err := vault.NewBlockStorageImp(cfg.BlockStorage)
	if err != nil {
		logger.Fatalf("failed to initialize Vault storage: %v", err)
	}

	// Connect to Redis for the Asynq task queue.
	redisConnOpt, err := asynq.ParseRedisURI(cfg.Redis.URI)
	if err != nil {
		logger.Fatalf("failed to parse redis URI: %v", err)
	}

	// Initialize the Asynq Client (to enqueue new tasks if needed by sub-services).
	client := asynq.NewClient(redisConnOpt)

	// Initialize the Asynq Server (Consumer).
	// This component listens to the Redis queue and pulls tasks to process.
	// Concurrency: 10 indicates it can handle 10 tasks in parallel.
	consumer := asynq.NewServer(
		redisConnOpt,
		asynq.Config{
			Logger:      logger,
			Concurrency: 10,
			Queues: map[string]int{
				tasks.QUEUE_NAME: 10, // Priority 10 for the default queue
			},
		},
	)

	// Initialize Transaction Indexer Storage.
	// This connects to Postgres to store the status of signed transactions.
	// It allows the plugin to track if a tx was successfully mined or failed.
	txIndexerStore, err := storage.NewPostgresTxIndexStore(ctx, cfg.Postgres.DSN)
	if err != nil {
		panic(fmt.Errorf("storage.NewPostgresTxIndexStore: %w", err))
	}

	// Load supported chains configuration for the indexer.
	chains, err := tx_indexer.Chains()
	if err != nil {
		panic(fmt.Errorf("tx_indexer.Chains: %w", err))
	}

	txIndexerService := tx_indexer.NewService(
		logger,
		txIndexerStore,
		chains,
	)

	// Initialize the Vault Management Service.
	// This service contains the core logic for:
	// - Processing /reshare requests (DKLS protocol).
	// - Processing /sign requests (DKLS protocol).
	// It interacts with the Vault Storage to read/write key shares.
	vaultService, err := vault.NewManagementService(
		cfg.VaultService,
		client, // Uses client to re-enqueue tasks if needed (e.g. retries)
		vaultStorage,
		txIndexerService,
	)
	if err != nil {
		panic(fmt.Errorf("failed to create vault service: %w", err))
	}

	// Dial the Ethereum RPC node.
	// Required by the Trigger service to construct and broadcast transactions.
	rpcClient, err := ethclient.Dial(cfg.Rpc)
	if err != nil {
		panic(fmt.Errorf("failed to create eth client: %w", err))
	}

	// Initialize the Trigger Service.
	// This exposes the HTTP endpoint (:8090/trigger) to manually start a transaction.
	// It uses keysign.NewSigner to communicate with the Vultisig Relay and Verifier.
	triggerService := NewTrigger(
		logger,
		vaultStorage,
		cfg.VaultService.EncryptionSecret,
		rpcClient,
		keysign.NewSigner(
			logger.WithField("pkg", "keysign.Signer").Logger,
			relay.NewRelayClient(cfg.VaultService.Relay.Server), // Connects to the p2p relay
			[]keysign.Emitter{
				// Emitter 1: Notifies the Verifier central server about the process
				keysign.NewVerifierEmitter(cfg.Verifier.URL, cfg.Verifier.Token),
				// Emitter 2: Enqueues a local task to track this signing operation
				keysign.NewPluginEmitter(client, tasks.TypeKeySignDKLS, tasks.QUEUE_NAME),
			},
			[]string{
				cfg.Verifier.PartyPrefix,          // e.g. "verifier"
				cfg.VaultService.LocalPartyPrefix, // e.g. "local-worker"
			},
		),
	)

	// Start the Trigger HTTP server in a background goroutine.
	go triggerService.Wait(ctx)

	// Route Task Types to their Handlers.
	// When a task of type 'TypeKeySignDKLS' comes in, execute 'vaultService.HandleKeySignDKLS'.
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeKeySignDKLS, vaultService.HandleKeySignDKLS)
	mux.HandleFunc(tasks.TypeReshareDKLS, vaultService.HandleReshareDKLS)

	// Start the Asynq worker server.
	// This blocks and processes tasks until the application is stopped.
	if err := consumer.Run(mux); err != nil {
		panic(fmt.Errorf("could not run server: %w", err))
	}
}

type workerCfg struct {
	VaultService vault_config.Config       `mapstructure:"vault_service"` // Settings for TSS parties
	BlockStorage vault_config.BlockStorage `mapstructure:"block_storage"` // S3 Credentials
	Postgres     config.Database           // DB connection for indexer
	Redis        config.Redis              // Redis connection for queue
	Verifier     config.Verifier           // Config for talking to Vultisig Verifier
	Rpc          string                    // RPC URL (e.g. Infura/Alchemy)
}

func newConfig() (workerCfg, error) {
	var cfg workerCfg
	err := envconfig.Process("", &cfg)
	if err != nil {
		return workerCfg{}, fmt.Errorf("failed to process env var: %w", err)
	}
	return cfg, nil
}
