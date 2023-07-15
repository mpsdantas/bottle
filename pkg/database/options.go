package database

import (
	"database/sql"
)

type options struct {
	conn                 *sql.DB
	migrations           []interface{}
	preferSimpleProtocol bool
}

type OptionFunc = func(option *options)

func WithConn(conn *sql.DB) OptionFunc {
	return func(option *options) {
		option.conn = conn
	}
}

func WithMigrations(values ...interface{}) OptionFunc {
	return func(option *options) {
		option.migrations = values
	}
}

func WithPreferSimpleProtocol(enable bool) OptionFunc {
	return func(option *options) {
		option.preferSimpleProtocol = enable
	}
}
