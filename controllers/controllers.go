package controllers

import (
	"encoding/json"
	"hmv-rest-api/database"
	"hmv-rest-api/models"
	"hmv-rest-api/services"
	"net/http"

	"github.com/gorilla/mux"
)

// GET /v1/hmv/questions
func ReturnAllQuestions(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.ReturnAllQuestionsDB())
}

// POST /v1/hmv/questions/{cpf}
func SaveAnswers(w http.ResponseWriter, r *http.Request) {
	var answers models.Answer

	// encripyt cpf
	cpf := services.Encrypt(mux.Vars(r)["cpf"])

	// armazena as respostas
	json.NewDecoder(r.Body).Decode(&answers)
	database.SaveAnswersDB(answers)

	// Criar o short_id/qr_code
	var codes models.ShortId
	codes.ShortId, codes.QrCode = services.GenerateCodes(cpf)

	// https://codebeautify.org/base64-to-image-converter conferir qrcode
	json.NewEncoder(w).Encode(codes)
}

// GET /v1/hmv/questions/qrcode/{shortId}
func ReturnQuestionsAndAnswersByShortId(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.ReturnQuestionsAndAnswersByShortIdDB(mux.Vars(r)["shortId"]))
}

// GET /v1/hmv/questions=cpf=123
func ReturnQuestionsAndAnswersByCpf(w http.ResponseWriter, r *http.Request) {
	cpf := services.Encrypt(mux.Vars(r)["cpf"])
	json.NewEncoder(w).Encode(database.ReturnQuestionsAndAnswersByCpfDB(cpf))
}

// PUT /v1/hmv/questions/{shortId}
func SaveAnswersWherePatientLeftOff(w http.ResponseWriter, r *http.Request) {
	var answers models.Answer
	json.NewDecoder(r.Body).Decode(&answers)
	database.UpdateAnswersDB(answers)
}

// POST /v1/hmv/questions/{shortId}/confirm
func SavePatientData(w http.ResponseWriter, r *http.Request) {
	var patientData models.PatientData
	database.SavePatientDataDB(patientData)
}
