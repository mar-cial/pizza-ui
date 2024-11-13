package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/mar-cial/food-ui/internal/adapters"
	"github.com/mar-cial/food-ui/internal/handlers"
)

var (
	ErrBadPort = errors.New("UI_PORT env variable needs to be set")
)

func main() {
	config := adapters.NewEnvConfig()
	port, err := config.GetPort()
	if err != nil {
		log.Fatal(ErrBadPort.Error())
	}

	mux := http.NewServeMux()

	pageHandler := handlers.NewPagesHandler()

	mux.HandleFunc("GET /", pageHandler.Homepage)

	server := adapters.NewHTTPServer(port, mux)
	if err := server.Start(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
