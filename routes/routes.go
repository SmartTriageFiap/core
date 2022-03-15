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
	r.HandleFunc("/v1/hmv/questions", controllers.ReturnAllQuestions).Methods("Get")
	r.HandleFunc("/v1/hmv/questions/{cpf}", controllers.SaveAnswers).Methods("Post")
	r.HandleFunc("/v1/hmv/questions/qrcode/{shortId}", controllers.ReturnQuestionsAndAnswersByShortId).Methods("Get")
	r.HandleFunc("/v1/hmv/questions/{cpf}", controllers.ReturnQuestionsAndAnswersByCpf).Methods("Get")
	r.HandleFunc("/v1/hmv/questions/{shortId}", controllers.SaveAnswersWherePatientLeftOff).Methods("Put")
	r.HandleFunc("/v1/hmv/questions/{shortId}/confirm", controllers.SavePatientData).Methods("Post")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
