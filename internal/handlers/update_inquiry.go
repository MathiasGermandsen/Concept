package handlers

import (
	"encoding/json"
	"net/http"

	"DjWeb-Backend/internal/database"
	"DjWeb-Backend/internal/logger"
	"DjWeb-Backend/internal/models"
)

func UpdateInquiry(responseWriter http.ResponseWriter, request *http.Request) {
	inquiryID, err := parseIDFromRequest(request)
	if err != nil {
		respondWithError(responseWriter, http.StatusBadRequest, "Invalid inquiry ID")
		return
	}

	var existingInquiry models.Inquiry
	result := database.DB.First(&existingInquiry, inquiryID)
	if result.Error != nil {
		logger.Log.Warn().Uint64("inquiry_id", inquiryID).Msg("Inquiry not found for update")
		respondWithError(responseWriter, http.StatusNotFound, "Inquiry not found")
		return
	}

	var updatedFields models.Inquiry
	if err := json.NewDecoder(request.Body).Decode(&updatedFields); err != nil {
		logger.Log.Error().Err(err).Msg("Failed to decode update payload")
		respondWithError(responseWriter, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer request.Body.Close()

	result = database.DB.Model(&existingInquiry).Updates(updatedFields)
	if result.Error != nil {
		logger.Log.Error().Err(result.Error).Msg("Failed to update inquiry")
		respondWithError(responseWriter, http.StatusInternalServerError, "Failed to update inquiry")
		return
	}

	logger.Log.Info().Uint64("inquiry_id", inquiryID).Msg("Inquiry updated")
	respondWithJSON(responseWriter, http.StatusOK, existingInquiry)
}
