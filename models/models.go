package models

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
	Steps        []Step
}

type Step struct {
	gorm.Model
	GameID      uint     `json:"game_id"`
	StepIndex   uint     `json:"step_index"`
	Type        string   `json:"step_type"`
	Points      uint     `json:"points"`
	Description *string  `json:"step_description"`
	Latitude    *float64 `json:"latitude"`
	Longitude   *float64 `json:"longitude"`
	ImageSource *string  `json:"image_source"`
	Questions   []Question
}

type Question struct {
	gorm.Model
	StepID uint   `json:"step_id"`
	Type   string `json:"question_type"`
	Text   string `json:"question"`
	// answers should be formatted like "ans1;ans2;ans3" and first n answers should be the right ones
	Answers             string `json:"answers"`
	NumOfCorrectAnswers uint   `json:"num_of_correct_ans" gorm:"default:1"`
}
