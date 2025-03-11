package database

import (
	"database/sql"
	"fmt"

	"github.com/iqthuc/sport-shop/config"
)

func NewPostgres() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.DB_SOURCE)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to postgres database: %w", err)
	}
	return db, nil
}
