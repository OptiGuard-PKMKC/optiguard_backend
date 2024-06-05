package controllers

import (
	"net/http"

	controller_intf "github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/controllers/interfaces"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/helpers"
	usecase_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/usecases/interfaces"
)

type AuthController struct {
	authUsecase usecase_intf.AuthUsecase
}

func NewAuthController(authUsecase usecase_intf.AuthUsecase) controller_intf.AuthController {
	return &AuthController{
		authUsecase: authUsecase,
	}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	req := controller_intf.RegisterRequest{}
	if err := helpers.JsonBodyDecoder(r.Body, &req); err != nil {
		helpers.SendResponse(w, controller_intf.Response{
			Status: "error",
			Error:  "Failed to parse request body",
		}, http.StatusBadRequest)
		return
	}

	_, err := c.authUsecase.Register(&usecase_intf.ParamsRegister{
		Name:            req.Name,
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	})
	if err != nil {
		helpers.SendResponse(w, controller_intf.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusBadRequest)
		return
	}

	res := controller_intf.Response{
		Status:  "success",
		Message: "User registered successfully",
	}

	helpers.SendResponse(w, res, http.StatusCreated)
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	req := controller_intf.LoginRequest{}
	if err := helpers.JsonBodyDecoder(r.Body, req); err != nil {
		helpers.SendResponse(w, controller_intf.Response{
			Status: "error",
			Error:  "Failed to parse request body",
		}, http.StatusBadRequest)
		return
	}

	user, err := c.authUsecase.Login(&usecase_intf.ParamsLogin{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		helpers.SendResponse(w, controller_intf.Response{
			Status: "error",
			Error:  err.Error(),
		}, http.StatusBadRequest)
		return
	}

	res := controller_intf.Response{
		Status:  "success",
		Message: "User logged in successfully",
		Data:    user,
	}

	helpers.SendResponse(w, res, http.StatusOK)
}
