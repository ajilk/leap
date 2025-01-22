package data

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "leap/logger"

	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	path := filepath.Join(home, ".leap", "data.db")
	err = os.MkdirAll(filepath.Dir(path), 0755) // Ensure the directory exists
	if err != nil {
		return err
	}

	db, err = sql.Open("sqlite3", path)
	if err != nil {
		return err
	}

	log.Debug("Successfully connected to the database!", "path", path)

	return db.Ping()
}
