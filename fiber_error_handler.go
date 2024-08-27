package main

import (
	"errors"

	"git.devminer.xyz/devminer/unitel"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/versia-pub/versia-go/internal/api_schema"
)

func fiberErrorHandler(c *fiber.Ctx, err error) error {
	var fiberErr *fiber.Error
	var apiErr *api_schema.APIError

	if errors.As(err, &fiberErr) {
		apiErr = api_schema.NewAPIError(fiberErr.Code, fiberErr.Error())(nil)
	} else if errors.As(err, &apiErr) {
		log.Error().Err(apiErr).Msg("API error")
	} else {
		if hub := unitel.GetHubFromFiberContext(c); hub != nil {
			hub.CaptureException(err)
		}

		log.Error().Err(err).Msg("Unhandled error")
		apiErr = api_schema.NewAPIError(fiber.StatusInternalServerError, "Internal Server Error")(nil)
	}

	log.Error().Err(apiErr).Msg("Error")

	return c.Status(apiErr.StatusCode).JSON(api_schema.NewFailedAPIResponse[any](apiErr))
}
