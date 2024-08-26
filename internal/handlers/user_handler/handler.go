package user_handler

import (
	"github.com/go-logr/logr"
	"github.com/gofiber/fiber/v2"
	"github.com/lysand-org/versia-go/internal/service"
	"github.com/lysand-org/versia-go/internal/validators"
)

type Handler struct {
	federationService service.FederationService
	requestSigner     service.RequestSigner

	userService  service.UserService
	inboxService service.InboxService

	bodyValidator    validators.BodyValidator
	requestValidator validators.RequestValidator

	log logr.Logger
}

func New(federationService service.FederationService, requestSigner service.RequestSigner, userService service.UserService, inboxService service.InboxService, bodyValidator validators.BodyValidator, requestValidator validators.RequestValidator, log logr.Logger) *Handler {
	return &Handler{
		federationService: federationService,
		requestSigner:     requestSigner,

		userService:  userService,
		inboxService: inboxService,

		bodyValidator:    bodyValidator,
		requestValidator: requestValidator,

		log: log,
	}
}

func (i *Handler) Register(r fiber.Router) {
	// TODO: Handle this differently
	// There might be other routes that might want to also add their stuff to the robots.txt
	r.Get("/robots.txt", i.RobotsTXT)

	r.Get("/.well-known/webfinger", i.Webfinger)

	r.Get("/api/app/users/search", i.SearchUser)
	r.Get("/api/app/users/:id", i.GetUser)
	r.Post("/api/app/users/", i.CreateUser)

	r.Get("/api/users/:id", i.GetLysandUser)
	r.Post("/api/users/:id/inbox", i.LysandInbox)
}
