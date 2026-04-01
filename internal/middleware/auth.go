package middleware

import (
	"crypto/subtle"
	"net/http"

	"DjWeb-Backend/internal/logger"
)

func APIKeyAuth(validAPIKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
			providedKey := request.Header.Get("X-API-Key")

			if validAPIKey == "" {
				logger.Log.Warn().Msg("API_KEY is not configured — all requests are allowed")
				next.ServeHTTP(responseWriter, request)
				return
			}

			if subtle.ConstantTimeCompare([]byte(providedKey), []byte(validAPIKey)) != 1 {
				logger.Log.Warn().
					Str("method", request.Method).
					Str("path", request.URL.Path).
					Msg("Unauthorized request — invalid or missing API key")

				http.Error(responseWriter, `{"error": "unauthorized"}`, http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(responseWriter, request)
		})
	}
}
