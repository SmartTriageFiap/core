package main

import (
	"fmt"
	"hmv-rest-api/mocks"
	"hmv-rest-api/routes"

	"go.mongodb.org/mongo-driver/mongo"
)

var BooksCollection *mongo.Collection

func main() {
	fmt.Println("Iniciando API HMV")
	mocks.LoadMockDB()
	routes.HandleRequest()
}
