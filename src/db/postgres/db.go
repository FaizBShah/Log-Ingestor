package postgres

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=postgres user=postgres password=postgres dbname=log_ingestor port=5432 sslmode=disable TimeZone=UTC"

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatalf("Failed to connect to postgres: %v", err)
	}

	return db
}
