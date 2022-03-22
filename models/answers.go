package models

type Answers struct {
	Cpf     string   `json:"cpf"`
	Salt    string   `json:"salt"`
	Answers []Answer `json:"answers"`
}

type Answer struct {
	Id     string `json:"id"`
	Answer int    `json:"answer"`
}
