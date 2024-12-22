package task

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
	"github.com/mpsdantas/bottle/pkg/log"
	"github.com/mpsdantas/bottle/pkg/pubsub"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Options struct {
	ProjectID      string
	LocationID     string
	Queue          string
	ServiceAccount string
}

type Client struct {
	clt            *cloudtasks.Client
	queue          string
	serviceAccount string
}

func New(ctx context.Context, opts *Options) *Client {
	client, err := cloudtasks.NewClient(ctx)
	if err != nil {
		log.Panic(ctx, "could not start google cloud task",
			log.Err(err),
		)
	}

	return &Client{
		clt:            client,
		queue:          fmt.Sprintf("projects/%s/locations/%s/queues/%s", opts.ProjectID, opts.LocationID, opts.Queue),
		serviceAccount: opts.ServiceAccount,
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
	baseUrl, err := getBaseURL(url)
	if err != nil {
		log.Error(ctx, "could not get base url",
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
				AuthorizationHeader: &cloudtaskspb.HttpRequest_OidcToken{
					OidcToken: &cloudtaskspb.OidcToken{
						ServiceAccountEmail: c.serviceAccount,
						Audience:            baseUrl,
					},
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

func getBaseURL(rawURL string) (string, error) {
	// Parse a URL string into a *url.URL structure
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// Reconstruct the base URL
	baseURL := fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)
	return baseURL, nil
}
