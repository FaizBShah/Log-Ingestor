package server

import (
	"log"
	"net/http"

	"github.com/FaizBShah/log-ingestor/src/db/mongo"
	"github.com/FaizBShah/log-ingestor/src/db/postgres"
	"github.com/FaizBShah/log-ingestor/src/handlers"
	"github.com/FaizBShah/log-ingestor/src/services"
	"github.com/gorilla/mux"
)

func Run() {
	pgClient := postgres.Connect()
	mongoClient := mongo.Connect()

	logService := services.Init(pgClient, mongoClient)
	logHandler := handlers.Init(logService)

	router := mux.NewRouter()

	router.HandleFunc("/logs/ingest", logHandler.IngestLog).Methods("POST")
	router.HandleFunc("/logs", logHandler.Query).Methods("GET")

	log.Println("server started on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
