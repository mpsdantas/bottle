package task

import (
	"context"
	"fmt"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"github.com/mpsdantas/bottle/pkg/core/log"
)

type Option func(*Client)

func WithProjectID(projectID string) Option {
	return func(c *Client) {
		c.projectID = projectID
	}
}

func WithLocationID(locationID string) Option {
	return func(c *Client) {
		c.locationID = locationID
	}
}

func WithQueue(queue string) Option {
	return func(c *Client) {
		c.queueID = queue
	}
}

func WithServiceAccount(serviceAccount string) Option {
	return func(c *Client) {
		c.serviceAccount = serviceAccount
	}
}

type Client struct {
	clt            *cloudtasks.Client
	projectID      string
	locationID     string
	queueID        string
	serviceAccount string
}

func New(opts ...Option) *Client {
	ctx := context.Background()

	client, err := cloudtasks.NewClient(ctx)
	if err != nil {
		log.Panic(ctx, "could not start google cloud task",
			log.Err(err),
		)
	}

	c := &Client{
		clt: client,
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.projectID == "" || c.locationID == "" || c.queueID == "" {
		log.Panic(ctx, "missing required cloud tasks configuration",
			log.String("project_id", c.projectID),
			log.String("location_id", c.locationID),
			log.String("queue_id", c.queueID),
		)
	}

	return c
}

func (c *Client) queuePath() string {
	return fmt.Sprintf("projects/%s/locations/%s/queues/%s", c.projectID, c.locationID, c.queueID)
}
