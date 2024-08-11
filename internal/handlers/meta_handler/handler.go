package meta_handler

import (
	"github.com/go-logr/logr"
	"github.com/gofiber/fiber/v2"
	"github.com/lysand-org/versia-go/config"
	"github.com/lysand-org/versia-go/pkg/webfinger"
)

type Handler struct {
	hostMeta webfinger.HostMeta

	log logr.Logger
}

func New(log logr.Logger) *Handler {
	return &Handler{
		hostMeta: webfinger.NewHostMeta(config.C.PublicAddress),

		log: log.WithName("users"),
	}
}

func (i *Handler) Register(r fiber.Router) {
	r.Get("/.well-known/lysand", i.GetLysandServerMetadata)
	r.Get("/.well-known/host-meta", i.GetHostMeta)
	r.Get("/.well-known/host-meta.json", i.GetHostMetaJSON)
}
