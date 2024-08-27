package protoretry

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
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
	l := log.With().Str("url", req.URL.String()).Logger()

	l.Debug().Msg("fetch")
	res, err := c.base.Do(req)
	if err != nil {
		var urlErr *url.Error

		if errors.Is(err, syscall.ECONNREFUSED) {
			goto onTlsError
		}

		if errors.As(err, &urlErr) && errors.Is(urlErr.Err, http.ErrSchemeMismatch) {
			goto onTlsError
		}

		l.Error().Err(err).Msg("failed")
		return nil, err

	onTlsError:
		l.Debug().Msg("downgrading to http")
		req.URL.Scheme = "http"

		l.Debug().Msg("fetch")
		if res, err := c.base.Do(req); err == nil {
			return res, nil
		} else {
			l.Error().Err(err).Msg("failed")
			return nil, err
		}
	}

	l.Debug().Msg("fetched")
	return res, nil
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
