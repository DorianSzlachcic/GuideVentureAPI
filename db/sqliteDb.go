package db

import (
	"database/sql"
	"guideventureapi/db/queries"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDb struct {
	file string
	db   *sql.DB
}

func NewDb() (*SQLiteDb, error) {
	file := "sqlite.db"
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(queries.CREATE); err != nil {
		return nil, err
	}

	sqliteDb := SQLiteDb{file, db}
	return &sqliteDb, nil
}
