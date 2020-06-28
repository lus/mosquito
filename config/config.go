package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config represents the basic configuration model
type Config struct {
	Address string
}

// Current holds the current application configuration
var Current *Config

// LoadFromEnv loads the current configuration from the environment variables
func LoadFromEnv() {
	// Load the .env file
	godotenv.Load()

	// Set the current configuration
	Current = &Config{
		Address: os.Getenv("MOS_ADDRESS"),
	}
}
