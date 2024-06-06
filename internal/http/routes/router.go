package routes

import (
	route_intf "github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/routes/interfaces"
	"github.com/gorilla/mux"
)

func SetupRouter(secretKey string, c route_intf.Controllers) *mux.Router {
	router := mux.NewRouter()

	groupRouter := router.PathPrefix("/api").Subrouter()

	// Auth routes
	AuthRoutes(groupRouter, c)
	FundusRoutes(groupRouter, c, secretKey)
	UserRoutes(groupRouter, c, secretKey)

	return router
}
