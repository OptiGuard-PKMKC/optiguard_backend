package controller_intf

import "net/http"

type FundusController interface {
	DetectImage(w http.ResponseWriter, r *http.Request)
	ViewFundus(w http.ResponseWriter, r *http.Request)
	DeleteFundus(w http.ResponseWriter, r *http.Request)
	RequestVerifyFundusByPatient(w http.ResponseWriter, r *http.Request)
	VerifyFundusByDoctor(w http.ResponseWriter, r *http.Request)
}
