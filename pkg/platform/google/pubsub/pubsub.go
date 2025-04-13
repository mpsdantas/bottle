package pubsub

import (
	"context"
	"encoding/json"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/mpsdantas/bottle/pkg/core/log"
)

type Event interface {
	Attributes() map[string]string
}

type Client struct {
	clt        *pubsub.Client
	retries    int
	retryDelay time.Duration
}

func New(opt ...OptionFunc) *Client {
	ctx := context.Background()
	opts := options{
		project:    "",
		retries:    100,
		retryDelay: 1 * time.Second,
	}

	for _, optionFunc := range opt {
		optionFunc(&opts)
	}

	if opts.project == "" {
		log.Fatal(ctx, "project is required")
	}

	client, err := pubsub.NewClient(ctx, opts.project)
	if err != nil {
		log.Panic(ctx, "could not start google pubsub",
			log.Err(err),
		)
	}

	return &Client{
		clt:        client,
		retries:    opts.retries,
		retryDelay: opts.retryDelay,
	}
}

func (c *Client) Close() {
	ctx := context.Background()
	err := c.clt.Close()
	if err != nil {
		log.Error(ctx, "could not close pubsub",
			log.Err(err),
		)
	}
}

func (c *Client) Topic(id string) *Topic {
	return &Topic{
		tp:         c.clt.Topic(id),
		retries:    c.retries,
		retryDelay: c.retryDelay,
	}
}

type Topic struct {
	tp         *pubsub.Topic
	retries    int
	retryDelay time.Duration
}

func (t *Topic) Publish(ctx context.Context, value interface{}) {
	go func(c context.Context) {
		if err := t.PublishSync(c, value); err != nil {
			log.Error(c, "could not publish message after max retries",
				log.String("topic", t.tp.ID()),
				log.Int("retries", t.retries),
				log.Err(err),
			)
		}
	}(WithoutCancel(ctx))
}

func (t *Topic) PublishSync(ctx context.Context, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	for i := 0; i < t.retries; i++ {
		message := &pubsub.Message{
			Data: data,
		}

		if attributes, ok := value.(Event); ok {
			message.Attributes = attributes.Attributes()
		}

		result := t.tp.Publish(ctx, message)

		_, err := result.Get(ctx)
		if err != nil {
			log.Error(ctx, "could not publish message",
				log.String("topic", t.tp.ID()),
				log.Int("attempt", i+1),
				log.Err(err),
			)

			time.Sleep(t.retryDelay)
			continue
		}

		return nil
	}

	return err
}
