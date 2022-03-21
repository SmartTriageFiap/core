package models

type Questions struct {
	Questions []Question `json:"questions"`
}

type Question struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	TypeAnswer  int    `json:"typeanswer"`
	Answer      string `json:"answer"`
}
