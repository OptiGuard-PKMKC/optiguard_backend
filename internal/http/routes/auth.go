package routes

import (
	route_intf "github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/routes/interfaces"
	"github.com/gorilla/mux"
)

func AuthRoutes(router *mux.Router, controller route_intf.Controllers) {
	router.HandleFunc("/auth/register/validate", controller.Auth.RegisterValidate).Methods("POST")
	router.HandleFunc("/auth/register/complete", controller.Auth.Register).Methods("POST")
	router.HandleFunc("/auth/login", controller.Auth.Login).Methods("POST")
}
