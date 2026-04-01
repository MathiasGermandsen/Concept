package database

import (
	"DjWeb-Backend/internal/logger"
	"DjWeb-Backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to connect to the database")
	}

	logger.Log.Info().Msg("Database connection established")

	err = DB.AutoMigrate(&models.Inquiry{})
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to run database migrations")
	}

	logger.Log.Info().Msg("Database migrations completed")
}
