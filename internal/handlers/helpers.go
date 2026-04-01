package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"DjWeb-Backend/internal/logger"

	"github.com/gorilla/mux"
)

func parseIDFromRequest(request *http.Request) (uint64, error) {
	routeVars := mux.Vars(request)
	return strconv.ParseUint(routeVars["id"], 10, 64)
}

func respondWithJSON(responseWriter http.ResponseWriter, statusCode int, payload interface{}) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(statusCode)
	if err := json.NewEncoder(responseWriter).Encode(payload); err != nil {
		logger.Log.Error().Err(err).Msg("Failed to encode JSON response")
	}
}

func respondWithError(responseWriter http.ResponseWriter, statusCode int, message string) {
	respondWithJSON(responseWriter, statusCode, map[string]string{"error": message})
}
