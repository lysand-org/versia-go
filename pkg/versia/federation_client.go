package versia

import (
	"fmt"
	"github.com/go-logr/logr"
	"github.com/lysand-org/versia-go/pkg/protoretry"
	"net/http"
	"net/url"
	"time"
)

type ResponseError struct {
	StatusCode int
	URL        *url.URL
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("error from %s: %d", e.URL, e.StatusCode)
}

type FederationClient struct {
	log   logr.Logger
	httpC *http.Client
	hc    *protoretry.Client
}

type Opt func(c *FederationClient)

func WithHTTPClient(h *http.Client) Opt {
	return func(c *FederationClient) {
		c.httpC = h
	}
}

func WithLogger(l logr.Logger) Opt {
	return func(c *FederationClient) {
		c.log = l
	}
}

func NewClient(opts ...Opt) *FederationClient {
	c := &FederationClient{
		httpC: http.DefaultClient,
		log:   logr.Discard(),
	}

	for _, opt := range opts {
		opt(c)
	}

	c.httpC.Transport = &federationClientHTTPTransport{
		inner:     c.httpC.Transport,
		useragent: "github.com/lysand-org/versia-go/pkg/lysand#0.0.1",
	}

	c.hc = protoretry.New(c.httpC)

	return c
}

type federationClientHTTPTransport struct {
	inner     http.RoundTripper
	useragent string
	l         logr.Logger
}

func (t *federationClientHTTPTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", t.useragent)

	if req.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	start := time.Now()
	res, err := t.inner.RoundTrip(req)
	elapsed := time.Since(start)
	if err == nil {
		t.l.V(1).Info("fetch succeeded", "url", req.URL.String(), "status", res.StatusCode, "duration", elapsed)
	} else {
		t.l.V(1).Error(err, "fetch failed", "url", req.URL.String(), "duration", elapsed)
	}

	return res, err
}
