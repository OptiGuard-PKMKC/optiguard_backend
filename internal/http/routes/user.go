package routes

import (
	"net/http"

	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/middleware"
	route_intf "github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/routes/interfaces"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router, controller route_intf.Controllers, secretKey string) {
	// Protected routes
	router.Handle(
		"/user/profile",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.User.Profile)),
	).Methods("GET")

	/*
		@desc Get doctor profile by patient
		@route /user/doctor/profile/{id}
		@method GET
	*/

	/*
		@desc Create schedule for doctor
		@route /user/doctor/schedule
		@method POST
		@body { "start_day", "end_day", "start_hour", "end_hour" }
	*/

	/*
		@desc Update schedule for doctor
		@route /user/doctor/schedule
		@method PUT
		@body { "start_day", "end_day", "start_hour", "end_hour" }
	*/

	/*
		@route /user/doctor/profile
		@method POST
		@body { "" }
	*/
}
