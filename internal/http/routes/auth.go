package routes

import (
	route_intf "github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/routes/interfaces"
	"github.com/gorilla/mux"
)

func AuthRoutes(router *mux.Router, controller route_intf.Controllers) {
	// User routes
	router.HandleFunc("/auth/register", controller.Auth.Register).Methods("POST")
	router.HandleFunc("/auth/login", controller.Auth.Login).Methods("POST")
}
