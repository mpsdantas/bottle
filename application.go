package bottle

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/mpsdantas/bottle/pkg/env"
	"github.com/mpsdantas/bottle/pkg/log"
)

type Application struct {
	*fiber.App
}

func New() *Application {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          ErrorHandler,
	})
	app.Use(recover.New())
	app.Use(requestid.New())

	return &Application{
		app,
	}
}

func (a *Application) Run() {
	ctx := context.Background()

	log.Info(ctx, "starting server",
		log.String("name", env.Application),
		log.String("environment", env.Environment),
		log.String("scope", env.Scope),
	)

	a.App.Hooks().OnListen(OnListen(ctx))
	a.App.Hooks().OnShutdown(OnShutdown(ctx))

	errChan := make(chan error)
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := a.start(); err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		a.stop(ctx)

		log.Fatal(ctx, "could not start server",
			log.Err(err),
		)
	case <-stopChan:
		a.stop(ctx)
	}
}

func (a *Application) start() error {
	addr := fmt.Sprintf(":%v", env.Port)
	if err := a.Listen(addr); err != nil {
		return err
	}

	return nil
}

func (a *Application) stop(ctx context.Context) {
	log.Info(ctx, "stopping server")

	if err := a.Shutdown(); err != nil {
		log.Fatal(ctx, "could not shutdown server",
			log.Err(err),
		)
	}
}
