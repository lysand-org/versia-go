package follow_handler

import (
	"github.com/go-logr/logr"
	"github.com/gofiber/fiber/v2"
	"github.com/versia-pub/versia-go/internal/config"
	"github.com/versia-pub/versia-go/internal/service"
	"github.com/versia-pub/versia-go/pkg/webfinger"
)

type Handler struct {
	followService     service.FollowService
	federationService service.FederationService

	hostMeta webfinger.HostMeta

	log logr.Logger
}

func New(followService service.FollowService, federationService service.FederationService, log logr.Logger) *Handler {
	return &Handler{
		followService:     followService,
		federationService: federationService,

		hostMeta: webfinger.NewHostMeta(config.C.PublicAddress),

		log: log.WithName("users"),
	}
}

func (i *Handler) Register(r fiber.Router) {
	r.Get("/api/follows/:id", i.GetVersiaFollow)
}
