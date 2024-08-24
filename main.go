package main

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"git.devminer.xyz/devminer/unitel/unitelhttp"
	"git.devminer.xyz/devminer/unitel/unitelsql"
	"github.com/lysand-org/versia-go/ent/instancemetadata"
	"github.com/lysand-org/versia-go/internal/handlers/follow_handler"
	"github.com/lysand-org/versia-go/internal/handlers/meta_handler"
	"github.com/lysand-org/versia-go/internal/handlers/note_handler"
	"github.com/lysand-org/versia-go/internal/repository"
	"github.com/lysand-org/versia-go/internal/repository/repo_impls"
	"github.com/lysand-org/versia-go/internal/service/svc_impls"
	"github.com/lysand-org/versia-go/internal/validators/val_impls"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	pgx "github.com/jackc/pgx/v5/stdlib"
	"github.com/lysand-org/versia-go/config"
	"github.com/lysand-org/versia-go/ent"
	"github.com/lysand-org/versia-go/internal/database"
	"github.com/lysand-org/versia-go/internal/handlers/user_handler"
	"github.com/lysand-org/versia-go/internal/tasks"
	"github.com/lysand-org/versia-go/internal/utils"
	"github.com/lysand-org/versia-go/pkg/taskqueue"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"modernc.org/sqlite"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func shouldPropagate(r *http.Request) bool {
	return config.C.ForwardTracesTo.Match([]byte(r.URL.String()))
}

func main() {
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
	tq, err := taskqueue.NewClient(context.Background(), config.C.NATSStreamName, nc, tel, zerologr.New(&log.Logger).WithName("taskqueue-client"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create taskqueue client")
	}
	defer tq.Close()

	log.Debug().Msg("Running schema migration")
	if err := migrateDB(db, zerologr.New(&log.Logger).WithName("migrate-db"), tel); err != nil {
		log.Fatal().Err(err).Msg("Failed to run schema migration")
	}

	// Stateless services

	requestSigner := svc_impls.NewRequestSignerImpl(tel, zerologr.New(&log.Logger).WithName("request-signer"))
	federationService := svc_impls.NewFederationServiceImpl(httpClient, tel, zerologr.New(&log.Logger).WithName("federation-service"))
	taskService := svc_impls.NewTaskServiceImpl(tq, tel, zerologr.New(&log.Logger).WithName("task-service"))

	// Manager

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

	// Services

	userService := svc_impls.NewUserServiceImpl(repos, federationService, tel, zerologr.New(&log.Logger).WithName("user-service"))
	noteService := svc_impls.NewNoteServiceImpl(federationService, taskService, repos, tel, zerologr.New(&log.Logger).WithName("note-service"))
	followService := svc_impls.NewFollowServiceImpl(federationService, repos, tel, zerologr.New(&log.Logger).WithName("follow-service"))
	inboxService := svc_impls.NewInboxService(federationService, repos, tel, zerologr.New(&log.Logger).WithName("inbox-service"))
	instanceMetadataService := svc_impls.NewInstanceMetadataServiceImpl(federationService, repos, tel, zerologr.New(&log.Logger).WithName("instance-metadata-service"))

	// Handlers

	userHandler := user_handler.New(federationService, requestSigner, userService, inboxService, bodyValidator, requestValidator, zerologr.New(&log.Logger).WithName("user-handler"))
	noteHandler := note_handler.New(noteService, bodyValidator, zerologr.New(&log.Logger).WithName("notes-handler"))
	followHandler := follow_handler.New(followService, federationService, zerologr.New(&log.Logger).WithName("follow-handler"))
	metaHandler := meta_handler.New(instanceMetadataService, zerologr.New(&log.Logger).WithName("meta-handler"))

	// Initialization

	if err := initServerActor(db, tel); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize server actor")
	}

	web := fiber.New(fiber.Config{
		ProxyHeader:           "X-Forwarded-For",
		ErrorHandler:          fiberErrorHandler,
		DisableStartupMessage: true,
		AppName:               "versia-go",
		EnablePrintRoutes:     true,
	})

	web.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return true
		},
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, b3, traceparent, sentry-trace, baggage",
		AllowCredentials: true,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	web.Use(unitelhttp.FiberMiddleware(tel, unitelhttp.FiberMiddlewareConfig{
		Repanic:         false,
		WaitForDelivery: false,
		Timeout:         5 * time.Second,
		// host for incoming requests
		TraceRequestHeaders: []string{"origin", "x-nonce", "x-signature", "x-signed-by", "sentry-trace", "sentry-baggage"},
		// origin for outgoing requests
		TraceResponseHeaders: []string{"host", "x-nonce", "x-signature", "x-signed-by", "sentry-trace", "sentry-baggage"},
		// IgnoredRoutes:        nil,
		Logger:          zerologr.New(&log.Logger).WithName("http-server"),
		TracePropagator: shouldPropagate,
	}))
	web.Use(unitelhttp.RequestLogger(zerologr.New(&log.Logger).WithName("http-server"), true, true))

	log.Debug().Msg("Registering handlers")

	userHandler.Register(web.Group("/"))
	noteHandler.Register(web.Group("/"))
	followHandler.Register(web.Group("/"))
	metaHandler.Register(web.Group("/"))

	wg := sync.WaitGroup{}
	wg.Add(2)

	// TODO: Run these in separate processes, if wanted
	go func() {
		defer wg.Done()

		log.Debug().Msg("Starting taskqueue consumer")

		tasks.NewHandler(federationService, repos, tel, zerologr.New(&log.Logger).WithName("task-handler")).
			Register(tq)

		if err := tq.StartConsumer(context.Background(), "consumer"); err != nil {
			log.Fatal().Err(err).Msg("failed to start taskqueue client")
		}
	}()

	go func() {
		defer wg.Done()

		log.Debug().Msg("Starting server")

		addr := fmt.Sprintf(":%d", config.C.Port)

		var err error
		if config.C.TLSKey != nil {
			err = web.ListenTLS(addr, *config.C.TLSCert, *config.C.TLSKey)
		} else {
			err = web.Listen(addr)
		}
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	<-signalCh

	log.Info().Msg("Shutting down")

	tq.Close()
	if err := web.Shutdown(); err != nil {
		log.Error().Err(err).Msg("Failed to shutdown server")
	}

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

func initServerActor(db *ent.Client, telemetry *unitel.Telemetry) error {
	s := telemetry.StartSpan(context.Background(), "function", "main.initServerActor")
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
