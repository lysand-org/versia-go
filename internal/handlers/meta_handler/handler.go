package meta_handler

import (
	"github.com/go-logr/logr"
	"github.com/gofiber/fiber/v2"
	"github.com/lysand-org/versia-go/config"
	"github.com/lysand-org/versia-go/internal/service"
	"github.com/lysand-org/versia-go/pkg/webfinger"
)

type Handler struct {
	instanceMetadataService service.InstanceMetadataService

	hostMeta webfinger.HostMeta

	log logr.Logger
}

func New(instanceMetadataService service.InstanceMetadataService, log logr.Logger) *Handler {
	return &Handler{
		instanceMetadataService: instanceMetadataService,

		hostMeta: webfinger.NewHostMeta(config.C.PublicAddress),

		log: log.WithName("users"),
	}
}

func (i *Handler) Register(r fiber.Router) {
	r.Get("/.well-known/versia", i.GetLysandInstanceMetadata)
	r.Get("/.well-known/versia/admins", i.GetLysandInstanceMetadata)
	r.Get("/.well-known/versia/moderators", i.GetLysandInstanceMetadata)

	// Webfinger host meta spec
	r.Get("/.well-known/host-meta", i.GetHostMeta)
	r.Get("/.well-known/host-meta.json", i.GetHostMetaJSON)
}
