package services

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type LogService struct {
	PgClient    *gorm.DB
	MongoClient *mongo.Client
}

func Init(pgClient *gorm.DB, mongoClient *mongo.Client) *LogService {
	return &LogService{
		PgClient:    pgClient,
		MongoClient: mongoClient,
	}
}
