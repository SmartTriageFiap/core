package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Database {
	fmt.Println("Conectando no mongodb")
	connectionURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/",
		os.Getenv("MG_USER"),
		os.Getenv("MG_PASS"),
		os.Getenv("MG_ADDR"),
		os.Getenv("MG_PORT"))

	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("hmv")

	fmt.Println("Conectado!")

	return db
}
