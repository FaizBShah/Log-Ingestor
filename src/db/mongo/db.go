package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("Failed to connect to mongodb: %v", err)
	}

	if err := mongoClient.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("Failed to connect to mongodb: %v", err)
	}

	return mongoClient
}
