package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func New(storagePath string) (*sql.DB, error) {
	const op = "storage.New"
	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return db, nil
}
