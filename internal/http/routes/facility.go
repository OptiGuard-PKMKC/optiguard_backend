package routes

import (
	"net/http"

	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/middleware"
	route_intf "github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/routes/interfaces"
	"github.com/gorilla/mux"
)

func HealthFacilityRoutes(router *mux.Router, controller route_intf.Controllers, secretKey string) {
	// Protected routes

	/*
		@desc Create schedule for adaptor
		@route /adaptor/schedule
		@method POST
		@body { "facility_id", "schedule_id", "date" }
	*/
	router.Handle(
		"/adaptor/schedule",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Facility.CreateAdaptorSchedule)),
	).Methods("POST")

	/*
		@desc Get all health facilities that have lens adaptor
		@route /adaptor/facility
		@method GET
	*/

	/*
		@desc Get all lens adaptors by facility
		@route /adaptor/facility/{id}
		@method GET
	*/
}
