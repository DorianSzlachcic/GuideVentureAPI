package models

type Question struct {
	Type                string   `json:"question_type"`
	Text                string   `json:"question"`
	Answers             []string `json:"answers"`
	NumOfCorrectAnswers int      `json:"num_of_correct_ans"`
}
