package database

import (
	"fmt"
	"hmv-rest-api/models"
)

func SaveAnswersDB(answer models.Answer) {
	fmt.Println("Save:")
	fmt.Println(answer)
}

func InitServiceDB(cpf string) {
	fmt.Println("Save:", cpf)
}

func UpdateAnswersDB(answer models.Answer) {
	fmt.Println("Update:")
	fmt.Println(answer)
}

func ReturnAllQuestionsDB() []models.Question {
	questions := []models.Question{
		{Id: "1", Description: "oi", TypeAnswer: 1},
		{Id: "2", Description: "oi2", TypeAnswer: 0},
		{Id: "3", Description: "oi3", TypeAnswer: 0},
	}
	return questions
}

func ReturnQuestionsAndAnswersByShortIdDB(shortId string) models.Answer {
	var answers models.Answer
	return answers
}

func ReturnQuestionsAndAnswersByCpfDB(cpf string) models.Answer {
	var answers models.Answer
	return answers
}

func SavePatientDataDB(patientData models.PatientData) {
	fmt.Println("Save data:")
	fmt.Println(patientData)
}

func GetShortId(cpf string) string {
	return cpf
}
