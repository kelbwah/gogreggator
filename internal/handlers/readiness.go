package handlers

import (
	"net/http"

	"github.com/kelbwah/gogreggator/utils"
)

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	payload := map[string]string{
		"status": "ok",
	}
	utils.RespondWithJSON(w, http.StatusOK, payload)
}
