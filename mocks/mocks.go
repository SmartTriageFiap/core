package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"hmv-rest-api/database"
	"hmv-rest-api/models"
	"io/ioutil"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	CollectionQueue     *mongo.Collection
	CollectionPatient   *mongo.Collection
	CollectionQuestions *mongo.Collection
	Ctx                 = context.TODO()
)

func LoadMockDB() {

	//Database
	db := database.Connect()
	CollectionQuestions = db.Collection("questions")

	//Json Questions
	jsonFile, err := os.Open("mocks/questions.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened questions.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var questions models.Questions
	json.Unmarshal(byteValue, &questions)

	//Load data
	for i := 0; i < len(questions.Questions); i++ {
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("User Id: " + questions.Questions[i].Id)
		fmt.Println("User Description: " + questions.Questions[i].Description)
		fmt.Println("User TypeAnswer: " + strconv.Itoa(questions.Questions[i].TypeAnswer))
		fmt.Println("User Answer: " + questions.Questions[i].Answer)

		question := models.Question{
			Id:          questions.Questions[i].Id,
			Description: questions.Questions[i].Description,
			TypeAnswer:  questions.Questions[i].TypeAnswer,
			Answer:      questions.Questions[i].Answer,
		}

		result, err := CollectionQuestions.InsertOne(Ctx, question)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(result.InsertedID)
	}

}
