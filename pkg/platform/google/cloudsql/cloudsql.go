package cloudsql

import (
	"context"
	"fmt"
	"net"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v5"
	"github.com/mpsdantas/bottle/pkg/core/log"
	"github.com/mpsdantas/bottle/pkg/platform/postgresql/connection"
)

func PostgresqlDialer(conn *connection.Connection, _ *pgx.ConnConfig) connection.ConnDial {
	lctx := context.Background()

	dialer, err := cloudsqlconn.NewDialer(lctx)
	if err != nil {
		log.Fatal(lctx, "could not create cloud sql dialer", log.Err(err))
	}

	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		if conn.ConnectionName == "" {
			return nil, fmt.Errorf("could not start postgresql, no connection name")
		}

		return dialer.Dial(ctx, conn.ConnectionName)
	}
}
