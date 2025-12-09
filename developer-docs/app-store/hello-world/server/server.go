package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"

	"github.com/vultisig/verifier/plugin"
	"github.com/vultisig/verifier/plugin/config"
	smetrics "github.com/vultisig/verifier/plugin/metrics"
	"github.com/vultisig/verifier/plugin/policy"
	"github.com/vultisig/verifier/plugin/policy/policy_pg"
	"github.com/vultisig/verifier/plugin/redis"
	"github.com/vultisig/verifier/plugin/scheduler"
	"github.com/vultisig/verifier/plugin/server"
	"github.com/vultisig/verifier/vault"
	"github.com/vultisig/verifier/vault_config"
)

// main is the entry point for the Vultisig Plugin Server.
// It initializes all necessary infrastructure components (database, cache, storage),
// sets up the policy engine, and starts the HTTP server to handle requests.
func main() {
	// Create a root context that can be cancelled to gracefully shut down all services.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load configuration from environment variables.
	// This defines connection strings for Postgres, Redis, and S3 storage.
	cfg, err := newConfig()
	if err != nil {
		logrus.Fatalf("failed to load config: %v", err)
	}

	logger := logrus.New()

	// Initialize Redis client.
	// Redis is used for two primary purposes:
	// 1. Caching frequently accessed data to improve performance.
	// 2. Distributed locking and coordination between plugin instances.
	redisClient, err := redis.NewRedis(config.Redis{
		URI: cfg.Redis.URI,
	})
	if err != nil {
		logger.Fatalf("failed to initialize Redis client: %v", err)
	}

	// Parse the same Redis URI for the 'asynq' library.
	asynqConnOpt, err := asynq.ParseRedisURI(cfg.Redis.URI)
	if err != nil {
		logger.Fatalf("failed to parse redis URI: %v", err)
	}

	// Initialize Asynq client and inspector.
	// Asynq is a task queue used to handle background jobs asynchronously.
	// - Client: Enqueues tasks (e.g., "sign transaction", "reshare key").
	// - Inspector: Allows querying the state of queues and tasks (used for monitoring/UI).
	asynqClient := asynq.NewClient(asynqConnOpt)
	asynqInspector := asynq.NewInspector(asynqConnOpt)

	// Initialize Block Storage (Vault).
	// This connects to an S3-compatible object storage service (like Minio or AWS S3).
	// It securely stores the TSS (Threshold Signature Scheme) key shares required for
	// signing transactions. These files are encrypted and critical for operation.
	vaultStorage, err := vault.NewBlockStorageImp(cfg.BlockStorage)
	if err != nil {
		logger.Fatalf("failed to initialize Vault storage: %v", err)
	}

	// Initialize PostgreSQL Connection Pool.
	// Postgres is the primary persistent store for the application.
	// It holds:
	// - Policy definitions (automations created by users).
	// - Transaction history and logs.
	// - Plugin configuration state.
	pgPool, err := pgxpool.New(ctx, cfg.Postgres.DSN)
	if err != nil {
		logger.Fatalf("failed to initialize Postgres pool: %v", err)
	}

	// Initialize Policy Storage with Auto-Migrations.
	// plugin.WithMigrations ensures the database schema is up-to-date by applying
	// migration files found in "policy/policy_pg/migrations" before the app starts.
	// It returns a repository interface for interacting with policy data.
	policyStorage, err := plugin.WithMigrations(
		logger,
		pgPool,
		policy_pg.NewRepo,
		"policy/policy_pg/migrations",
	)
	if err != nil {
		logger.Fatalf("failed to initialize policy storage: %v", err)
	}

	// Initialize the Policy Service.
	// This service acts as the brain of the plugin, managing the lifecycle of automations.
	// It validates user requests against defined rules before any action is taken.
	// Note: scheduler.NewNilService() is a placeholder. If your plugin requires
	// time-based execution (e.g., "execute every hour"), a real scheduler implementation goes here.
	policyService, err := policy.NewPolicyService(
		policyStorage,
		scheduler.NewNilService(),
		logger,
	)
	if err != nil {
		logger.Fatalf("failed to initialize policy service: %v", err)
	}

	// Create the HTTP Server instance.
	// This bundles all the services (Policy, Redis, Storage, Asynq) into a runnable server.
	// It also registers the 'spec' (defined in newSpec()), which tells the Vultisig
	// ecosystem what capabilities this plugin has (e.g., supported chains, API recipes).
	srv := server.NewServer(
		cfg.Server,
		policyService,
		redisClient,
		vaultStorage,
		asynqClient,
		asynqInspector,
		newSpec(), // See spec.go for the plugin definition
		server.DefaultMiddlewares(),
		smetrics.NewNilPluginServerMetrics(),
		logger,
	)

	// Handle Graceful Shutdown.
	// Listen for OS signals (SIGINT/SIGTERM) to allow the server to finish
	// ongoing requests before shutting down, preventing data corruption.
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		logger.Infof("received shutdown signal, shutting down gracefully")
		cancel()
	}()

	// Start the server and block until it stops or an error occurs.
	if err := srv.Start(ctx); err != nil {
		logger.Fatalf("failed to start server: %v", err)
	}
}

// serverCfg holds the runtime configuration structure.
// Values are populated from environment variables (e.g., POSTGRES_DSN -> Postgres.DSN).
type serverCfg struct {
	Server       server.Config             // HTTP server settings (port, host)
	BlockStorage vault_config.BlockStorage `mapstructure:"block_storage"` // S3 bucket credentials
	Postgres     config.Database           // Database connection string
	Redis        config.Redis              // Redis connection URI
}

// newConfig parses environment variables into the serverCfg struct.
// It uses the envconfig library to map ENV_VARS to struct fields.
func newConfig() (serverCfg, error) {
	var cfg serverCfg
	err := envconfig.Process("", &cfg)
	if err != nil {
		return serverCfg{}, fmt.Errorf("failed to process env var: %w", err)
	}
	return cfg, nil
}
