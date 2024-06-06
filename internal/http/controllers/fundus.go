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
