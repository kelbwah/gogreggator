package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelbwah/gogreggator/internal/handlers"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env vars: %v", err)
	}
	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	handlers.HandlersInit(mux)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Listening on port ':%s'\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
