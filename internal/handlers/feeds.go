package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kelbwah/gogreggator/internal/database"
	"github.com/kelbwah/gogreggator/utils"
)

func (cfg *apiConfig) HandleFeedsCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	}

	createdFeed, err := cfg.DB.CreateFeed(r.Context(), feedParams)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    feedParams.ID,
		UserID:    user.ID,
	}

	createdFeedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), feedFollowParams)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	utils.RespondWithJSON(w, http.StatusCreated, struct {
		Feed       database.Feed       `json:"feed"`
		FeedFollow database.FeedFollow `json:"feed_follow"`
	}{
		Feed:       createdFeed,
		FeedFollow: createdFeedFollow,
	})
}

func (cfg *apiConfig) HandleFeedsGet(w http.ResponseWriter, r *http.Request) {
	createdFeed, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, createdFeed)
}
