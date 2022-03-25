package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"hmv-rest-api/controllers"
	"hmv-rest-api/database"
	"hmv-rest-api/middleware"
	"hmv-rest-api/models"
	"hmv-rest-api/services"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	CollectionQueue     *mongo.Collection
	CollectionPatients  *mongo.Collection
	CollectionQuestions *mongo.Collection
	Ctx                 = context.TODO()
	SHORTID             string
)

func SetupEnvVarsToTest() {
	os.Setenv("GO_CRIPYT", "45f84bddefa6c5212b60223ceaf64e61")
	os.Setenv("MONGO_ADDR", "localhost")
	os.Setenv("MONGO_PORT", "27017")
	os.Setenv("MONGO_USER", "admin")
	os.Setenv("MONGO_PASS", "admin")
}

func SetupHandleRequests() *mux.Router {
	routes := mux.NewRouter()
	return routes
}

func createPatientQueue() {
	db := database.Connect()
	CollectionQueue = db.Collection("queue")

	patient := models.Answers{
		Cpf:  services.Encrypt("40050060070020"),
		Salt: services.Salt("40050060070020"),
		Answers: []models.Answer{
			{Id: "1", Answer: 0},
			{Id: "2", Answer: 0},
		},
	}

	result, err := CollectionQueue.InsertOne(Ctx, patient)
	if err != nil {
		fmt.Println(err)
	}

	SHORTID = strings.ToUpper(fmt.Sprintf("%s/", result.InsertedID))
}

func deletePatientQueue() {
	db := database.Connect()
	CollectionQueue = db.Collection("queue")

	filter := bson.D{{"salt", services.Salt("40050060070020")}}

	result, err := CollectionQueue.DeleteOne(Ctx, filter)
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted", result.DeletedCount)
}

func TestShouldReturnTheQuestions(t *testing.T) {
	SetupEnvVarsToTest()
	r := SetupHandleRequests()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/v1/hmv/questions", controllers.ReturnQuestions).Methods("Get")
	req, _ := http.NewRequest("GET", "/v1/hmv/questions", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestShouldReturnStatusCodeNotFound404(t *testing.T) {
	SetupEnvVarsToTest()
	r := SetupHandleRequests()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/v1/hmv/questions/{cpf}", controllers.CheckPacient).Methods("Get")
	req, _ := http.NewRequest("GET", "/v1/hmv/questions/40050060070021", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestShouldReturnANewEnterInQueue(t *testing.T) {
	SetupEnvVarsToTest()

	createPatientQueue()
	defer deletePatientQueue()

	patientMock := models.Answers{
		Cpf:  services.Encrypt("40050060070020"),
		Salt: services.Salt("40050060070020"),
		Answers: []models.Answer{
			{Id: "1", Answer: 0},
			{Id: "2", Answer: 0},
		},
	}

	valorJson, _ := json.Marshal(patientMock)

	r := SetupHandleRequests()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/v1/hmv/questions/{cpf}", controllers.SaveAnswers).Methods("Post")

	req, _ := http.NewRequest("POST", "/v1/hmv/questions/40050060070020", bytes.NewBuffer(valorJson))
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

}
