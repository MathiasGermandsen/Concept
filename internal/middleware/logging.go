package middleware

import (
	"net/http"
	"time"

	"DjWeb-Backend/internal/logger"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		startTime := time.Now()

		logger.Log.Info().
			Str("method", request.Method).
			Str("path", request.URL.Path).
			Msg("Endpoint hit")

		next.ServeHTTP(responseWriter, request)

		duration := time.Since(startTime)

		logger.Log.Info().
			Str("method", request.Method).
			Str("path", request.URL.Path).
			Dur("duration", duration).
			Msg("Request completed")
	})
}
