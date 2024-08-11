package note_handler

import (
	"github.com/go-logr/logr"
	"github.com/gofiber/fiber/v2"
	"github.com/lysand-org/versia-go/config"
	"github.com/lysand-org/versia-go/internal/service"
	"github.com/lysand-org/versia-go/internal/validators"
	"github.com/lysand-org/versia-go/pkg/webfinger"
)

type Handler struct {
	noteService   service.NoteService
	bodyValidator validators.BodyValidator

	hostMeta webfinger.HostMeta

	log logr.Logger
}

func New(noteService service.NoteService, bodyValidator validators.BodyValidator, log logr.Logger) *Handler {
	return &Handler{
		noteService:   noteService,
		bodyValidator: bodyValidator,

		hostMeta: webfinger.NewHostMeta(config.C.PublicAddress),

		log: log.WithName("users"),
	}
}

func (i *Handler) Register(r fiber.Router) {
	r.Get("/api/app/notes/:id", i.GetNote)
	r.Post("/api/app/notes/", i.CreateNote)
}
