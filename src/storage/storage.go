package storage

import (
	"database/sql"
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
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// TODO: https://foxcpp.dev/articles/the-right-way-to-use-go-sqlite3
	db.SetMaxOpenConns(1)

	if err = migrate(db); err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}
