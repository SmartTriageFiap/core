package routes

import (
	"hmv-rest-api/controllers"
	"hmv-rest-api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/v1/hmv/questions", controllers.ReturnQuestions).Methods("Get")
	r.HandleFunc("/v1/hmv/questions/{cpf}", controllers.CheckPacient).Methods("Get")
	r.HandleFunc("/v1/hmv/questions/{cpf}", controllers.SaveAnswers).Methods("Post")
	r.HandleFunc("/v1/hmv/questions/{cpf}", controllers.UpdadeAnswers).Methods("Put")
	r.HandleFunc("/v1/hmv/questions/{cpf}/confirm", controllers.SavePatientData).Methods("Post")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
