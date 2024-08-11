package sqlite

import (
	"guideventureapi/db/models"

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

	err = db.AutoMigrate(&models.Game{}, &models.Step{}, &models.Question{})
	if err != nil {
		return nil, err
	}

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
	var games []models.Game
	err := s.db.Find(&games).Error
	return games, err
}

func (s *SQLiteDb) GetGame(gameId string) (*models.Game, error) {
	var game models.Game
	err := s.db.First(&game, gameId).Error
	return &game, err
}

func (s *SQLiteDb) GetSteps(gameId string) ([]models.Step, error) {
	var steps []models.Step
	err := s.db.Find(&steps, "game_id = ?", gameId).Error
	return steps, err
}

func (s *SQLiteDb) GetStep(gameId string, stepIndex string) (*models.Step, error) {
	var step models.Step
	err := s.db.First(&step, "step_index = ? AND game_id = ?", stepIndex, gameId).Error
	if err != nil {
		return nil, err
	}
	err = s.db.Find(&step.Questions, "step_id = ?", step.ID).Error
	return &step, err
}
