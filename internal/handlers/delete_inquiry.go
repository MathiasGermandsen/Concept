package handlers

import (
	"net/http"

	"DjWeb-Backend/internal/database"
	"DjWeb-Backend/internal/logger"
	"DjWeb-Backend/internal/models"
)

func DeleteInquiry(responseWriter http.ResponseWriter, request *http.Request) {
	inquiryID, err := parseIDFromRequest(request)
	if err != nil {
		respondWithError(responseWriter, http.StatusBadRequest, "Invalid inquiry ID")
		return
	}

	result := database.DB.Delete(&models.Inquiry{}, inquiryID)
	if result.Error != nil {
		logger.Log.Error().Err(result.Error).Msg("Failed to delete inquiry")
		respondWithError(responseWriter, http.StatusInternalServerError, "Failed to delete inquiry")
		return
	}

	if result.RowsAffected == 0 {
		respondWithError(responseWriter, http.StatusNotFound, "Inquiry not found")
		return
	}

	logger.Log.Info().Uint64("inquiry_id", inquiryID).Msg("Inquiry deleted")
	respondWithJSON(responseWriter, http.StatusOK, map[string]string{"message": "Inquiry deleted successfully"})
}
