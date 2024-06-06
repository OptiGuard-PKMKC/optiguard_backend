package controller_intf

import "net/http"

type UserController interface {
	Profile(w http.ResponseWriter, r *http.Request)
}
