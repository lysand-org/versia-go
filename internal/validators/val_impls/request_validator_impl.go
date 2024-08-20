package val_impls

import (
	"context"
	"errors"
	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/gofiber/fiber/v2"
	"github.com/lysand-org/versia-go/internal/repository"
	"github.com/lysand-org/versia-go/internal/utils"
	"github.com/lysand-org/versia-go/internal/validators"
	"github.com/lysand-org/versia-go/pkg/lysand"
	"net/http"
)

var (
	ErrInvalidSignature = errors.New("invalid signature")

	_ validators.RequestValidator = (*RequestValidatorImpl)(nil)
)

type RequestValidatorImpl struct {
	repositories repository.Manager

	telemetry *unitel.Telemetry
	log       logr.Logger
}

func NewRequestValidator(repositories repository.Manager, telemetry *unitel.Telemetry, log logr.Logger) *RequestValidatorImpl {
	return &RequestValidatorImpl{
		repositories: repositories,

		telemetry: telemetry,
		log:       log,
	}
}

func (i RequestValidatorImpl) Validate(ctx context.Context, r *http.Request) error {
	s := i.telemetry.StartSpan(ctx, "function", "val_impls/RequestValidatorImpl.Validate")
	defer s.End()
	ctx = s.Context()

	r = r.WithContext(ctx)

	fedHeaders, err := lysand.ExtractFederationHeaders(r.Header)
	if err != nil {
		return err
	}

	// TODO: Fetch user from database instead of using the URI
	user, err := i.repositories.Users().Resolve(ctx, lysand.URLFromStd(fedHeaders.SignedBy))
	if err != nil {
		return err
	}

	body, err := utils.CopyBody(r)
	if err != nil {
		return err
	}

	if !(lysand.Verifier{PublicKey: user.PublicKey.Key}).Verify(r.Method, r.URL, body, fedHeaders) {
		i.log.WithCallDepth(1).Info("signature verification failed", "user", user.URI, "url", r.URL.Path)
		s.CaptureError(ErrInvalidSignature)

		return ErrInvalidSignature
	} else {
		i.log.V(2).Info("signature verification succeeded", "user", user.URI, "url", r.URL.Path)
	}

	return nil
}

func (i RequestValidatorImpl) ValidateFiberCtx(ctx context.Context, c *fiber.Ctx) error {
	s := i.telemetry.StartSpan(ctx, "function", "val_impls/RequestValidatorImpl.ValidateFiberCtx")
	defer s.End()
	ctx = s.Context()

	r, err := utils.ConvertToStdRequest(c)
	if err != nil {
		return err
	}

	return i.Validate(ctx, r)
}
