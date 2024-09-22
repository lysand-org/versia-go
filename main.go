package main

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"database/sql"
	"database/sql/driver"
	"github.com/versia-pub/versia-go/pkg/versia"
	"net/http"
	"os"
	"os/signal"
	"slices"
	"strings"
	"sync"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"git.devminer.xyz/devminer/unitel"
	"git.devminer.xyz/devminer/unitel/unitelhttp"
	"git.devminer.xyz/devminer/unitel/unitelsql"
	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	pgx "github.com/jackc/pgx/v5/stdlib"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/versia-pub/versia-go/ent"
	"github.com/versia-pub/versia-go/ent/instancemetadata"
	"github.com/versia-pub/versia-go/internal/config"
	"github.com/versia-pub/versia-go/internal/database"
	"github.com/versia-pub/versia-go/internal/repository"
	"github.com/versia-pub/versia-go/internal/repository/repo_impls"
	"github.com/versia-pub/versia-go/internal/service/svc_impls"
	"github.com/versia-pub/versia-go/internal/task"
	"github.com/versia-pub/versia-go/internal/task/task_impls"
	"github.com/versia-pub/versia-go/internal/utils"
	"github.com/versia-pub/versia-go/internal/validators/val_impls"
	"github.com/versia-pub/versia-go/pkg/taskqueue"
	"modernc.org/sqlite"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	rootCtx, cancelRoot := context.WithCancel(context.Background())

	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	zerologr.NameFieldName = "logger"
	zerologr.NameSeparator = "/"
	zerologr.SetMaxV(2)

	config.Load()

	tel, err := unitel.Initialize(config.C.Telemetry)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize telemetry")
	}

	httpClient := &http.Client{
		Transport: unitelhttp.NewTracedTransport(
			tel,
			&http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: config.C.Telemetry.Environment == "development"}},
			[]string{"origin", "x-nonce", "x-signature", "x-signed-by", "sentry-trace", "sentry-baggage"},
			[]string{"host", "x-nonce", "x-signature", "x-signed-by", "sentry-trace", "sentry-baggage"},
			unitelhttp.WithLogger(zerologr.New(&log.Logger).WithName("http-client")),
			unitelhttp.WithTracePropagation(shouldPropagate),
		),
	}

	log.Debug().Msg("Opening database connection")
	var db *ent.Client
	if strings.HasPrefix(config.C.DatabaseURI, "postgres://") {
		db, err = openDB(tel, false, config.C.DatabaseURI)
	} else {
		db, err = openDB(tel, true, config.C.DatabaseURI)
	}
	if err != nil {
		log.Fatal().Err(err).Msg("Failed opening connection to the database")
	}
	defer db.Close()

	nc, err := nats.Connect(config.C.NATSURI)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to NATS")
	}

	log.Debug().Msg("Starting taskqueue client")
	tq, err := taskqueue.NewClient(config.C.NATSStreamName, nc, tel, zerologr.New(&log.Logger).WithName("taskqueue-client"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create taskqueue client")
	}

	log.Debug().Msg("Running schema migration")
	if err := migrateDB(db, zerologr.New(&log.Logger).WithName("migrate-db"), tel); err != nil {
		log.Fatal().Err(err).Msg("Failed to run schema migration")
	}

	log.Debug().Msg("Initializing instance")
	if err := initInstance(db, tel); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize instance")
	}

	// Stateless services

	requestSigner := svc_impls.NewRequestSignerImpl(tel, zerologr.New(&log.Logger).WithName("request-signer"))
	federationService := svc_impls.NewFederationServiceImpl(httpClient, tel, zerologr.New(&log.Logger).WithName("federation-service"))

	// Repositories

	repos := repo_impls.NewManagerImpl(
		db, tel, zerologr.New(&log.Logger).WithName("repositories"),
		func(db *ent.Client, log logr.Logger, telemetry *unitel.Telemetry) repository.UserRepository {
			return repo_impls.NewUserRepositoryImpl(federationService, db, log, telemetry)
		},
		repo_impls.NewNoteRepositoryImpl,
		repo_impls.NewFollowRepositoryImpl,
		func(db *ent.Client, log logr.Logger, telemetry *unitel.Telemetry) repository.InstanceMetadataRepository {
			return repo_impls.NewInstanceMetadataRepositoryImpl(federationService, db, log, telemetry)
		},
	)

	// Validators

	bodyValidator := val_impls.NewBodyValidator(zerologr.New(&log.Logger).WithName("validation-service"))
	requestValidator := val_impls.NewRequestValidator(repos, tel, zerologr.New(&log.Logger).WithName("request-validator"))

	// Task handlers

	notes := task_impls.NewNoteHandler(federationService, repos, tel, zerologr.New(&log.Logger).WithName("task-note-handler"))
	notesSet := registerTaskHandler(rootCtx, "notes", tq, notes)

	taskManager := task_impls.NewManager(notes, tel, zerologr.New(&log.Logger).WithName("task-manager"))

	// Services

	taskService := svc_impls.NewTaskServiceImpl(taskManager, tel, zerologr.New(&log.Logger).WithName("task-service"))
	userService := svc_impls.NewUserServiceImpl(repos, federationService, tel, zerologr.New(&log.Logger).WithName("user-service"))
	noteService := svc_impls.NewNoteServiceImpl(federationService, taskService, repos, tel, zerologr.New(&log.Logger).WithName("note-service"))
	followService := svc_impls.NewFollowServiceImpl(federationService, repos, tel, zerologr.New(&log.Logger).WithName("follow-service"))
	inboxService := svc_impls.NewInboxService(federationService, repos, tel, zerologr.New(&log.Logger).WithName("inbox-service"))
	instanceMetadataService := svc_impls.NewInstanceMetadataServiceImpl(federationService, repos, tel, zerologr.New(&log.Logger).WithName("instance-metadata-service"))

	wg := sync.WaitGroup{}

	if config.C.Mode == config.ModeWeb || config.C.Mode == config.ModeCombined {
		wg.Add(1)
		go func() {
			defer wg.Done()

			if err := server(
				rootCtx,
				tel,
				db,
				nc,
				federationService,
				requestSigner,
				bodyValidator,
				requestValidator,
				userService,
				noteService,
				followService,
				instanceMetadataService,
				inboxService,
			); err != nil {
				log.Fatal().Err(err).Msg("Failed to start server")
			}
		}()
	}

	maybeRunTaskHandler(rootCtx, "notes", notesSet, &wg)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	<-signalCh

	log.Info().Msg("Shutting down")

	cancelRoot()

	wg.Wait()
}

func openDB(tel *unitel.Telemetry, isSqlite bool, uri string) (*ent.Client, error) {
	s := tel.StartSpan(context.Background(), "function", "main.openDB")
	defer s.End()

	var drv driver.Driver
	var dialectType string
	var dbType string

	if isSqlite {
		log.Debug().Msg("Opening SQLite database connection")
		drv = &sqliteDriver{Driver: &sqlite.Driver{}}
		dialectType = dialect.SQLite
		dbType = "sqlite"
	} else {
		log.Debug().Msg("Opening PostgreSQL database connection")
		drv = &pgx.Driver{}
		dialectType = dialect.Postgres
		dbType = "postgres"
	}

	sql.Register(dialectType+"-traced", unitelsql.NewTracedSQL(tel, drv, dbType))

	db, err := sql.Open(dialectType+"-traced", uri)
	if err != nil {
		return nil, err
	}

	entDrv := entsql.OpenDB(dialectType, db)
	return ent.NewClient(ent.Driver(entDrv)), nil
}

func migrateDB(db *ent.Client, log logr.Logger, telemetry *unitel.Telemetry) error {
	s := telemetry.StartSpan(context.Background(), "function", "main.migrateDB")
	defer s.End()
	ctx := s.Context()

	log.V(1).Info("Migrating database schema")
	if err := db.Schema.Create(ctx); err != nil {
		log.Error(err, "Failed to migrate database schema")
		return err
	}

	log.V(1).Info("Database migration complete")

	return nil
}

func initInstance(db *ent.Client, telemetry *unitel.Telemetry) error {
	s := telemetry.StartSpan(context.Background(), "function", "main.initInstance")
	defer s.End()
	ctx := s.Context()

	tx, err := database.BeginTx(ctx, db, telemetry)
	if err != nil {
		return err
	}
	defer func(tx *database.Tx) {
		if err := tx.Finish(); err != nil {
			log.Error().Err(err).Msg("Failed to finish transaction")
		}
	}(tx)
	ctx = tx.Context()

	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate keypair")
		return err
	}

	err = tx.InstanceMetadata.Create().
		SetIsRemote(false).
		SetURI(utils.InstanceMetadataAPIURL().String()).
		SetExtensions(versia.Extensions{}).
		SetName(config.C.InstanceName).
		SetNillableDescription(config.C.InstanceDescription).
		SetHost(config.C.Host).
		SetPrivateKey(priv).
		SetPublicKey(pub).
		SetPublicKeyAlgorithm("ed25519").
		SetSoftwareName("versia-go").
		SetSoftwareVersion("0.0.1").
		SetSharedInboxURI(utils.SharedInboxAPIURL().String()).
		SetAdminsURI(utils.InstanceMetadataAdminsAPIURL().String()).
		SetModeratorsURI(utils.InstanceMetadataModeratorsAPIURL().String()).
		SetSupportedVersions([]string{"0.4.0"}).
		SetSupportedExtensions([]string{}).
		//
		OnConflictColumns(instancemetadata.FieldHost).
		UpdateName().
		UpdateDescription().
		UpdateHost().
		UpdateSoftwareName().
		UpdateSoftwareVersion().
		UpdateSharedInboxURI().
		UpdateAdminsURI().
		UpdateModeratorsURI().
		UpdateSupportedVersions().
		UpdateSupportedExtensions().
		Exec(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create server metadata")
		return err
	}

	tx.MarkForCommit()

	return tx.Finish()
}

func registerTaskHandler[T task.Handler](ctx context.Context, name string, tq *taskqueue.Client, handler T) *taskqueue.Set {
	s, err := tq.Set(ctx, name)
	if err != nil {
		log.Fatal().Err(err).Str("handler", name).Msg("Could not create taskset for task handler")
	}

	handler.Register(s)

	return s
}

func maybeRunTaskHandler(ctx context.Context, name string, set *taskqueue.Set, wg *sync.WaitGroup) {
	l := log.With().Str("handler", name).Logger()

	if config.C.Mode == config.ModeWeb {
		l.Warn().Strs("requested", config.C.Consumers).Msg("Not starting task handler, as this process is running in web mode")
		return
	}

	if config.C.Mode == config.ModeConsumer && !slices.Contains(config.C.Consumers, name) {
		l.Warn().Strs("requested", config.C.Consumers).Msg("Not starting task handler, as it wasn't requested")
		return
	}

	wg.Add(1)

	c := set.Consumer(name)
	if err := c.Start(ctx); err != nil {
		l.Fatal().Err(err).Msg("Could not start task handler")
	}

	l.Info().Msg("Started task handler")

	go func() {
		defer wg.Done()

		<-ctx.Done()
		l.Debug().Msg("Got signal to stop task handler")

		c.Close()

		l.Info().Msg("Stopped task handler")
	}()
}
