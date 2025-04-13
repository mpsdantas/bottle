package connection

import (
	"context"
	"net"

	"github.com/jackc/pgx/v5"
)

type ConnDial func(ctx context.Context, network, addr string) (net.Conn, error)

func DialLocal(conn *Connection, pgxCfg *pgx.ConnConfig) ConnDial {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		pgxCfg.Config.RuntimeParams["sslmode"] = conn.SSLMode
		return net.Dial("tcp", net.JoinHostPort(conn.Host, conn.Port))
	}
}
