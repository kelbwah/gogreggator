package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelbwah/gogreggator/internal/database"
	"github.com/kelbwah/gogreggator/internal/handlers"
	"github.com/kelbwah/gogreggator/internal/scraper"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")
	env, err := initVars()
	if err != nil {
		log.Fatalf("Error loading var(s): %s\n", err.Error())
	}

	db, err := sql.Open("postgres", env["dbUrl"])
	if err != nil {
		log.Fatalf("Database Error: %s\n", err.Error())
	}
	dbQueries := database.New(db)

	mux := http.NewServeMux()
	handlers.HandlersInit(mux, dbQueries)

	server := &http.Server{
		Addr:    ":" + env["port"],
		Handler: mux,
	}

	go scraper.ScrapeFeedData()

	log.Printf("Listening on port ':%s'\n", env["port"])
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}

func initVars() (map[string]string, error) {
	varMap := make(map[string]string)

	port := os.Getenv("PORT")
	if port == "" {
		return nil, errors.New("PORT environment variable is not set")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		return nil, errors.New("DATABASE_URL environment variable is not set")
	}

	varMap["port"] = port
	varMap["dbUrl"] = dbUrl

	return varMap, nil
}
