package service

import (
	"database/sql"

	"git.wh64.net/devproje/kuma-archive/config"
	_ "github.com/mattn/go-sqlite3"
)

func Open() (*sql.DB, error) {
	return sql.Open("sqlite3", (config.ROOT_DIRECTORY))
}
