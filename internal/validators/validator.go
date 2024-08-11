package validators

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type BodyValidator interface {
	Validate(v any) error
}

type RequestValidator interface {
	Validate(ctx context.Context, r *http.Request) error
	ValidateFiberCtx(ctx context.Context, c *fiber.Ctx) error
}
