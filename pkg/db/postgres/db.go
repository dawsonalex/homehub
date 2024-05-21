package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Config struct {
	User   string
	Dbname string
}

type DB struct {
	*sql.DB
}

// OpenConnection opens a connection to a database given some Config
func OpenConnection(cfg Config) (*DB, error) {
	//connStr := fmt.Sprintf("user=%s dbname=%s sslmode=verify-full", cfg.User, cfg.Dbname)
	connStr := "postgres://postgres:postgres@postgres/macrod?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
