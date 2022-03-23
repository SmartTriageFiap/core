package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"hmv-rest-api/database"
	"hmv-rest-api/models"
	"hmv-rest-api/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	CollectionQueue     *mongo.Collection
	CollectionPatients  *mongo.Collection
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

	var insertAnswersModel models.Answers
	insertAnswersModel.Cpf = services.Encrypt(cpf)
	insertAnswersModel.Salt = salt
	json.NewDecoder(r.Body).Decode(&insertAnswersModel.Answers)

	fmt.Println(insertAnswersModel)

	result, err := CollectionQueue.InsertOne(Ctx, insertAnswersModel)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result.InsertedID)

	var codes models.ShortId
	codes.ShortId, codes.QrCode = services.GenerateCodes(fmt.Sprintf("%s/", result.InsertedID))

	json.NewEncoder(w).Encode(codes)

}

func UpdadeAnswers(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	CollectionQueue = db.Collection("queue")

	vars := mux.Vars(r)
	cpf := vars["cpf"]
	salt := services.Salt(cpf)

	var updadeAnswersModel models.Answers
	updadeAnswersModel.Cpf = services.Encrypt(cpf)
	updadeAnswersModel.Salt = salt
	json.NewDecoder(r.Body).Decode(&updadeAnswersModel.Answers)

	fmt.Println(updadeAnswersModel)

	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{"salt", salt}}
	update := bson.D{{"$set", updadeAnswersModel}}

	var updatedDocument bson.M
	err := CollectionQueue.FindOneAndUpdate(
		Ctx,
		filter,
		update,
		opts,
	).Decode(&updatedDocument)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		log.Fatal(err)
	}
	fmt.Printf("updated document %v", updatedDocument["_id"])

	var codes models.ShortId
	codes.ShortId, codes.QrCode = services.GenerateCodes(fmt.Sprintf("%s/", updatedDocument["_id"]))

	json.NewEncoder(w).Encode(codes)

}

func SavePatientData(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	CollectionPatients = db.Collection("patients")

	vars := mux.Vars(r)
	cpf := vars["cpf"]
	salt := services.Salt(cpf)

	var savePatientData models.PatientData

	json.NewDecoder(r.Body).Decode(&savePatientData)
	savePatientData.Patient.Salt = salt
	savePatientData.Patient.Cpf = services.Encrypt(cpf)

	result, err := CollectionPatients.InsertOne(Ctx, savePatientData)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result.InsertedID)

}
