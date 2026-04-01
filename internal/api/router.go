package api

import (
	"net/http"

	"DjWeb-Backend/internal/handlers"
	"DjWeb-Backend/internal/middleware"

	"github.com/gorilla/mux"
)

func NewRouter(apiKey string) http.Handler {
	router := mux.NewRouter()

	// Global middleware — applied to every request
	router.Use(middleware.RequestLogger)

	// API sub-router with authentication
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.Use(middleware.APIKeyAuth(apiKey))

	// Inquiry CRUD routes
	apiRouter.HandleFunc("/inquiries", handlers.CreateInquiry).Methods(http.MethodPost)
	apiRouter.HandleFunc("/inquiries", handlers.GetAllInquiries).Methods(http.MethodGet)
	apiRouter.HandleFunc("/inquiries/{id:[0-9]+}", handlers.GetInquiryByID).Methods(http.MethodGet)
	apiRouter.HandleFunc("/inquiries/{id:[0-9]+}", handlers.UpdateInquiry).Methods(http.MethodPut)
	apiRouter.HandleFunc("/inquiries/{id:[0-9]+}", handlers.DeleteInquiry).Methods(http.MethodDelete)

	// Health check — no auth required
	router.HandleFunc("/health", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Write([]byte(`{"status": "ok"}`))
	}).Methods(http.MethodGet)

	return router
}
