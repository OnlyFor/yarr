package storage

import (
	"database/sql"
	"log"
	"strings"
	// sqlite3 "github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error) {
	// sql.Register("sqlite3_with_extensions",
	// 	&sqlite3.SQLiteDriver{
	// 		Extensions: []string{
	// 			"/usr/local/lib/libsqlite_zstd.so",
	// 		},
	// 	})
	// db, err := sql.Open("sqlite3_with_extensions", path)
	if pos := strings.IndexRune(path, '?'); pos == -1 {
		params := "_journal=WAL&_sync=NORMAL&_busy_timeout=5000&cache=shared"
		log.Printf("opening db with params: %s", params)
		path = path + "?" + params
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// _, err = db.Exec("PRAGMA auto_vacuum = full;")
	// if err != nil {
	// 	return nil, err
	// }
	// _, err = db.Exec("PRAGMA journal_mode = MEMORY;")
	// if err != nil {
	// 	return nil, err
	// }

	// TODO: https://foxcpp.dev/articles/the-right-way-to-use-go-sqlite3
	db.SetMaxOpenConns(1)

	if err = migrate(db); err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}
