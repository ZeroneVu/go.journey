package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ZeroneVu/go.journey/minibar/config"
	"github.com/ZeroneVu/go.journey/minibar/routing"
)

func main() {

	// Load in the config file and define the environment
	c := config.NewConfig("./config.toml")

	// Define the port
	port := fmt.Sprintf(":%d", c.Server.Port)
	if port == ":0" {
		port = ":8080"
	}

	// Create the middleware stack
	stack := NewStack(routing.MyRoutes, c)

	// Listen and serve using the middleware stack
	log.Printf("Server is locally listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, stack))
}
