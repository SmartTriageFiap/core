package controllers

import (
	"context"
	"encoding/json"
	"hmv-rest-api/database"
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
	//vars := mux.Vars(r)
	//cpf := vars["cpf"]
	//salt := services.Salt(cpf)

}

func UpdadeAnswers(w http.ResponseWriter, r *http.Request) {

}

func SavePatientData(w http.ResponseWriter, r *http.Request) {

}
