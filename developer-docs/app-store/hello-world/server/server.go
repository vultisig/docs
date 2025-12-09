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

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := newConfig()
	if err != nil {
		logrus.Fatalf("failed to load config: %v", err)
	}

	logger := logrus.New()

	// We use Redis as a cache and for managing async tasks
	redisClient, err := redis.NewRedis(config.Redis{
		URI: cfg.Redis.URI,
	})
	if err != nil {
		logger.Fatalf("failed to initialize Redis client: %v", err)
	}

	asynqConnOpt, err := asynq.ParseRedisURI(cfg.Redis.URI)
	if err != nil {
		logger.Fatalf("failed to parse redis URI: %v", err)
	}

	asynqClient := asynq.NewClient(asynqConnOpt)
	asynqInspector := asynq.NewInspector(asynqConnOpt)

	// The storage block refers to the S3 storage where the reshares required for signing transactions are stored.
	// Minio can be used here
	vaultStorage, err := vault.NewBlockStorageImp(cfg.BlockStorage)
	if err != nil {
		logger.Fatalf("failed to initialize Vault storage: %v", err)
	}

	// We use Postgres as our SQL database to store information about policies, transactions, and other data necessary
	// for the application to function
	pgPool, err := pgxpool.New(ctx, cfg.Postgres.DSN)
	if err != nil {
		logger.Fatalf("failed to initialize Postgres pool: %v", err)
	}

	policyStorage, err := plugin.WithMigrations(
		logger,
		pgPool,
		policy_pg.NewRepo,
		// all necessary migrations will be applied
		"policy/policy_pg/migrations",
	)
	if err != nil {
		logger.Fatalf("failed to initialize policy storage: %v", err)
	}

	policyService, err := policy.NewPolicyService(
		policyStorage,
		// Can be implemented for your tasks, while we will use nil implementation
		scheduler.NewNilService(),
		logger,
	)
	if err != nil {
		logger.Fatalf("failed to initialize policy service: %v", err)
	}

	srv := server.NewServer(
		cfg.Server,
		policyService,
		redisClient,
		vaultStorage,
		asynqClient,
		asynqInspector,
		newSpec(),
		server.DefaultMiddlewares(),
		smetrics.NewNilPluginServerMetrics(),
		logger,
	)

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		logger.Infof("received shutdown signal, shutting down gracefully")
		cancel()
	}()

	if err := srv.Start(ctx); err != nil {
		logger.Fatalf("failed to start server: %v", err)
	}
}

type serverCfg struct {
	Server       server.Config
	BlockStorage vault_config.BlockStorage `mapstructure:"block_storage"`
	Postgres     config.Database
	Redis        config.Redis
}

func newConfig() (serverCfg, error) {
	var cfg serverCfg
	err := envconfig.Process("", &cfg)
	if err != nil {
		return serverCfg{}, fmt.Errorf("failed to process env var: %w", err)
	}
	return cfg, nil
}
