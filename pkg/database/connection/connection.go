package connection

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/mpsdantas/bottle/pkg/log"
)

type Config struct {
	User     string
	Password string
	Database string
	Host     string
	Port     string
	SSLMode  string
}

func New(ctx context.Context, config *Config) *sql.DB {
	var (
		sslmode = "disable"
		port    = "5432"
	)

	if config.SSLMode != "" {
		sslmode = config.SSLMode
	}

	if config.Port != "" {
		port = config.Port
	}

	conn := fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v?sslmode=%v",
		config.User,
		config.Password,
		config.Host,
		port,
		config.Database,
		sslmode,
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(ctx, "could not open postgres connection", log.Err(err))
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(ctx, "could not ping database connection", log.Err(err))
	}

	return db
}
