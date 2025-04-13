package rest

import (
	"context"
	"io"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Rest struct {
	client HTTPClient
}

func New(client HTTPClient) *Rest {
	return &Rest{
		client: client,
	}
}

type Response struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

func (r *Rest) Do(ctx context.Context, req *http.Request) (*Response, error) {
	req = req.WithContext(ctx)

	if appName, ok := ctx.Value("application").(string); ok && appName != "" {
		req.Header.Set("x-application", appName)
	}

	if requestID, ok := ctx.Value("requestid").(string); ok && requestID != "" {
		req.Header.Set("x-request-id", requestID)
	}

	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: res.StatusCode,
		Headers:    res.Header,
		Body:       body,
	}, nil
}
