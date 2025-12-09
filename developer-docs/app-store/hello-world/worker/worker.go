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

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := newConfig()
	if err != nil {
		logrus.Fatalf("failed to load config: %v", err)
	}

	logger := logrus.New()

	// The storage block refers to the S3 storage where the reshares required for signing transactions are stored.
	// Minio can be used here
	vaultStorage, err := vault.NewBlockStorageImp(cfg.BlockStorage)
	if err != nil {
		logger.Fatalf("failed to initialize Vault storage: %v", err)
	}

	redisConnOpt, err := asynq.ParseRedisURI(cfg.Redis.URI)
	if err != nil {
		logger.Fatalf("failed to parse redis URI: %v", err)
	}

	client := asynq.NewClient(redisConnOpt)
	consumer := asynq.NewServer(
		redisConnOpt,
		asynq.Config{
			Logger:      logger,
			Concurrency: 10,
			Queues: map[string]int{
				tasks.QUEUE_NAME: 10,
			},
		},
	)

	// Prepare txIndexStore to collect all signed txs and then control their status by tx_indexer if needed
	txIndexerStore, err := storage.NewPostgresTxIndexStore(ctx, cfg.Postgres.DSN)
	if err != nil {
		panic(fmt.Errorf("storage.NewPostgresTxIndexStore: %w", err))
	}

	chains, err := tx_indexer.Chains()
	if err != nil {
		panic(fmt.Errorf("tx_indexer.Chains: %w", err))
	}

	txIndexerService := tx_indexer.NewService(
		logger,
		txIndexerStore,
		chains,
	)

	// vaultService is necessary for async /reshare and /sign processes
	vaultService, err := vault.NewManagementService(
		cfg.VaultService,
		client,
		vaultStorage,
		txIndexerService,
	)
	if err != nil {
		panic(fmt.Errorf("failed to create vault service: %w", err))
	}

	rpcClient, err := ethclient.Dial(cfg.Rpc)
	if err != nil {
		panic(fmt.Errorf("failed to create eth client: %w", err))
	}

	triggerService := NewTrigger(
		logger,
		vaultStorage,
		cfg.VaultService.EncryptionSecret,
		rpcClient,
		keysign.NewSigner(
			logger.WithField("pkg", "keysign.Signer").Logger,
			relay.NewRelayClient(cfg.VaultService.Relay.Server),
			[]keysign.Emitter{
				keysign.NewVerifierEmitter(cfg.Verifier.URL, cfg.Verifier.Token),
				keysign.NewPluginEmitter(client, tasks.TypeKeySignDKLS, tasks.QUEUE_NAME),
			},
			[]string{
				cfg.Verifier.PartyPrefix,
				cfg.VaultService.LocalPartyPrefix,
			},
		),
	)

	go triggerService.Wait(ctx)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeKeySignDKLS, vaultService.HandleKeySignDKLS)
	mux.HandleFunc(tasks.TypeReshareDKLS, vaultService.HandleReshareDKLS)
	if err := consumer.Run(mux); err != nil {
		panic(fmt.Errorf("could not run server: %w", err))
	}
}

type workerCfg struct {
	VaultService vault_config.Config       `mapstructure:"vault_service"`
	BlockStorage vault_config.BlockStorage `mapstructure:"block_storage"`
	Postgres     config.Database
	Redis        config.Redis
	Verifier     config.Verifier
	Rpc          string
}

func newConfig() (workerCfg, error) {
	var cfg workerCfg
	err := envconfig.Process("", &cfg)
	if err != nil {
		return workerCfg{}, fmt.Errorf("failed to process env var: %w", err)
	}
	return cfg, nil
}
