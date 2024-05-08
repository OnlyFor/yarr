package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error) {

	sql.Register("sqlite3_with_extensions", &SQLiteDriver{true})

	db, err := sql.Open("sqlite3_with_extensions", path)
	if err != nil {
		return nil, err
	}

	// _, err = db.Exec("select load_extension('sqlite3_mod_regexp.dll')")
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
