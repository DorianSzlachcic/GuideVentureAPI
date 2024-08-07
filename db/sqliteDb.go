package db

import (
	"guideventureapi/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ptr[T any](val T) *T {
	return &val
}

type SQLiteDb struct {
	db *gorm.DB
}

func NewSQLiteDb() (*SQLiteDb, error) {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Game{}, &models.Step{}, &models.Question{})

	sqliteDb := SQLiteDb{db}
	return &sqliteDb, nil
}

func (s *SQLiteDb) CreateDummyData() error {
	questions := []models.Question{{Type: "single_choice", Text: "Gdzie się znajdujesz?", Answers: "Wrocław;Gdańsk;Poznań;Inne"},
		{Type: "text_answer", Text: "Jaki jest najlepszy uniwersytet we Wrocławiu?", Answers: "UWR"},
		{Type: "multiple_choice", Text: "Jakie są największe atrakcje Wrocławia?", Answers: "Rynek;WFiA;Wyspa Słodowa;Nadodrze", NumOfCorrectAnswers: 2}}
	steps := []models.Step{{StepIndex: 1, Type: "navigate", Points: 50, Description: ptr("Przejdź do punktu startowego"), Latitude: ptr(51.109730), Longitude: ptr(17.030655)},
		{StepIndex: 2, Type: "quiz", Points: 300, Description: ptr("Odpowiedz na poniższe pytania"), Questions: questions},
		{StepIndex: 3, Type: "navigate", Points: 50, Description: ptr("Przejdź do następnego punktu"), Latitude: ptr(51.114730), Longitude: ptr(17.042187)},
		{StepIndex: 4, Type: "photo", Points: 100, Description: ptr("Wykonaj sobie zdjęcie")},
		{StepIndex: 5, Type: "navigate", Points: 50, Description: ptr("Przejdź do następnego punktu"), Latitude: ptr(51.114046), Longitude: ptr(17.031246)},
		{StepIndex: 6, Type: "puzzle", Points: 300, Description: ptr("Ułóżcie puzzle"), ImageSource: ptr("https://upload.wikimedia.org/wikipedia/commons/thumb/b/b6/Wroclaw_-_Uniwersytet_Wroclawski_o_poranku.jpg/1200px-Wroclaw_-_Uniwersytet_Wroclawski_o_poranku.jpg")}}
	game := models.Game{Title: "Przykładowa gra", Introduction: "Gra przykładowa stworzona na potrzebny developmentu", Steps: steps}
	return s.db.Create(&game).Error
}

func (s *SQLiteDb) GetGames() ([]models.Game, error) {
	// rows, err := s.db.Query(queries.SelectGames)
	// if err != nil {
	// 	return nil, err
	// }

	// games := []models.Game{}
	// for rows.Next() {
	// 	var game models.Game
	// 	err = rows.Scan(&game.Id, &game.Title, &game.Introduction)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	games = append(games, game)
	// }

	// return games, nil
	return nil, nil
}

func (s *SQLiteDb) GetGame(gameId string) (*models.Game, error) {
	// row := s.db.QueryRow(queries.SelectGameById, gameId)

	// var game models.Game
	// err := row.Scan(&game.Id, &game.Title, &game.Introduction)
	// if err != nil {
	// 	return nil, err
	// }

	// return &game, nil
	return nil, nil
}

func (s *SQLiteDb) GetSteps(gameId string) ([]models.Step, error) {
	// rows, err := s.db.Query(queries.SelectSteps, gameId)
	// if err != nil {
	// 	return nil, err
	// }

	// steps := []models.Step{}
	// for rows.Next() {
	// 	var step models.Step
	// 	err = rows.Scan(
	// 		&step.Id,
	// 		&step.GameId,
	// 		&step.StepIndex,
	// 		&step.Type,
	// 		&step.Points,
	// 		&step.Description,
	// 		&step.Geolocation[0],
	// 		&step.Geolocation[1],
	// 		&step.ImageSource,
	// 	)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	steps = append(steps, step)
	// }

	// return steps, nil
	return nil, nil
}

func (s *SQLiteDb) GetStep(gameId string, stepIndex string) (*models.Step, error) {
	// row := s.db.QueryRow(queries.SelectStepByIndex, gameId, stepIndex)

	// var step models.Step
	// err := row.Scan(
	// 	&step.Id,
	// 	&step.GameId,
	// 	&step.StepIndex,
	// 	&step.Type,
	// 	&step.Points,
	// 	&step.Description,
	// 	&step.Geolocation[0],
	// 	&step.Geolocation[1],
	// 	&step.ImageSource,
	// )
	// if err != nil {
	// 	return nil, err
	// }

	// return &step, nil
	return nil, nil
}

func (s *SQLiteDb) GetQuestions(gameId string, stepIndex string) ([]models.Question, error) {
	// step := s.db.QueryRow(queries.SelectStepId, gameId, stepIndex)

	// var stepId int
	// err := step.Scan(&stepId)
	// if err != nil {
	// 	return nil, err
	// }

	// rows, err := s.db.Query(queries.SelectQuestions, stepId)
	// if err != nil {
	// 	return nil, err
	// }

	// questions := []models.Question{}
	// for rows.Next() {
	// 	var question models.Question
	// 	err = rows.Scan(
	// 		&question.Id,
	// 		&question.StepId,
	// 		&question.Type,
	// 		&question.Text,
	// 		&question.Answers,
	// 		&question.NumOfCorrectAnswers,
	// 	)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	questions = append(questions, question)
	// }

	// return questions, nil
	return nil, nil
}
