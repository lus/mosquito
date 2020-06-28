package server

import (
	"github.com/Lukaesebrot/mosquito/config"
	routing "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// Start starts the server
func Start(address string) {
	// Create a new router
	router := routing.New()
	router.ServeFilesCustom("/{filepath:*}", &fasthttp.FS{
		Root:               "./data",
		GenerateIndexPages: config.Current.DirectoryIndexes,
		CacheDuration:      config.Current.CacheDuration,
	})

	// Start the fasthttp server
	fasthttp.ListenAndServe(address, router.Handler)
}
