package postgresql

import (
	"database/sql"
)

type options struct {
	db                   *sql.DB
	preferSimpleProtocol bool
}

type OptionFunc = func(option *options)

func WithConn(db *sql.DB) OptionFunc {
	return func(option *options) {
		option.db = db
	}
}

func WithPreferSimpleProtocol(enable bool) OptionFunc {
	return func(option *options) {
		option.preferSimpleProtocol = enable
	}
}
