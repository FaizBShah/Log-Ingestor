package handlers

import (
	"net/http"

	"github.com/FaizBShah/log-ingestor/src/services"
)

type LogHandler struct {
	LogService *services.LogService
}

func Init(logService *services.LogService) *LogHandler {
	return &LogHandler{
		LogService: logService,
	}
}

func (lh *LogHandler) IngestLog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (lh *LogHandler) Query(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
