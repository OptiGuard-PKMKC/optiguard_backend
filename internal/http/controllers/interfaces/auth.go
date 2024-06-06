package controller_intf

import (
	"net/http"
)

type AuthController interface {
	RegisterValidate(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}
