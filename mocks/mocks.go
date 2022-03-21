package mocks

import (
	"encoding/json"
	"fmt"
	"hmv-rest-api/models"
	"io/ioutil"
	"os"
	"strconv"
)

func LoadQuestions() {

	jsonFile, err := os.Open("mocks/questions.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened questions.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var questions models.Questions

	json.Unmarshal(byteValue, &questions)

	for i := 0; i < len(questions.Questions); i++ {
		fmt.Println("User Id: " + questions.Questions[i].Id)
		fmt.Println("User Description: " + questions.Questions[i].Description)
		fmt.Println("User TypeAnswer: " + strconv.Itoa(questions.Questions[i].TypeAnswer))
		fmt.Println("User Answer: " + questions.Questions[i].Answer)
		fmt.Println("----------------------------------------------------------------")
	}

}
