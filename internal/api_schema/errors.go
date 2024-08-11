package api_schema

var (
	ErrBadRequest         = NewAPIError(400, "Bad request")
	ErrInvalidRequestBody = NewAPIError(400, "Invalid request body")
	ErrUnauthorized       = NewAPIError(401, "Unauthorized")
	ErrForbidden          = NewAPIError(403, "Forbidden")
	ErrNotFound           = NewAPIError(404, "Not found")
	ErrConflict           = NewAPIError(409, "Conflict")
	ErrUsernameTaken      = NewAPIError(409, "Username is taken")
	ErrRateLimitExceeded  = NewAPIError(429, "Rate limit exceeded")

	ErrInternalServerError = NewAPIError(500, "Internal server error")
	ErrNotImplemented      = NewAPIError(501, "Not implemented")
)
