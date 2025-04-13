package idtoken

import (
	"context"
	"net/http"

	"github.com/mpsdantas/bottle/pkg/core/log"
	"google.golang.org/api/idtoken"
)

func New(base string) *http.Client {
	ctx := context.Background()

	client, err := idtoken.NewClient(ctx, base)
	if err != nil {
		log.Fatal(ctx, "could not get client", log.Err(err))
	}

	return client
}
