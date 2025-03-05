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

func InsertBookmark(key string, value string) something.Bookmark {
	ctx := context.Background()
	queries := something.New(db)

	bookmark, err := queries.CreateBookmark(ctx, something.CreateBookmarkParams{
		Key:   key,
		Value: value,
	})

	if err != nil {
		panic(err)
	}

	return bookmark
}

func ListBookmarks() []something.Bookmark {
	ctx := context.Background()
	queries := something.New(db)

	bookmarks, err := queries.ListBookmarks(ctx)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	return bookmarks
}

func DeleteBookmark(id int) {
	ctx := context.Background()
	queries := something.New(db)

	err := queries.DeleteBookmark(ctx, int64(id))
	if err != nil {
		log.Error(err)
		panic(err)
	}
}

func DeleteAllBookmarks() {
	ctx := context.Background()
	queries := something.New(db)

	err := queries.DeleteAllBookmarks(ctx)
	if err != nil {
		log.Error(err)
		panic(err)
	}
}
