package database

import (
	"context"
	"fmt"
	"hmv-rest-api/models"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Database {
	fmt.Println("Conectando no mongodb")
	connectionURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/",
		os.Getenv("MG_USER"),
		os.Getenv("MG_PASS"),
		os.Getenv("MG_ADDR"),
		os.Getenv("MG_PORT"))

	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("hmv")

	fmt.Println("Conectado!")

	return db
}

func SaveAnswersDB(answer models.Answer) {
	fmt.Println("Save:")
	db := Connect()
	_, err := db.Collection("answers").InsertOne(context.TODO(), answer)

	if err != nil {
		log.Fatal(err)
	}

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
