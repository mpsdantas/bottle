package task

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
	"github.com/mpsdantas/bottle/pkg/log"
	"github.com/mpsdantas/bottle/pkg/pubsub"
	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Client struct {
	clt    *cloudtasks.Client
	tokens map[string]*oauth2.Token
	queue  string
}

func New(ctx context.Context, projectID string, locationID string, queue string) *Client {
	client, err := cloudtasks.NewClient(ctx)
	if err != nil {
		log.Panic(ctx, "could not start google cloud task",
			log.Err(err),
		)
	}

	return &Client{
		clt:    client,
		tokens: map[string]*oauth2.Token{},
		queue:  fmt.Sprintf("projects/%s/locations/%s/queues/%s", projectID, locationID, queue),
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

func (c *Client) Create(ctx context.Context, url string, data any) {
	go func(ctx context.Context, lurl string, ldata any) {
		c.create(ctx, lurl, ldata, nil)
	}(pubsub.WithoutCancel(ctx), url, data)
}

func (c *Client) Schedule(ctx context.Context, url string, data any, schedule time.Time) {
	go func(ctx context.Context, lurl string, ldata any, lschedule time.Time) {
		c.create(ctx, lurl, ldata, &lschedule)
	}(pubsub.WithoutCancel(ctx), url, data, schedule)
}

func (c *Client) create(ctx context.Context, url string, data any, schedule *time.Time) {
	token, err := c.getToken(ctx, url)
	if err != nil {
		log.Error(ctx, "could not get token",
			log.Err(err),
		)
		return
	}

	payload, err := json.Marshal(data)
	if err != nil {
		log.Error(ctx,
			"could not marshal payload",
			log.Err(err),
		)
	}

	task := &cloudtaskspb.Task{
		MessageType: &cloudtaskspb.Task_HttpRequest{
			HttpRequest: &cloudtaskspb.HttpRequest{
				HttpMethod: cloudtaskspb.HttpMethod_POST,
				Url:        url,
				Headers: map[string]string{
					"Authorization": fmt.Sprintf("Bearer %s", token.AccessToken),
				},
				Body: payload,
			},
		},
	}

	if schedule != nil {
		task.ScheduleTime = timestamppb.New(*schedule)
	}

	_, err = c.clt.CreateTask(ctx, &cloudtaskspb.CreateTaskRequest{
		Parent: c.queue,
		Task:   task,
	})
	if err != nil {
		log.Error(ctx,
			"could not schedule task",
			log.Err(err),
		)
	}
}

func (c *Client) getToken(ctx context.Context, audience string) (*oauth2.Token, error) {
	value, ok := c.tokens[audience]
	if ok {
		if value.Valid() {
			return value, nil
		}
	}

	tokenSource, err := idtoken.NewTokenSource(ctx, audience)
	if err != nil {
		return nil, fmt.Errorf("could not create token source: %w", err)
	}

	token, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("could not get token: %w", err)
	}

	c.tokens[audience] = token

	return token, nil
}
