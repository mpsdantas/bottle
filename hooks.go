package bottle

import (
	"context"

	"github.com/mpsdantas/bottle/pkg/log"
)

func OnListen(ctx context.Context) func() error {
	return func() error {
		log.Info(ctx, "server is running")
		return nil
	}
}

func OnShutdown(ctx context.Context) func() error {
	return func() error {
		_ = log.Sync()
		log.Info(ctx, "server was stopped")

		return nil
	}
}
