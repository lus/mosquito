package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config represents the basic configuration model
type Config struct {
	Address       string
	CacheDuration time.Duration
}

// Current holds the current application configuration
var Current *Config

// LoadFromEnv loads the current configuration from the environment variables
func LoadFromEnv() {
	// Load the .env file
	godotenv.Load()

	// Parse the cache duration
	cacheDuration, err := time.ParseDuration(os.Getenv("MOS_CACHE_DURATION"))
	if err != nil {
		cacheDuration = 10 * time.Minute
	}

	// Set the current configuration
	Current = &Config{
		Address:       os.Getenv("MOS_ADDRESS"),
		CacheDuration: cacheDuration,
	}
}
