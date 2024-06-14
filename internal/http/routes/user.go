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
		@desc Create doctor profile
		@route /user/doctor
		@method POST
		@body { "specialization", "str_number", "bio_desc", []"practices"{ "city", "province", "office_name", "address", "start_date", "end_date" }, "[]educations{ "degree", "school_name", "start_date", "end_date" }" }
	*/
	router.Handle(
		"/user/doctor",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Doctor.CreateProfile)),
	).Methods("POST")

	/*
		@desc Get all doctor profile by patient
		@route /user/doctor/profile?start_date={start_date}&end_date={end_date}&start_hour={start_hour}&end_hour={end_hour}
		@method GET
	*/
	router.Handle(
		"/user/doctor/profile",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Doctor.ViewAll)),
	).Methods("GET")

	/*
		@desc Get doctor profile by patient
		@route /user/doctor/profile/{id}
		@method GET
	*/
	router.Handle(
		"/user/doctor/profile/{id}",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Doctor.Profile)),
	).Methods("GET")

	/*
		@desc Create available schedule for doctor
		@route /user/doctor/schedule
		@method POST
		@body []{ "day", "start_hour", "end_hour" }
	*/
	router.Handle(
		"/user/doctor/schedule",
		middleware.Authentication(secretKey, http.HandlerFunc(controller.Doctor.CreateSchedule)),
	).Methods("POST")

	/*
		@desc Update schedule for doctor
		@route /user/doctor/schedule
		@method PUT
		@body { "id" }
	*/

	/*
		@route /user/doctor/profile
		@method POST
		@body { "" }
	*/
}
