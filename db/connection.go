package db

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib" // pgx Postgres driver
)

func Connect(databaseURL string) (*sql.DB, error) {
	return sql.Open("pgx", databaseURL)
}
