package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/mpsdantas/bottle/pkg/log"
	"google.golang.org/api/option"
)

type Client struct {
	clt *pubsub.Client
}

func New(ctx context.Context, project string, opts ...option.ClientOption) *Client {
	client, err := pubsub.NewClient(ctx, project, opts...)
	if err != nil {
		log.Panic(ctx, "could not start google pubsub",
			log.Err(err),
		)
	}

	return &Client{
		clt: client,
	}
}

func (c *Client) Close() {
	err := c.clt.Close()
	if err != nil {
		log.Error(context.Background(), "could not close pubsub",
			log.Err(err),
		)
	}
}

func (c *Client) Topic(id string) *Topic {
	return &Topic{
		tp: c.clt.Topic(id),
	}
}
