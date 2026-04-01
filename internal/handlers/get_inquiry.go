package handlers

import (
	"net/http"

	"DjWeb-Backend/internal/database"
	"DjWeb-Backend/internal/logger"
	"DjWeb-Backend/internal/models"
)

func GetAllInquiries(responseWriter http.ResponseWriter, request *http.Request) {
	var inquiries []models.Inquiry

	result := database.DB.Find(&inquiries)
	if result.Error != nil {
		logger.Log.Error().Err(result.Error).Msg("Failed to fetch inquiries")
		respondWithError(responseWriter, http.StatusInternalServerError, "Failed to fetch inquiries")
		return
	}

	respondWithJSON(responseWriter, http.StatusOK, inquiries)
}

func GetInquiryByID(responseWriter http.ResponseWriter, request *http.Request) {
	inquiryID, err := parseIDFromRequest(request)
	if err != nil {
		respondWithError(responseWriter, http.StatusBadRequest, "Invalid inquiry ID")
		return
	}

	var inquiry models.Inquiry
	result := database.DB.First(&inquiry, inquiryID)
	if result.Error != nil {
		logger.Log.Warn().Uint64("inquiry_id", inquiryID).Msg("Inquiry not found")
		respondWithError(responseWriter, http.StatusNotFound, "Inquiry not found")
		return
	}

	respondWithJSON(responseWriter, http.StatusOK, inquiry)
}
