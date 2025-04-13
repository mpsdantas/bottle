package task

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskOption func(*cloudtaskspb.Task)

func WithSchedule(t time.Time) TaskOption {
	return func(task *cloudtaskspb.Task) {
		task.ScheduleTime = timestamppb.New(t)
	}
}

func (c *Client) Create(ctx context.Context, url string, data any, opts ...TaskOption) error {
	baseURL, err := getBaseURL(url)
	if err != nil {
		return fmt.Errorf("could not parse base url from %s: %w", url, err)
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("could not marshal payload for url %s: %w", url, err)
	}

	task := &cloudtaskspb.Task{
		MessageType: &cloudtaskspb.Task_HttpRequest{
			HttpRequest: &cloudtaskspb.HttpRequest{
				HttpMethod: cloudtaskspb.HttpMethod_POST,
				Url:        url,
				AuthorizationHeader: &cloudtaskspb.HttpRequest_OidcToken{
					OidcToken: &cloudtaskspb.OidcToken{
						ServiceAccountEmail: c.serviceAccount,
						Audience:            baseURL,
					},
				},
				Body: payload,
			},
		},
	}

	for _, opt := range opts {
		opt(task)
	}

	_, err = c.clt.CreateTask(ctx, &cloudtaskspb.CreateTaskRequest{
		Parent: c.queuePath(),
		Task:   task,
	})

	if err != nil {
		return fmt.Errorf("failed to create task for url %s: %w", url, err)
	}

	return nil
}

func getBaseURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host), nil
}
