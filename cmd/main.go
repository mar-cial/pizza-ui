package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/mar-cial/food-ui/internal/handlers"
)

var (
	ErrBadPort = errors.New("UI_PORT env variable needs to be set")
)

func main() {
	port := os.Getenv("UI_PORT")
	if port == "" {
		log.Fatalln(ErrBadPort)
	}

	mux := http.NewServeMux()

	pageHandler := handlers.NewPagesHandler()

	mux.HandleFunc("/", pageHandler.Homepage)

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	log.Printf("UI server running on port %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
