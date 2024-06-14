package controllers

import (
	"net/http"

	controller_intf "github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/controllers/interfaces"
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/request"
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/response"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/helpers"
	usecase_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/usecases/interfaces"
	"github.com/mitchellh/mapstructure"
)

type DoctorController struct {
	doctorUsecase usecase_intf.DoctorUsecase
}

func NewDoctorController(doctorUsecase usecase_intf.DoctorUsecase) controller_intf.DoctorController {
	return &DoctorController{
		doctorUsecase: doctorUsecase,
	}
}

func (c *DoctorController) ViewAll(w http.ResponseWriter, r *http.Request) {
	var filterQuery map[string]string
	if err := helpers.QueryDecoder(r, &filterQuery); err != nil {
		helpers.FailedParsingQuery(w, err)
		return
	}

	filter := &request.FilterAppointmentSchedule{}
	if err := mapstructure.Decode(filterQuery, filter); err != nil {
		helpers.FailedParsingQuery(w, err)
		return
	}

	doctors, err := c.doctorUsecase.FindAll(filter)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Fetch all doctors",
		Data:    doctors,
	}, http.StatusOK)
}

func (c *DoctorController) Profile(w http.ResponseWriter, r *http.Request) {
	doctorID, err := helpers.StringToInt64(helpers.UrlVars(r, "id"))
	if err != nil {
		helpers.FailedGetUrlVars(w, err, nil)
		return
	}

	doctor, err := c.doctorUsecase.GetProfile(*doctorID)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Fetch doctor profile success",
		Data:    doctor,
	}, http.StatusOK)
}

func (c *DoctorController) CreateSchedule(w http.ResponseWriter, r *http.Request) {
	req := []*request.CreateDoctorSchedule{}
	if err := helpers.JsonBodyDecoder(r.Body, &req); err != nil {
		helpers.FailedParsingBody(w, err)
		return
	}

	for _, s := range req {
		err := validate.Struct(s)
		if err != nil {
			helpers.FailedValidation(w, err)
			return
		}
	}

	currentUser, err := helpers.GetCurrentUser(r)
	if err != nil {
		helpers.FailedGetCurrentUser(w, err)
		return
	}

	if err = c.doctorUsecase.CreateSchedule(currentUser.ID, req); err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusInternalServerError)
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "Schedule created",
	}, http.StatusCreated)
}
