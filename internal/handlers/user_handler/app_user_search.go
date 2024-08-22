package user_handler

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lysand-org/versia-go/internal/api_schema"
	"github.com/lysand-org/versia-go/pkg/webfinger"
	"net"
	"syscall"
)

func (i *Handler) SearchUser(c *fiber.Ctx) error {
	var req api_schema.SearchUserRequest
	if err := c.QueryParser(&req); err != nil {
		return api_schema.ErrInvalidRequestBody(nil)
	}

	if err := i.bodyValidator.Validate(req); err != nil {
		return err
	}

	u, err := i.userService.Search(c.UserContext(), req)
	if err != nil {
		// TODO: Move into service error
		if errors.Is(err, syscall.ECONNREFUSED) {
			return api_schema.ErrBadRequest(map[string]any{"reason": "Remote server is offline"})
		}

		if errors.Is(err, webfinger.ErrUserNotFound) {
			return api_schema.ErrUserNotFound
		}

		var dnsErr *net.DNSError
		if errors.As(err, &dnsErr) {
			if dnsErr.IsNotFound {
				return api_schema.ErrBadRequest(map[string]any{"reason": fmt.Sprintf("Could not resolve %s", dnsErr.Name)})
			}

			if dnsErr.IsTimeout {
				return api_schema.ErrInternalServerError(map[string]any{"reason": "Local DNS server timed out"})
			}
		}

		i.log.Error(err, "Failed to search for user", "username", req.Username)

		return api_schema.ErrInternalServerError(nil)
	}

	return c.JSON((*api_schema.VersiaUser)(u.ToLysand()))
}
