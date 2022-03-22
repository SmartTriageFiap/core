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
	connectionURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/",
		os.Getenv("MONGO_USER"),
		os.Getenv("MONGO_PASS"),
		os.Getenv("MONGO_ADDR"),
		os.Getenv("MONGO_PORT"))

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

	return db
}
