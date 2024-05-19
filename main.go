package main

import (
	"github.com/jayphan14/simplewebserver/handlers"
	"github.com/jayphan14/simplewebserver/server"
)

func main() {
	// Create server
	s := server.NewServer("8080")
	s.AddEndpoint("GET", "/v1/users", handlers.GetUsers)
	s.StartServer()
}
