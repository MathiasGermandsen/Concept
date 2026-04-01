package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"DjWeb-Backend/internal/database"
	"DjWeb-Backend/internal/logger"
	"DjWeb-Backend/internal/models"

	"github.com/gorilla/mux"
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

// --- shared helpers (DRY) ---

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
