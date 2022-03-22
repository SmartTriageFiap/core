package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"hmv-rest-api/database"
	"hmv-rest-api/models"
	"hmv-rest-api/services"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	CollectionQueue     *mongo.Collection
	CollectionPatient   *mongo.Collection
	CollectionQuestions *mongo.Collection
	Ctx                 = context.TODO()
)

func CheckPacient(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	CollectionQueue = db.Collection("queue")

	vars := mux.Vars(r)
	cpf := vars["cpf"]
	salt := services.Salt(cpf)

	var result bson.M
	err := CollectionQueue.FindOne(Ctx, bson.D{{"salt", salt}}).Decode(&result)
	if err != nil || err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(result)
}

func ReturnQuestions(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	CollectionQueue = db.Collection("questions")

	cursor, err := CollectionQueue.Find(Ctx, bson.D{})
	if err != nil {
		panic(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(results)
}

func SaveAnswers(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	CollectionQueue = db.Collection("queue")

	vars := mux.Vars(r)
	cpf := vars["cpf"]
	salt := services.Salt(cpf)

	var answers []models.Answer
	json.NewDecoder(r.Body).Decode(&answers)

	var insertAnswersModel models.Answers
	insertAnswersModel.Cpf = services.Encrypt(cpf)
	insertAnswersModel.Salt = salt
	insertAnswersModel.Answers = answers

	fmt.Println(insertAnswersModel)

	result, err := CollectionQueue.InsertOne(Ctx, insertAnswersModel)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result.InsertedID)

	var codes models.ShortId
	codes.ShortId, codes.QrCode = services.GenerateCodes(fmt.Sprintf("%s/", result.InsertedID))

	// https://codebeautify.org/base64-to-image-converter conferir qrcode
	json.NewEncoder(w).Encode(codes)

}

func UpdadeAnswers(w http.ResponseWriter, r *http.Request) {

}

func SavePatientData(w http.ResponseWriter, r *http.Request) {

}
