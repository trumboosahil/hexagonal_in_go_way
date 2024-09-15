package postgres

import (
	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func NewPostgresConnection(connectionString string) (*sql.DB, error) {
	return sql.Open("postgres", connectionString)
}
