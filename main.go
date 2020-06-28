package main

import (
	"github.com/Lukaesebrot/mosquito/config"
	"github.com/Lukaesebrot/mosquito/server"
)

func main() {
	// Load the application configuration
	config.LoadFromEnv()

	// Start the fasthttp server
	server.Start(config.Current.Address)
}
