package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kelbwah/gogreggator/internal/database"
	"github.com/kelbwah/gogreggator/utils"
)

func (cfg *apiConfig) HandleFeedFollowsCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	}

	createdFeedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), feedFollowParams)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, createdFeedFollow)
}

func (cfg *apiConfig) HandleFeedFollowsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	retrievedFeedFollows, err := cfg.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, retrievedFeedFollows)
}

func (cfg *apiConfig) HandleFeedFollowsDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := r.PathValue("feedFollowID")
	feedFollowUUID, parseErr := uuid.Parse(feedFollowIDStr)
	if parseErr != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid feedFollowID")
		return
	}

	feedFollowParams := database.DeleteFeedFollowParams{
		ID:     feedFollowUUID,
		UserID: user.ID,
	}

	err := cfg.DB.DeleteFeedFollow(r.Context(), feedFollowParams)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, struct{}{})
}
