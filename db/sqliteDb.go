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

func (s *SQLiteDb) GetGame(gameId string) (*models.Game, error) {
	row := s.db.QueryRow(queries.SelectGameById, gameId)

	var game models.Game
	err := row.Scan(&game.Id, &game.Title, &game.Introduction)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

func (s *SQLiteDb) GetSteps(gameId string) ([]models.Step, error) {
	rows, err := s.db.Query(queries.SelectSteps, gameId)
	if err != nil {
		return nil, err
	}

	steps := []models.Step{}
	for rows.Next() {
		var step models.Step
		err = rows.Scan(
			&step.Id,
			&step.GameId,
			&step.StepIndex,
			&step.Type,
			&step.Points,
			&step.Description,
			&step.Geolocation[0],
			&step.Geolocation[1],
			&step.ImageSource,
		)
		if err != nil {
			return nil, err
		}
		steps = append(steps, step)
	}

	return steps, nil
}

func (s *SQLiteDb) GetStep(gameId string, stepIndex string) (*models.Step, error) {
	row := s.db.QueryRow(queries.SelectStepByIndex, gameId, stepIndex)

	var step models.Step
	err := row.Scan(
		&step.Id,
		&step.GameId,
		&step.StepIndex,
		&step.Type,
		&step.Points,
		&step.Description,
		&step.Geolocation[0],
		&step.Geolocation[1],
		&step.ImageSource,
	)
	if err != nil {
		return nil, err
	}

	return &step, nil
}

func (s *SQLiteDb) GetQuestions(gameId string, stepIndex string) ([]models.Question, error) {
	step := s.db.QueryRow(queries.SelectStepId, gameId, stepIndex)

	var stepId int
	err := step.Scan(&stepId)
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(queries.SelectQuestions, stepId)
	if err != nil {
		return nil, err
	}

	questions := []models.Question{}
	for rows.Next() {
		var question models.Question
		err = rows.Scan(
			&question.Id,
			&question.StepId,
			&question.Type,
			&question.Text,
			&question.Answers,
			&question.NumOfCorrectAnswers,
		)
		if err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	return questions, nil
}
