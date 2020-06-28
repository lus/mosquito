package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config represents the basic configuration model
type Config struct {
	Address          string
	CacheDuration    time.Duration
	DirectoryIndexes bool
}

// Current holds the current application configuration
var Current *Config

// LoadFromEnv loads the current configuration from the environment variables
func LoadFromEnv() {
	// Load the .env file
	godotenv.Load()

	// Define the server address
	address := os.Getenv("MOS_ADDRESS")
	if address == "" {
		address = ":8080"
	}

	// Parse the cache duration
	cacheDuration, err := time.ParseDuration(os.Getenv("MOS_CACHE_DURATION"))
	if err != nil {
		cacheDuration = 10 * time.Minute
	}

	// Parse the directory indexes flag
	// We don't need to handle errors here, because the fallback return value is 'false'
	directoryIndexes, _ := strconv.ParseBool(os.Getenv("MOS_DIRECTORY_INDEXES"))

	// Set the current configuration
	Current = &Config{
		Address:          os.Getenv("MOS_ADDRESS"),
		CacheDuration:    cacheDuration,
		DirectoryIndexes: directoryIndexes,
	}
}
