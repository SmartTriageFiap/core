package models

type Answer struct {
	Cpf       string     `json:cpf`
	Questions []Question `json:questions`
}
