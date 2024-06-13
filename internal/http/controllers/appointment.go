package controllers

import (
	"net/http"

	controller_intf "github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/controllers/interfaces"
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/request"
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/response"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/helpers"
	usecase_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/usecases/interfaces"
)

type AppointmentController struct {
	aptUsecase usecase_intf.AppointmentUsecase
}

func NewAppointmentController(aptUsecase usecase_intf.AppointmentUsecase) controller_intf.AppointmentController {
	return &AppointmentController{
		aptUsecase: aptUsecase,
	}
}

func (c *AppointmentController) Create(w http.ResponseWriter, r *http.Request) {
	req := request.CreateAppointment{}
	if err := helpers.JsonBodyDecoder(r.Body, &req); err != nil {
		helpers.FailedParsingBody(w, err)
		return
	}

	if err := c.aptUsecase.Create(&req); err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Create appointment success",
	}, http.StatusCreated)
}

func (c *AppointmentController) ViewAll(w http.ResponseWriter, r *http.Request) {
	currentUser, err := helpers.GetCurrentUser(r)
	if err != nil {
		helpers.FailedGetCurrentUser(w, err)
		return
	}

	user := request.ViewAppointment{
		UserID:   currentUser.ID,
		UserRole: currentUser.Role,
	}

	apts, err := c.aptUsecase.FindAll(&user)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Get appointments success",
		Data:    apts,
	}, http.StatusOK)
}

func (c *AppointmentController) Confirm(w http.ResponseWriter, r *http.Request) {
	req := request.ConfirmAppointment{}
	if err := helpers.JsonBodyDecoder(r.Body, &req); err != nil {
		helpers.FailedParsingBody(w, err)
		return
	}

	// Validate request body
	err := validate.Struct(&req)
	if err != nil {
		helpers.FailedValidation(w, err)
		return
	}

	aptID, err := helpers.StringToInt64(helpers.UrlVars(r, "apt_id"))
	if err != nil {
		helpers.FailedGetUrlVars(w, err, nil)
		return
	}

	if err := c.aptUsecase.UpdateStatus(*aptID, req.Confirm); err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Update appointment status success",
	}, http.StatusOK)
}
