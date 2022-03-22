package main

import (
	"fmt"
	"hmv-rest-api/mocks"
	"hmv-rest-api/routes"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

var BooksCollection *mongo.Collection

func main() {
	fmt.Println("Iniciando API HMV")
	os.Setenv("GO_CRIPYT", "45f84bddefa6c5212b60223ceaf64e61")

	os.Setenv("MG_USER", "admin")
	os.Setenv("MG_PASS", "admin")
	os.Setenv("MG_ADDR", "localhost")
	os.Setenv("MG_PORT", "27017")

	mocks.LoadMockDB()

	routes.HandleRequest()
}

// export GO111MODULE=on
// make docker-stop && make docker-mongo && go run main.go
