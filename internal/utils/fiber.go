package utils

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"io"
	"net/http"
)

func ConvertToStdRequest(c *fiber.Ctx) (*http.Request, error) {
	stdReq := &http.Request{}
	if err := fasthttpadaptor.ConvertRequest(c.Context(), stdReq, true); err != nil {
		return nil, err
	}

	return stdReq, nil
}

func CopyBody(req *http.Request) ([]byte, error) {
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
