package database

import (
	"database/sql"
	"fmt"

	_ "github.com/tursodatabase/go-libsql"
)

func ConnectLibsql(connString string) (*sql.DB, error) {
	db, err := sql.Open("libsql", connString)
	if err != nil {
		return nil, fmt.Errorf("opening %q using libsql driver: %w", connString, err)
	}
	return db, nil
}
