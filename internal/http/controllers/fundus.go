package controllers

import (
	"encoding/base64"
	"io"
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
	// Request image file
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to parse form",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	// Get image file
	file, _, err := r.FormFile("fundus_image")
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to get image file",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read file content
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to read image file",
			Error:   err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	// Get current user
	user, err := helpers.GetCurrentUser(r)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to get current user",
			Error:   err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	// Encode image to base64
	req := request.DetectFundusImage{
		PatientID:   user.ID,
		FundusImage: base64.StdEncoding.EncodeToString(fileBytes),
	}

	newFundus, message, err := c.fundusUsecase.DetectImage(&req)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusBadRequest)
		return
	}

	if message != "" {
		helpers.SendResponse(w, response.Response{
			Status:  "success",
			Message: message,
		}, http.StatusOK)
		return
	}

	resData := response.Fundus{
		ID:        newFundus.ID,
		ImageBlob: newFundus.ImageBlob,
		Verified:  newFundus.Verified,
		Status:    newFundus.Status,
		Condition: newFundus.Condition,
		CreatedAt: newFundus.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	res := response.Response{
		Status:  "success",
		Message: "Detect fundus success",
		Data:    resData,
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
