package db

import "guideventureapi/db/models"

type Database interface {
	CreateDummyData() error
	GetGames() ([]models.Game, error)
	GetGame(string) (*models.Game, error)
	GetSteps(string) ([]models.Step, error)
	GetStep(string, string) (*models.Step, error)
}
