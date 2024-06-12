package controllers

import (
	"net/http"

	controller_intf "github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/controllers/interfaces"
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/request"
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/response"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/helpers"
	usecase_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/usecases/interfaces"
)

type FundusController struct {
	fundusUsecase usecase_intf.FundusUsecase
}

func NewFundusController(fundusUsecase usecase_intf.FundusUsecase) controller_intf.FundusController {
	return &FundusController{
		fundusUsecase: fundusUsecase,
	}
}

func (c *FundusController) DetectImage(w http.ResponseWriter, r *http.Request) {
	req := request.DetectFundusImage{}
	if err := helpers.JsonBodyDecoder(r.Body, &req); err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to parse request body",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	// Validate the request body
	err := validate.Struct(&req)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Validation error",
			Error:   helpers.GetValidationErrors(err),
		}, http.StatusBadRequest)
		return
	}

	fundusID, err := c.fundusUsecase.DetectImage(&req)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusBadRequest)
		return
	}

	res := response.Response{
		Status:  "success",
		Message: "Detect fundus success",
		Data:    response.FundusID{ID: fundusID},
	}

	helpers.SendResponse(w, res, http.StatusCreated)
}

func (c *FundusController) ViewFundus(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.StringToInt64(helpers.UrlVars(r, "id"))
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Invalid fundus ID",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	fundus, err := c.fundusUsecase.ViewFundus(*id)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "View fundus success",
		Data:    fundus,
	}, http.StatusOK)
}

func (c *FundusController) DeleteFundus(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.StringToInt64(helpers.UrlVars(r, "id"))
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Invalid fundus ID",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	if err = c.fundusUsecase.DeleteFundus(*id); err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Delete fundus success",
	}, http.StatusOK)
}

func (c *FundusController) RequestVerifyFundusByPatient(w http.ResponseWriter, r *http.Request) {

}

func (c *FundusController) VerifyFundusByDoctor(w http.ResponseWriter, r *http.Request) {
	fundusID, err := helpers.StringToInt64(helpers.UrlVars(r, "id"))
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Invalid fundus ID",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	req := request.VerifyFundus{}
	if err := helpers.JsonBodyDecoder(r.Body, &req); err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to parse request body",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	err = c.fundusUsecase.VerifyFundusByDoctor(int(*fundusID), int(req.DoctorID), req.Status, req.Feedbacks)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to verify fundus",
			Error:   err.Error(),
		}, http.StatusInternalServerError)
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Verify fundus success",
	}, http.StatusOK)
}
