package controllers

import (
	"net/http"

	controller_intf "github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/controllers/interfaces"
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/response"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/helpers"
	usecase_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/usecases/interfaces"
)

type UserController struct {
	userUsecase usecase_intf.UserUsecase
}

func NewUserController(userUsecase usecase_intf.UserUsecase) controller_intf.UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (c *UserController) Profile(w http.ResponseWriter, r *http.Request) {
	currentUser, err := helpers.GetCurrentUser(r)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to get current user",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	user, err := c.userUsecase.GetProfile(currentUser.ID)
	if err != nil {
		helpers.SendResponse(w, response.Response{
			Status:  "error",
			Message: "Failed to get user profile",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	helpers.SendResponse(w, response.Response{
		Status:  "success",
		Message: "User profile",
		Data:    user,
	}, http.StatusOK)
}
