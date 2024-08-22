package handlers

import (
	"net/http"

	"github.com/kelbwah/gogreggator/internal/types"
)

func HandlersInit(mux *http.ServeMux, apiCfg *types.APIConfig) {
	/* -- GET ENDPOINTS -- */
	mux.HandleFunc("GET /v1/healthz", readinessHandler)
	mux.HandleFunc("GET /v1/err", errorHandler)

	/* -- POST ENDPOINTS -- */
	mux.HandleFunc("POST /v1/users", apiCfg.HandleUsersCreate)
}
