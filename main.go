package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"DjWeb-Backend/internal/api"
	"DjWeb-Backend/internal/config"
	"DjWeb-Backend/internal/database"
	"DjWeb-Backend/internal/logger"
)

func main() {
	appConfig := config.Load()

	logger.Init(appConfig.LogLevel)

	database.Connect(appConfig.DatabaseDSN)

	router := api.NewRouter(appConfig.APIKey)

	server := &http.Server{
		Addr:         ":" + appConfig.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine so we can handle graceful shutdown
	go func() {
		logger.Log.Info().Str("port", appConfig.Port).Msg("Server is starting")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal().Err(err).Msg("Server failed to start")
		}
	}()

	// Wait for interrupt signal for graceful shutdown
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM)
	<-shutdownSignal

	logger.Log.Info().Msg("Shutting down server...")

	shutdownContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownContext); err != nil {
		logger.Log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	logger.Log.Info().Msg("Server stopped gracefully")
}
