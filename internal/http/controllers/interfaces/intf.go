package controller_intf

import (
	"net/http"
)

type AuthController interface {
	RegisterValidate(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type AppointmentController interface {
	Create(w http.ResponseWriter, r *http.Request)
	ViewAll(w http.ResponseWriter, r *http.Request)
	Confirm(w http.ResponseWriter, r *http.Request)
}

type FundusController interface {
	DetectImage(w http.ResponseWriter, r *http.Request)
	ViewFundus(w http.ResponseWriter, r *http.Request)
	DeleteFundus(w http.ResponseWriter, r *http.Request)
	RequestVerifyFundusByPatient(w http.ResponseWriter, r *http.Request)
	VerifyFundusByDoctor(w http.ResponseWriter, r *http.Request)
}

type UserController interface {
	Profile(w http.ResponseWriter, r *http.Request)
}
