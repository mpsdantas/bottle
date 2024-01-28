package idtoken

import (
	"context"
	"io"
	"net/http"

	"github.com/mpsdantas/bottle/pkg/log"
	"google.golang.org/api/idtoken"
)

func NewHTTPClient(base string) *http.Client {
	ctx := context.Background()

	client, err := idtoken.NewClient(ctx, base)
	if err != nil {
		log.Fatal(ctx, "could not get client", log.Err(err))
	}

	return client
}

type Response struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

type Client struct {
	http *http.Client
}

func New(http *http.Client) *Client {
	return &Client{
		http: http,
	}
}

func (c *Client) Do(ctx context.Context, req *http.Request) (*Response, error) {
	requestid, ok := ctx.Value("requestid").(string)
	if requestid != "" && ok {
		req.Header.Add("x-request-id", requestid)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       body,
	}, nil
}
