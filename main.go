package main

import "github.com/Lukaesebrot/mosquito/server"

func main() {
	// Start the fasthttp server
	// TODO: Make address configurable
	server.Start(":8080")
}
