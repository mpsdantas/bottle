package database

import (
	"context"

	"github.com/mpsdantas/bottle/pkg/env"
	"github.com/mpsdantas/bottle/pkg/errors"
	"github.com/mpsdantas/bottle/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	NotFound = gorm.ErrRecordNotFound
)

func New(ctx context.Context, opt ...OptionFunc) *gorm.DB {
	opts := options{}

	for _, optionFunc := range opt {
		optionFunc(&opts)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		PreferSimpleProtocol: opts.preferSimpleProtocol,
		Conn:                 opts.conn,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(ctx, "could not start gorm db",
			log.Err(err),
		)
	}

	if len(opts.migrations) > 0 && env.Environment == env.Local {
		if err := db.AutoMigrate(opts.migrations...); err != nil {
			log.Fatal(ctx, "could not run migrations", log.Err(err))
		}
	}

	return db
}

func HandlerError(ctx context.Context, err error) error {
	if err == NotFound {
		log.Warn(ctx, "could not get entity", log.Err(err))

		return errors.NotFound("not found")
	}

	log.Error(ctx, "internal error on get entity", log.Err(err))

	return errors.Internal("internal error")
}
