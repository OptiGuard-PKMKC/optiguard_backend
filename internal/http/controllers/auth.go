package controllers

import (
	"net/http"

	controller_intf "github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/controllers/interfaces"
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/request"
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/response"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/helpers"
	usecase_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/usecases/interfaces"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type AuthController struct {
	authUsecase usecase_intf.AuthUsecase
}

func NewAuthController(authUsecase usecase_intf.AuthUsecase) controller_intf.AuthController {
	return &AuthController{
		authUsecase: authUsecase,
	}
}

func (c *AuthController) RegisterValidate(w http.ResponseWriter, r *http.Request) {
	req := request.RegisterValidate{}
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

	err = c.authUsecase.RegisterValidate(&req)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusBadRequest)
		return
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "User data is valid",
	}, http.StatusOK)
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	req := request.Register{}
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

	user, err := c.authUsecase.Register(&req)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to register user",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	res := response.Response{
		Status:  "success",
		Message: "User registered successfully",
		Data:    user,
	}

	helpers.SendResponse(w, res, http.StatusCreated)
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	req := request.Login{}
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

	user, err := c.authUsecase.Login(&req)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to login user",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	res := response.Response{
		Status:  "success",
		Message: "User logged in successfully",
		Data:    user,
	}

	helpers.SendResponse(w, res, http.StatusOK)
}
