package handlers

import (
	"net/http"

	"github.com/kelbwah/gogreggator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func createApiCfg(db *database.Queries) *apiConfig {
	return &apiConfig{
		DB: db,
	}
}

func HandlersInit(mux *http.ServeMux, db *database.Queries) {
	// Initialize apiCfg
	apiCfg := createApiCfg(db)

	/* -- GET ENDPOINTS -- */
	mux.HandleFunc("GET /v1/healthz", readinessHandler) // Sanity check endpoint
	mux.HandleFunc("GET /v1/err", errorHandler)         // Sanity check endpoint
	mux.HandleFunc("GET /v1/users", apiCfg.middlewareAuth(apiCfg.HandleUsersGet))
	mux.HandleFunc("GET /v1/feeds", apiCfg.HandleFeedsGet)
	mux.HandleFunc("GET /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.HandleFeedFollowsGet))

	/* -- POST ENDPOINTS -- */
	mux.HandleFunc("POST /v1/users", apiCfg.HandleUsersCreate)
	mux.HandleFunc("POST /v1/feeds", apiCfg.middlewareAuth(apiCfg.HandleFeedsCreate))
	mux.HandleFunc("POST /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.HandleFeedFollowsCreate))

	/* -- DELETE ENDPOINTS */
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.HandleFeedFollowsDelete))
}
