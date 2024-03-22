package cloudstorage

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"github.com/mpsdantas/bottle/pkg/log"
)

type Client struct {
	bucket string
	c      *storage.Client
}

func New(ctx context.Context, bucket string) *Client {
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Panic(ctx, "could not start google cloudstorage",
			log.Err(err),
		)
	}

	return &Client{
		bucket: bucket,
		c:      client,
	}
}

func (c *Client) Upload(ctx context.Context, filename string, data io.Reader) error {
	o := c.c.Bucket(c.bucket).Object(filename)

	wc := o.NewWriter(ctx)

	if _, err := io.Copy(wc, data); err != nil {
		return fmt.Errorf("could not copy file: %w", err)
	}

	if err := wc.Close(); err != nil {
		return fmt.Errorf("could not close file: %w", err)
	}

	return nil
}

func (c *Client) Download(ctx context.Context, filename string) ([]byte, error) {
	rc, err := c.c.Bucket(c.bucket).Object(filename).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rc.Close(); err != nil {
			log.Error(ctx, "could not close file",
				log.Err(err),
			)
		}
	}()

	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	return data, nil
}
