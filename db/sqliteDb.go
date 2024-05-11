package db

import (
	"database/sql"
	"guideventureapi/db/queries"
	"guideventureapi/models"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDb struct {
	file string
	db   *sql.DB
}

func NewSQLiteDb() (*SQLiteDb, error) {
	file := "sqlite.db"
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(queries.Create); err != nil {
		return nil, err
	}

	sqliteDb := SQLiteDb{file, db}
	return &sqliteDb, nil
}

func (s *SQLiteDb) CreateDummyData() error {
	_, err := s.db.Exec(queries.DummyData)
	return err
}

func (s *SQLiteDb) GetGames() ([]models.Game, error) {
	rows, err := s.db.Query(queries.SelectGames)
	if err != nil {
		return nil, err
	}

	games := []models.Game{}
	for rows.Next() {
		var game models.Game
		err = rows.Scan(&game.Id, &game.Title, &game.Introduction)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}

	return games, nil
}

func (s *SQLiteDb) GetGame(id string) (*models.Game, error) {
	row, err := s.db.Query(queries.SelectGameById, id)
	if err != nil {
		return nil, err
	}

	var game models.Game
	row.Next()
	err = row.Scan(&game.Id, &game.Title, &game.Introduction)
	if err != nil {
		return nil, err
	}

	return &game, nil
}
