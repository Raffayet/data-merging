package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Raffayet/data-merging/backend/internal/service"
)

// ProfileHandler handles HTTP requests for profiles
type ProfileHandler struct {
	profileService service.ProfileService
}

func NewProfileHandler(profileService service.ProfileService) *ProfileHandler {
	return &ProfileHandler{profileService: profileService}
}

// GetProfilesHandler handles GET requests to fetch profiles
func (h *ProfileHandler) GetProfilesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	profiles, err := h.profileService.GetProfiles(ctx)
	if err != nil {
		http.Error(w, "Unable to fetch profiles", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(profiles)
}
