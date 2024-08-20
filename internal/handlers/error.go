package handlers

import (
	"net/http"

	"github.com/kelbwah/gogreggator/utils"
)

func errorHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
