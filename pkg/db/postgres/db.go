package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	user   string
	dbname string
}

type DB struct {
	*sql.DB
}

// OpenConnection opens a connection to a database given some Config
func OpenConnection(cfg Config) (*DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=verify-full", cfg.user, cfg.dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
