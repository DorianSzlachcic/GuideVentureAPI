package models

type Game struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
	Steps        []Step `json:"steps"`
}
