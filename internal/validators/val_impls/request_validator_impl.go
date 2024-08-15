package val_impls

import (
	"bytes"
	"context"
	"errors"
	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/gofiber/fiber/v2"
	"github.com/lysand-org/versia-go/internal/repository"
	"github.com/lysand-org/versia-go/internal/validators"
	"github.com/lysand-org/versia-go/pkg/lysand"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"io"
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
	s := i.telemetry.StartSpan(ctx, "function", "validator/val_impls.RequestValidatorImpl.Validate")
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

	body, err := copyBody(r)
	if err != nil {
		return err
	}

	if !(lysand.Verifier{PublicKey: user.PublicKey}).Verify(r.Method, r.URL, body, fedHeaders) {
		i.log.Info("signature verification failed", "user", user.URI, "url", r.URL.Path)
		s.CaptureError(ErrInvalidSignature)

		return ErrInvalidSignature
	} else {
		i.log.V(2).Info("signature verification succeeded", "user", user.URI, "url", r.URL.Path)
	}

	return nil
}

func (i RequestValidatorImpl) ValidateFiberCtx(ctx context.Context, c *fiber.Ctx) error {
	s := i.telemetry.StartSpan(ctx, "function", "validator/val_impls.RequestValidatorImpl.ValidateFiberCtx")
	defer s.End()
	ctx = s.Context()

	r, err := convertToStdRequest(c)
	if err != nil {
		return err
	}

	return i.Validate(ctx, r)
}

func convertToStdRequest(c *fiber.Ctx) (*http.Request, error) {
	stdReq := &http.Request{}
	if err := fasthttpadaptor.ConvertRequest(c.Context(), stdReq, true); err != nil {
		return nil, err
	}

	return stdReq, nil
}

func copyBody(req *http.Request) ([]byte, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	if err := req.Body.Close(); err != nil {
		return nil, err
	}

	req.Body = io.NopCloser(bytes.NewBuffer(body))
	return body, nil
}
