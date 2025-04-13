package storage

import (
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"
	"github.com/mpsdantas/bottle/pkg/core/log"
)

type Client struct {
	storageClient *storage.Client
	bucketHandler *storage.BucketHandle
}

func New(bucket string) *Client {
	ctx := context.Background()

	if bucket == "" {
		log.Panic(ctx, "bucket name is required")
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Panic(ctx, "could not start storage",
			log.Err(err),
		)
	}

	return &Client{
		storageClient: client,
		bucketHandler: client.Bucket(bucket),
	}
}

func (c *Client) Reader(ctx context.Context, filename string) (io.ReadCloser, error) {
	rc, err := c.bucketHandler.Object(filename).NewReader(ctx)
	if err != nil {
		return nil, err
	}

	return rc, nil
}

func (c *Client) Writer(ctx context.Context, filename string, opts ...WriterOption) (io.WriteCloser, error) {
	o := c.bucketHandler.Object(filename)
	wc := o.NewWriter(ctx)

	for _, opt := range opts {
		opt(wc)
	}

	return wc, nil
}

func (c *Client) SignedURL(filename, method string, expires time.Time) (string, error) {
	url, err := c.bucketHandler.SignedURL(filename, &storage.SignedURLOptions{
		Method:  method,
		Expires: expires,
	})
	if err != nil {
		return "", fmt.Errorf("could not get signed url: %w", err)
	}

	return url, nil
}

func (c *Client) Close() {
	if err := c.storageClient.Close(); err != nil {
		log.Error(context.Background(), "could not close storage", log.Err(err))
	}
}
