package data

import (
	"context"
	"database/sql"
	_ "embed"
	"os"
	"path/filepath"

	_ "leap/logger"

	"leap/data/something"

	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

//go:embed schema.sql
var ddl string

func OpenDatabase() error {
	home, err := os.UserHomeDir()
	if err != nil {
    log.Error(err)
		return err
	}

	path := filepath.Join(home, ".leap", "data.db")
	err = os.MkdirAll(filepath.Dir(path), 0755) // Ensure the directory exists
	if err != nil {
    log.Error(err)
		return err
	}

	db, err = sql.Open("sqlite3", path)
	if err != nil {
    log.Error(err)
		return err
	}

	log.Debug("Successfully connected to the database!", "path", path)

	return db.Ping()
}

func CreateBookmarksTable() error {
	ctx := context.Background()

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
    log.Error(err)
		return err
	}

	log.Debug("Table Created")
	return nil
}

func InsertAndFetch() error {
	ctx := context.Background()
	queries := something.New(db)

	bookmark, err := queries.CreateBookmark(ctx, something.CreateBookmarkParams{
		Key:   "Key Here",
		Value: sql.NullString{String: "Value Here", Valid: true},
	})

  if err != nil {
    log.Error(err)
    return err
  }


	bookmarks, err := queries.ListBookmarks(ctx)
	if err != nil {
    log.Error(err)
		return err
	}

	log.Debug(bookmarks)
	log.Debug(bookmark)
	return nil
}
