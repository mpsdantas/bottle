package cloudsql

import (
	"context"
	"database/sql"
	"fmt"
	"net"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/mpsdantas/bottle/pkg/database/connection"
	"github.com/mpsdantas/bottle/pkg/log"
)

func DB(ctx context.Context, cfg connection.Config) *sql.DB {
	dsn := fmt.Sprintf("user=%s password=%s database=%s", cfg.User, cfg.Password, cfg.Database)
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		log.Fatal(ctx, "could not parse config", log.Err(err))
	}

	d, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		log.Fatal(ctx, "could not create dialler", log.Err(err))
	}

	config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
		return d.Dial(ctx, cfg.ConnectionName)
	}

	dbURI := stdlib.RegisterConnConfig(config)
	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		log.Fatal(ctx, "could open sql connection", log.Err(err))
	}

	return dbPool
}
