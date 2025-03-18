package service

import (
	"database/sql"
	"path/filepath"

	"git.wh64.net/devproje/kuma-archive/config"
	_ "github.com/mattn/go-sqlite3"
)

func Open() (*sql.DB, error) {
	return sql.Open("sqlite3", filepath.Join(config.ROOT_DIRECTORY, "data.db"))
}
