package uuid

import (
	"strings"

	"github.com/google/uuid"
)

type Client struct{}

func New() *Client {
	return &Client{}
}

func (c *Client) Generate(prefix ...string) string {
	return c.join(prefix) + uuid.New().String()
}

func (c *Client) join(prefix []string) string {
	if len(prefix) == 0 {
		return ""
	}

	return strings.Join(prefix, "-") + "-"
}
