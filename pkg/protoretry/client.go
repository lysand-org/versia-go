package protoretry

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"syscall"
)

type Client struct {
	base *http.Client
}

func New(base *http.Client) *Client {
	return &Client{base}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	if res, err := c.base.Do(req); err == nil {
		return res, nil
	} else if !errors.Is(err, syscall.ECONNREFUSED) {
		return nil, err
	}

	req.URL.Scheme = "http"
	return c.base.Do(req)
}

func (c *Client) GET(ctx context.Context, u *url.URL) ([]byte, *http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	return c.DoReq(req)
}

func (c *Client) POST(ctx context.Context, u *url.URL, reqBody io.Reader) ([]byte, *http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), reqBody)
	if err != nil {
		return nil, nil, err
	}

	return c.DoReq(req)
}

func (c *Client) DoReq(req *http.Request) ([]byte, *http.Response, error) {
	res, err := c.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	return resBody, res, nil
}
