package connection

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/mpsdantas/bottle/pkg/core/log"
)

type DialFunc func(conn *Connection, pgxCfg *pgx.ConnConfig) ConnDial

type Connection struct {
	User           string
	Password       string
	Database       string
	Host           string
	Port           string
	SSLMode        string
	ConnectionName string
	DialFunc       DialFunc
}

func New(opts ...Option) *sql.DB {
	ctx := context.Background()
	cfg := &Connection{
		Host:    "localhost",
		Port:    "5432",
		SSLMode: "disable",
	}

	for _, opt := range opts {
		opt(cfg)
	}

	dsn := buildDSN(cfg)
	pgxCfg, err := pgx.ParseConfig(dsn)
	if err != nil {
		log.Fatal(ctx, "could not parse pgx config", log.Err(err))
	}

	if cfg.DialFunc == nil {
		cfg.DialFunc = DialLocal
	}

	pgxCfg.DialFunc = pgconn.DialFunc(cfg.DialFunc(cfg, pgxCfg))

	dbURI := stdlib.RegisterConnConfig(pgxCfg)
	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		log.Fatal(ctx, "could not open db", log.Err(err))
	}

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(ctx, "could not ping db", log.Err(err))
	}

	return db
}

func buildDSN(cfg *Connection) string {
	var parts []string

	if cfg.User != "" {
		parts = append(parts, fmt.Sprintf("user=%s", cfg.User))
	}

	if cfg.Password != "" {
		parts = append(parts, fmt.Sprintf("password=%s", cfg.Password))
	}

	if cfg.Database != "" {
		parts = append(parts, fmt.Sprintf("dbname=%s", cfg.Database))
	}

	return strings.Join(parts, " ")
}
