package server

import (
	"github.com/FaizBShah/log-ingestor/src/db/mongo"
	"github.com/FaizBShah/log-ingestor/src/db/postgres"
)

func Run() {
	postgres.Connect()
	mongo.Connect()
}
