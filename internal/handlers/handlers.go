package handlers

import "net/http"

func HandlersInit(mux *http.ServeMux) {
	mux.HandleFunc("GET /v1/healthz", readinessHandler)
	mux.HandleFunc("GET /v1/err", errorHandler)
}
