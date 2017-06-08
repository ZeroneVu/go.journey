package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ZeroneVu/go.journey/minibar/config"
	"github.com/ZeroneVu/go.journey/minibar/routing"
)

// Adapter ...
// Adapter is a wrapper handler
type Adapter func(http.Handler) http.Handler

// Adapt ...
// Wraps each handler with the middleware provided
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

type key int

const (
	confKey key = iota
)

// Log ...
// A home-made logger to log the route calls to standard output
func Logger(route routing.Route) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// server http
			h.ServeHTTP(w, r)

			// by doing this the log happen after all http request
			log.Printf(
				"%s\t%s\t%s\t%s",
				route.Method,
				route.Pattern,
				route.Name,
				time.Since(start),
			)
		})
	}
}

// WithConf ...
// Copies the configuration object around
// to all middleware for ease of access to global variables
func WithConf(c config.Config) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Copy the config object to the request context
			ctx := r.Context()
			ctx = context.WithValue(r.Context(), confKey, c)
			r = r.WithContext(ctx)

			// server http
			h.ServeHTTP(w, r)
		})
	}
}
