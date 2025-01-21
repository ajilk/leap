package data

import (
	"database/sql"

	_ "leap/logger"

	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {

	var err error

	db, err = sql.Open("sqlite3", "./.data.db")
	if err != nil {
		return err
	}

	log.Debug("Successfully connected to the database!")

	return db.Ping()
}
