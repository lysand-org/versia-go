package svc_impls

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/versia-pub/versia-go/internal/entity"
	"github.com/versia-pub/versia-go/internal/service"
	"github.com/versia-pub/versia-go/pkg/protoretry"
	"github.com/versia-pub/versia-go/pkg/versia"
	versiacrypto "github.com/versia-pub/versia-go/pkg/versia/crypto"
	versiautils "github.com/versia-pub/versia-go/pkg/versia/utils"
	"github.com/versia-pub/versia-go/pkg/webfinger"
	"net/http"
	"net/url"
)

var (
	_ service.FederationService = (*FederationServiceImpl)(nil)

	ErrSignatureValidationFailed = errors.New("signature validation failed")
)

type FederationServiceImpl struct {
	httpC *protoretry.Client

	telemetry *unitel.Telemetry

	log logr.Logger
}

func NewFederationServiceImpl(httpClient *http.Client, telemetry *unitel.Telemetry, log logr.Logger) *FederationServiceImpl {
	return &FederationServiceImpl{
		httpC:     protoretry.New(httpClient),
		telemetry: telemetry,
		log:       log,
	}
}

func (i *FederationServiceImpl) GetUser(ctx context.Context, uri *versiautils.URL) (*versia.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "svc_impls/FederationServiceImpl.GetUser").
		AddAttribute("userURI", uri.String())
	defer s.End()
	ctx = s.Context()

	body, resp, err := i.httpC.GET(ctx, uri.ToStd())
	if err != nil {
		s.SetSimpleStatus(unitel.Error, err.Error())
		return nil, err
	}

	u := &versia.User{}
	if err := json.Unmarshal(body, u); err != nil {
		s.SetSimpleStatus(unitel.Error, err.Error())
		return nil, err
	}

	fedHeaders, err := versiacrypto.ExtractFederationHeaders(resp.Header)
	if err != nil {
		s.SetSimpleStatus(unitel.Error, err.Error())
		return nil, err
	}

	v := versiacrypto.Verifier{PublicKey: u.PublicKey.Key.Key}
	if !v.Verify("GET", uri.ToStd(), body, fedHeaders) {
		s.SetSimpleStatus(unitel.Error, ErrSignatureValidationFailed.Error())
		i.log.V(1).Error(ErrSignatureValidationFailed, "signature validation failed", "user", u.URI.String())
		return nil, ErrSignatureValidationFailed
	}

	s.SetSimpleStatus(unitel.Ok, "")
	i.log.V(2).Info("signature verification succeeded", "user", u.URI.String())

	return u, nil
}

func (i *FederationServiceImpl) DiscoverUser(ctx context.Context, baseURL, username string) (*webfinger.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "svc_impls/FederationServiceImpl.DiscoverUser").
		AddAttribute("baseURL", baseURL).
		AddAttribute("username", username)
	defer s.End()
	ctx = s.Context()

	wf, err := webfinger.Discover(i.httpC, ctx, baseURL, username)
	if err != nil {
		s.SetSimpleStatus(unitel.Error, err.Error())
		return nil, err
	}

	s.SetSimpleStatus(unitel.Ok, "")

	return wf, nil
}

type ResponseError struct {
	StatusCode int
	URL        *url.URL
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("error from %s: %d", e.URL, e.StatusCode)
}

func (i *FederationServiceImpl) DiscoverInstance(ctx context.Context, baseURL string) (*versia.InstanceMetadata, error) {
	s := i.telemetry.StartSpan(ctx, "function", "svc_impls/FederationServiceImpl.DiscoverInstance").
		AddAttribute("baseURL", baseURL)
	defer s.End()
	ctx = s.Context()

	body, resp, err := i.httpC.GET(ctx, &url.URL{Scheme: "https", Host: baseURL, Path: "/.well-known/versia"})
	if err != nil {
		s.SetSimpleStatus(unitel.Error, err.Error())
		return nil, err
	} else if resp.StatusCode >= http.StatusBadRequest {
		s.SetSimpleStatus(unitel.Error, fmt.Sprintf("unexpected response code: %d", resp.StatusCode))
		return nil, &ResponseError{StatusCode: resp.StatusCode, URL: resp.Request.URL}
	}

	var metadata versia.InstanceMetadata
	if err := json.Unmarshal(body, &metadata); err != nil {
		s.SetSimpleStatus(unitel.Error, err.Error())
		return nil, err
	}

	s.SetSimpleStatus(unitel.Ok, "")

	return &metadata, nil
}

func (i *FederationServiceImpl) SendToInbox(ctx context.Context, author *entity.User, user *entity.User, object any) ([]byte, error) {
	s := i.telemetry.StartSpan(ctx, "function", "svc_impls/FederationServiceImpl.SendToInbox").
		SetUser(uint64(author.ID.ID()), author.Username, "", "").
		AddAttribute("author", author.ID).
		AddAttribute("authorURI", author.URI).
		AddAttribute("target", user.ID).
		AddAttribute("targetURI", user.URI)
	defer s.End()
	ctx = s.Context()

	uri := user.Inbox.ToStd()

	body, err := json.Marshal(object)
	if err != nil {
		s.SetSimpleStatus(unitel.Error, err.Error())
		return nil, err
	}

	nonce := make([]byte, 32)
	if _, err := rand.Read(nonce); err != nil {
		s.SetSimpleStatus(unitel.Error, err.Error())
		return nil, err
	}

	sigData := versiacrypto.NewSignatureData("POST", base64.StdEncoding.EncodeToString(nonce), uri, versiacrypto.SHA256(body))
	sig := author.Signer.Sign(*sigData)

	req, err := http.NewRequestWithContext(ctx, "POST", uri.String(), bytes.NewReader(body))
	if err != nil {
		s.SetSimpleStatus(unitel.Error, err.Error())
		return nil, err
	}

	sig.Inject(req.Header)

	body, _, err = i.httpC.DoReq(req)
	if err != nil {
		s.SetSimpleStatus(unitel.Error, err.Error())
		i.log.Error(err, "Failed to send to inbox", "author", author.URI, "target", user.URI)
		return nil, err
	}

	s.SetSimpleStatus(unitel.Ok, "")

	return body, nil
}
