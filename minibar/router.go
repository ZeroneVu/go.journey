package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ZeroneVu/go.journey/minibar/config"
	"github.com/ZeroneVu/go.journey/minibar/routing"
)

// NewStack ...
// A stack to wrap each http handler with the middleware provided
func NewStack(routes routing.Routes, c config.Config) *mux.Router {

	// Create the router
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		// A router is created with each handler defined in the routes array
		// which will be initiated with a logger and a mongo session
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(
				Adapt(http.Handler(handler),
					Logger(route),
					WithConf(c),
				),
			)
	}

	return router
}
