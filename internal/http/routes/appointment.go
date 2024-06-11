package routes

import (
	"net/http"

	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/middleware"
	route_intf "github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/routes/interfaces"
	"github.com/gorilla/mux"
)

func AppointmentRoutes(router *mux.Router, controller route_intf.Controllers, secretKey string) {
	// Protected routes
	router.Handle(
		"/appointment",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.User.Profile)),
	).Methods("POST")

	/*
		@desc Create appointment
		@route /appointment
		@method POST
		@body { "doctor_id", "date", "start_hour", "end_hour" }
	*/

	/*
		@desc Confirm appointment by doctor
		@router /appointment/confirm/{apt_id}
		@method POST
		@body { "confirm" }
	*/
}
