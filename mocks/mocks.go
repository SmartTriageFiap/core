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
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened questions.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var questions models.Questions

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &questions)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(questions.Questions); i++ {
		fmt.Println("User Id: " + questions.Questions[i].Id)
		fmt.Println("User Description: " + questions.Questions[i].Description)
		fmt.Println("User TypeAnswer: " + strconv.Itoa(questions.Questions[i].TypeAnswer))
		fmt.Println("User Answer: " + questions.Questions[i].Answer)
		fmt.Println("----------------------------------------------------------------")
	}

}
