package main

import (
	"context"
	"fmt"
	"git.devminer.xyz/devminer/unitel"
	"git.devminer.xyz/devminer/unitel/unitelhttp"
	"github.com/versia-pub/versia-go/internal/api_schema"
	"github.com/versia-pub/versia-go/internal/handlers/follow_handler"
	"github.com/versia-pub/versia-go/internal/handlers/meta_handler"
	"github.com/versia-pub/versia-go/internal/handlers/note_handler"
	"github.com/versia-pub/versia-go/internal/service"
	"github.com/versia-pub/versia-go/internal/validators"
	"net/http"
	"sync"
	"time"

	"github.com/go-logr/zerologr"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
	"github.com/versia-pub/versia-go/ent"
	"github.com/versia-pub/versia-go/internal/config"
	"github.com/versia-pub/versia-go/internal/handlers/user_handler"
)

func shouldPropagate(r *http.Request) bool {
	return config.C.ForwardTracesTo.Match([]byte(r.URL.String()))
}

func server(
	ctx context.Context,
	telemetry *unitel.Telemetry,
	database *ent.Client,
	natsConn *nats.Conn,
	federationService service.FederationService,
	requestSigner service.RequestSigner,
	bodyValidator validators.BodyValidator,
	requestValidator validators.RequestValidator,
	userService service.UserService,
	noteService service.NoteService,
	followService service.FollowService,
	instanceMetadataService service.InstanceMetadataService,
	inboxService service.InboxService,
) error {
	// Handlers

	userHandler := user_handler.New(federationService, requestSigner, userService, inboxService, bodyValidator, requestValidator, zerologr.New(&log.Logger).WithName("user-handler"))
	noteHandler := note_handler.New(noteService, bodyValidator, requestSigner, zerologr.New(&log.Logger).WithName("notes-handler"))
	followHandler := follow_handler.New(followService, federationService, zerologr.New(&log.Logger).WithName("follow-handler"))
	metaHandler := meta_handler.New(instanceMetadataService, zerologr.New(&log.Logger).WithName("meta-handler"))

	// Initialization

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

	web.Use(unitelhttp.FiberMiddleware(telemetry, unitelhttp.FiberMiddlewareConfig{
		Repanic:         true,
		WaitForDelivery: false,
		Timeout:         5 * time.Second,
		// host for incoming requests
		TraceRequestHeaders: []string{"origin", "x-nonce", "x-signature", "x-signed-by", "sentry-trace", "sentry-baggage"},
		// origin for outgoing requests
		TraceResponseHeaders: []string{"host", "x-nonce", "x-signature", "x-signed-by", "sentry-trace", "sentry-baggage"},
		IgnoredRoutes:        []string{"/api/health"},
		Logger:               zerologr.New(&log.Logger).WithName("http-server"),
		TracePropagator:      shouldPropagate,
	}))
	web.Use(unitelhttp.RequestLogger(zerologr.New(&log.Logger).WithName("http-server"), true, true))

	log.Debug().Msg("Registering handlers")

	web.Get("/api/health", healthCheck(database, natsConn))

	userHandler.Register(web.Group("/"))
	noteHandler.Register(web.Group("/"))
	followHandler.Register(web.Group("/"))
	metaHandler.Register(web.Group("/"))

	wg := sync.WaitGroup{}
	wg.Add(2)

	addr := fmt.Sprintf(":%d", config.C.Port)

	log.Info().Str("addr", addr).Msg("Starting server")

	go func() {
		<-ctx.Done()

		if err := web.Shutdown(); err != nil {
			log.Error().Err(err).Msg("Failed to shutdown server")
		}
	}()

	var err error
	if config.C.TLSKey != nil {
		err = web.ListenTLS(addr, *config.C.TLSCert, *config.C.TLSKey)
	} else {
		err = web.Listen(addr)
	}

	return err
}

func healthCheck(db *ent.Client, nc *nats.Conn) fiber.Handler {
	return func(c *fiber.Ctx) error {
		dbWorking := true
		if err := db.Ping(); err != nil {
			log.Error().Err(err).Msg("Database healthcheck failed")
			dbWorking = false
		}

		natsWorking := true
		if status := nc.Status(); status != nats.CONNECTED {
			log.Error().Str("status", status.String()).Msg("NATS healthcheck failed")
			natsWorking = false
		}

		if dbWorking && natsWorking {
			return c.SendString("lookin' good")
		}

		return api_schema.ErrInternalServerError(map[string]any{
			"database": dbWorking,
			"nats":     natsWorking,
		})
	}
}
