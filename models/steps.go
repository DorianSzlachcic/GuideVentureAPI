package models

type Step struct {
	Type        string `json:"step_type"`
	Points      int    `json:"points"`
	Description string `json:"step_description"`
}

type StepNavigate struct {
	Step
	Geolocation []float64 `json:"geolocation"`
}

type StepQuiz struct {
	Step
	Questions []Question `json:"questions"`
}

type StepPuzzle struct {
	Step
	ImageSource string `json:"image_source"`
}
