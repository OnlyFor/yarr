package storage

import (
	"database/sql"
	sqlite3 "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error) {

	sql.Register("sqlite3_with_extensions",
		&sqlite3.SQLiteDriver{
			Extensions: []string{
				"/usr/local/lib/zstd_vfs.so",
			},
		})

	db, err := sql.Open("sqlite3_with_extensions", ":memory:")
	if err != nil {
		return nil, err
	}
    _, err = db.Exec(fmt.Sprintf("ATTACH DATABASE '%s' AS loaded_db", path))
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
