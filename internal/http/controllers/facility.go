package controllers

import (
	"net/http"

	controller_intf "github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/controllers/interfaces"
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/request"
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/response"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/helpers"
	usecase_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/usecases/interfaces"
)

type HealthFacilityController struct {
	facilityUsecase usecase_intf.HealthFacilityUsecase
}

func NewHealthFacilityController(facilityUsecase usecase_intf.HealthFacilityUsecase) controller_intf.HealthFacilityController {
	return &HealthFacilityController{
		facilityUsecase: facilityUsecase,
	}
}

func (c *HealthFacilityController) CreateAdaptorSchedule(w http.ResponseWriter, r *http.Request) {
	currentUser, err := helpers.GetCurrentUser(r)
	if err != nil {
		helpers.FailedGetCurrentUser(w, err)
		return
	}

	req := request.CreateAdaptorSchedule{}
	if err := helpers.JsonBodyDecoder(r.Body, &req); err != nil {
		helpers.FailedParsingBody(w, err)
		return
	}

	// Validate req body
	if err := validate.Struct(&req); err != nil {
		helpers.FailedValidation(w, err)
		return
	}

	if err = c.facilityUsecase.CreateSchedule(*currentUser, &req); err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "successfully created adaptor schedule",
	}, http.StatusCreated)
}

func (c *HealthFacilityController) ViewAllFacility(w http.ResponseWriter, r *http.Request) {
	facilities, err := c.facilityUsecase.FindAll()
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "All health facilities retrieved successfully",
		Data:    facilities,
	}, http.StatusOK)
}

func (c *HealthFacilityController) ViewAllLensAdaptorByFacility(w http.ResponseWriter, r *http.Request) {
	facilityID, err := helpers.StringToInt64(helpers.UrlVars(r, "id"))
	if err != nil {
		helpers.FailedGetUrlVars(w, err, nil)
		return
	}

	facilities, err := c.facilityUsecase.FindByID(*facilityID)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "All lens adaptors retrieved successfully",
		Data:    facilities,
	}, http.StatusOK)
}
