package mongo

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoClient, err := mongo.Connect(nil, clientOptions)

	if err != nil {
		log.Fatalf("Failed to connect to mongodb: %v", err)
	}

	if err := mongoClient.Ping(nil, nil); err != nil {
		log.Fatalf("Failed to connect to mongodb: %v", err)
	}

	return mongoClient
}
