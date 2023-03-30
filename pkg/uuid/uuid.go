package uuid

import (
	"strings"

	"github.com/google/uuid"
)

//go:generate mockgen -source=./uuid.go -package=uuid -destination=./uuid_mock.go
type UUID interface {
	New(prefix ...string) string
}

type client struct{}

func New() UUID {
	return &client{}
}

func (c *client) New(prefix ...string) string {
	return c.join(prefix) + uuid.New().String()
}

func (c *client) join(prefix []string) string {
	if len(prefix) == 0 {
		return ""
	}

	return strings.Join(prefix, "-") + "-"
}
