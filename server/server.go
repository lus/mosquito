package server

import (
	routing "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// Start starts the server
func Start(address string) {
	// Create a new router
	router := routing.New()

	// Start the fasthttp server
	fasthttp.ListenAndServe(address, router.Handler)
}
