package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all the configuration for the application.
type Config struct {
	DatabaseDSN string
}

// Load reads the .env file and environment variables to populate the Config struct.
func Load() *Config {
	// Load .env file if it exists. We ignore the error because in production, 
	// env vars might be set directly by the hosting provider without a .env file.
	err := godotenv.Load()
	if err != nil {
		log.Println("Note: No .env file found. Falling back to system environment variables.")
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is required to start the application")
	}

	return &Config{
		DatabaseDSN: dsn,
	}
}
