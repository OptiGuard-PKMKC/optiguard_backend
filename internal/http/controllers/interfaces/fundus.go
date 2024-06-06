package controller_intf

import "net/http"

type FundusController interface {
	DetectImage(w http.ResponseWriter, r *http.Request)
}
