package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	// Struct used in place of gorm.gormModel, specifies naming in json
	ID        uint            `json:"id" gorm:"primarykey"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type Game struct {
	Model
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
	Steps        []Step `json:"-"`
}

type Step struct {
	Model
	GameID      uint       `json:"game_id"`
	StepIndex   uint       `json:"step_index"`
	Type        string     `json:"step_type"`
	Points      uint       `json:"points"`
	Description *string    `json:"step_description"`
	Latitude    *float64   `json:"latitude"`
	Longitude   *float64   `json:"longitude"`
	ImageSource *string    `json:"image_source"`
	Questions   []Question `json:"-"`
}

type Question struct {
	Model
	StepID uint   `json:"step_id"`
	Type   string `json:"question_type"`
	Text   string `json:"question"`
	// answers should be formatted like "ans1;ans2;ans3" and first n answers should be the right ones
	Answers             string `json:"answers"`
	NumOfCorrectAnswers uint   `json:"num_of_correct_ans" gorm:"default:1"`
}
