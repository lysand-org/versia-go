package user_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lysand-org/versia-go/ent"
	"github.com/lysand-org/versia-go/internal/api_schema"
)

func (i *Handler) CreateUser(c *fiber.Ctx) error {
	var req api_schema.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return api_schema.ErrInvalidRequestBody(nil)
	}

	if err := i.bodyValidator.Validate(req); err != nil {
		return err
	}

	u, err := i.userService.NewUser(c.UserContext(), req.Username, req.Password)
	if err != nil {
		// TODO: Wrap this in a custom error
		if ent.IsConstraintError(err) {
			return api_schema.ErrUsernameTaken(nil)
		}

		i.log.Error(err, "Failed to create user", "username", req.Username)

		return api_schema.ErrInternalServerError(nil)
	}

	return c.JSON(api_schema.User{
		ID:       u.ID,
		Username: u.Username,
	})
}
