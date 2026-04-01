package handlers

import (
	"encoding/json"
	"net/http"

	"DjWeb-Backend/internal/database"
	"DjWeb-Backend/internal/logger"
	"DjWeb-Backend/internal/models"
)

func CreateInquiry(responseWriter http.ResponseWriter, request *http.Request) {
	var inquiry models.Inquiry

	if err := json.NewDecoder(request.Body).Decode(&inquiry); err != nil {
		logger.Log.Error().Err(err).Msg("Failed to decode request body")
		respondWithError(responseWriter, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer request.Body.Close()

	if inquiry.Email == "" || inquiry.CustomerName == "" {
		respondWithError(responseWriter, http.StatusBadRequest, "Customer name and email are required")
		return
	}

	result := database.DB.Create(&inquiry)
	if result.Error != nil {
		logger.Log.Error().Err(result.Error).Msg("Failed to create inquiry")
		respondWithError(responseWriter, http.StatusInternalServerError, "Failed to create inquiry")
		return
	}

	logger.Log.Info().Uint("inquiry_id", inquiry.ID).Msg("Inquiry created")
	respondWithJSON(responseWriter, http.StatusCreated, inquiry)
}
