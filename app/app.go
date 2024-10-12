package app

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/drornir/pizi-go/app/database"
)

func NewLocalLibSQLApp(dir string) (*App, error) {
	err := os.MkdirAll(dir, os.ModeDir)
	if err != nil {
		return nil, fmt.Errorf("creating path to db %q: %w", dir, err)
	}
	db, err := database.ConnectLibsql(filepath.Join(dir, "app.db"))
	if err != nil {
		return nil, fmt.Errorf("connecting to local libsql app.db: %w", err)
	}

	// todo add more system databases like message queues, db migrations etc.

	return &App{
		db: db,
	}, nil
}

type App struct {
	db *sql.DB
}
