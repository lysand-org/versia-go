package lysand

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"io"
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

	return c
}

func (c *FederationClient) rawGET(ctx context.Context, uri *url.URL) (*http.Response, []byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", uri.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	return c.doReq(req)
}

func (c *FederationClient) rawPOST(ctx context.Context, uri *url.URL, body io.Reader) (*http.Response, []byte, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", uri.String(), body)
	if err != nil {
		return nil, nil, err
	}

	return c.doReq(req)
}

func (c *FederationClient) doReq(req *http.Request) (*http.Response, []byte, error) {
	resp, err := c.httpC.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return resp, nil, &ResponseError{
			StatusCode: resp.StatusCode,
			URL:        req.URL,
		}
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, nil, err
	}

	return resp, respBody, nil
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
