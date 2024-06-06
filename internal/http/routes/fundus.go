package routes

import (
	"net/http"

	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/middleware"
	route_intf "github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/routes/interfaces"
	"github.com/gorilla/mux"
)

func FundusRoutes(router *mux.Router, controller route_intf.Controllers, secretKey string) {
	// Protected routes
	router.Handle(
		"/fundus/detect",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Fundus.DetectImage)),
	).Methods("POST")
}
