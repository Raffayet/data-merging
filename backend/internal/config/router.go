package config

import (
	"github.com/Raffayet/data-merging/backend/internal/api"
	"github.com/gorilla/mux"
)

func SetupRouter(profileHandler *api.ProfileHandler) *mux.Router {
	// Setup HTTP routing
	router := mux.NewRouter()
	router.HandleFunc("/profiles", profileHandler.GetProfilesHandler).Methods("GET")
	return router
}
