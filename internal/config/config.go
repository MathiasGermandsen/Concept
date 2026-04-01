package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseDSN string
	LogLevel    string
	APIKey      string
}

func Load() Config {
	_ = godotenv.Load()

	return Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseDSN: getEnv("DATABASE_DSN", "host=localhost user=postgres password=postgres dbname=mobiledisco port=5432 sslmode=disable"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		APIKey:      getEnv("API_KEY", ""),
	}
}

func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

// GetEnvAsInt reads an environment variable as an integer with a default fallback.
func GetEnvAsInt(key string, fallback int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return intValue
}
