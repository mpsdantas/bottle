package database

import (
	"context"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mpsdantas/bottle/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Mock(ctx context.Context) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(ctx, "could not start database mock",
			log.Err(err),
		)
	}

	open, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       db,
	}))
	if err != nil {
		log.Fatal(ctx, "could not connect database mock", log.Err(err))
	}

	return open, mock
}
