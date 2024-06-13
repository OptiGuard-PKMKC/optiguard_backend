package controllers

import (
	"net/http"

	controller_intf "github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/controllers/interfaces"
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/request"
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/response"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/helpers"
	usecase_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/usecases/interfaces"
)

type DoctorController struct {
	doctorUsecase usecase_intf.DoctorUsecase
}

func NewDoctorController(doctorUsecase usecase_intf.DoctorUsecase) controller_intf.DoctorController {
	return &DoctorController{
		doctorUsecase: doctorUsecase,
	}
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
