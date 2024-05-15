package models

import "database/sql"

type Game struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
}

type Step struct {
	Id          int                `json:"id"`
	GameId      int                `json:"game_id"`
	StepIndex   int                `json:"step_index"`
	Type        string             `json:"step_type"`
	Points      int                `json:"points"`
	Description sql.NullString     `json:"step_description"`
	Geolocation [2]sql.NullFloat64 `json:"geolocation"`
	ImageSource sql.NullString     `json:"image_source"`
}

type Question struct {
	Id                  int    `json:"id"`
	StepId              int    `json:"step_id"`
	Type                string `json:"question_type"`
	Text                string `json:"question"`
	Answers             string `json:"answers"`
	NumOfCorrectAnswers int    `json:"num_of_correct_ans"`
}
