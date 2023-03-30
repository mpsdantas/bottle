package rest

import (
	"context"
	"time"

	"github.com/valyala/fasthttp"
)

//go:generate mockgen -source=./rest.go -package=rest -destination=./rest_mock.go
type Client interface {
	Do(ctx context.Context, req *fasthttp.Request) (*Response, error)
}

type Response struct {
	StatusCode int
	Headers    *fasthttp.ResponseHeader
	Body       []byte
}

type rest struct {
	application string
}

func (c *rest) Do(ctx context.Context, req *fasthttp.Request) (*Response, error) {
	req.Header.Add("x-application", c.application)
	requestid, ok := ctx.Value("requestid").(string)
	if requestid != "" && ok {
		req.Header.Add("x-request-id", requestid)
	}

	res := fasthttp.AcquireResponse()
	err := fasthttp.DoTimeout(req, res, 40*time.Second)

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	return &Response{
		StatusCode: res.StatusCode(),
		Headers:    &res.Header,
		Body:       res.Body(),
	}, err
}

func New(name string) Client {
	return &rest{
		application: name,
	}
}
