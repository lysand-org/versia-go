package svc_impls

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/gofiber/fiber/v2"
	"github.com/lysand-org/versia-go/internal/service"
	versiacrypto "github.com/lysand-org/versia-go/pkg/versia/crypto"
	"net/url"
)

var _ service.RequestSigner = (*RequestSignerImpl)(nil)

type RequestSignerImpl struct {
	telemetry *unitel.Telemetry

	log logr.Logger
}

func NewRequestSignerImpl(telemetry *unitel.Telemetry, log logr.Logger) *RequestSignerImpl {
	return &RequestSignerImpl{
		telemetry: telemetry,
		log:       log,
	}
}

func (i *RequestSignerImpl) SignAndSend(c *fiber.Ctx, signer versiacrypto.Signer, body any) error {
	s := i.telemetry.StartSpan(c.UserContext(), "function", "svc_impls/RequestSignerImpl.SignAndSend")
	defer s.End()

	j, err := json.Marshal(body)
	if err != nil {
		return err
	}

	rawNonce := make([]byte, 64)
	if _, err := rand.Read(rawNonce); err != nil {
		return err
	}
	nonce := base64.StdEncoding.EncodeToString(rawNonce)

	uri, err := url.ParseRequestURI(string(c.Request().RequestURI()))
	if err != nil {
		return err
	}

	digest := versiacrypto.SHA256(j)

	d := versiacrypto.NewSignatureData(c.Method(), nonce, uri, digest)

	signed := signer.Sign(*d)
	for k, v := range signed.Headers() {
		c.Set(k, v)
	}

	i.log.V(2).Info("signed response", "digest", base64.StdEncoding.EncodeToString(digest), "nonce", nonce)

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	return c.Send(j)
}
