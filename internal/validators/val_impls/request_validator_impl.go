package val_impls

import (
	"context"
	"errors"
	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/gofiber/fiber/v2"
	"github.com/versia-pub/versia-go/internal/repository"
	"github.com/versia-pub/versia-go/internal/utils"
	"github.com/versia-pub/versia-go/internal/validators"
	versiacrypto "github.com/versia-pub/versia-go/pkg/versia/crypto"
	versiautils "github.com/versia-pub/versia-go/pkg/versia/utils"
	"net/http"
)

var (
	ErrInvalidSignature    = errors.New("invalid signature")
	ErrInvalidOriginHeader = errors.New("invalid origin header")

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

	origin := r.Header.Get("Origin")
	if origin != "" {
		return ErrInvalidOriginHeader
	}

	l := i.log.WithValues("url", r.URL.Path)

	r = r.WithContext(ctx)

	fedHeaders, err := versiacrypto.ExtractFederationHeaders(r.Header)
	if err != nil {
		return err
	}

	var key *versiacrypto.SPKIPublicKey
	var signerURI *versiautils.URL
	if fedHeaders.SignedBy == nil {
		metadata, err := i.repositories.InstanceMetadata().Resolve(ctx, origin)
		if err != nil {
			return err
		}
		signerURI = metadata.URI
	} else {
		user, err := i.repositories.Users().Resolve(ctx, versiautils.URLFromStd(fedHeaders.SignedBy))
		if err != nil {
			return err
		}
		signerURI = user.URI
	}

	l = l.WithValues("signer", signerURI)

	body, err := utils.CopyBody(r)
	if err != nil {
		return err
	}

	if !(versiacrypto.Verifier{PublicKey: key}).Verify(r.Method, r.URL, body, fedHeaders) {
		l.WithCallDepth(1).Info("signature verification failed")
		s.CaptureError(ErrInvalidSignature)

		return ErrInvalidSignature
	} else {
		l.V(2).Info("signature verification succeeded")
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
