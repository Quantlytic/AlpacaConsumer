package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	AlpacaAPIKey    string
	AlpacaSecretKey string
	AlpacaBaseURL   string
}

// Load loads the configuration from environment variables
func Load() *Config {
	// load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		AlpacaAPIKey:    getEnv("ALPACA_API_KEY", ""),
		AlpacaSecretKey: getEnv("ALPACA_SECRET_KEY", ""),
		AlpacaBaseURL:   getEnv("ALPACA_BASE_URL", "https://paper-api.alpaca.markets/v2"),
	}
}

// getEnv retrieves an environment variable or returns a fallback value
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
