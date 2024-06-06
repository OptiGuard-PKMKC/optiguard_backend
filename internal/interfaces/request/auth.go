package request

import (
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/helpers/customtypes"
)

type (
	RegisterValidate struct {
		Name            string `json:"name" validate:"required,min=3,max=100"`
		Email           string `json:"email" validate:"required,email"`
		Password        string `json:"password" validate:"required,min=8,max=100"`
		ConfirmPassword string `json:"confirm_password" validate:"required,min=8,max=100"`
	}

	Register struct {
		Name            string           `json:"name" validate:"required,min=3,max=100"`
		Email           string           `json:"email" validate:"required,email"`
		Password        string           `json:"password" validate:"required,min=8,max=100"`
		ConfirmPassword string           `json:"confirm_password" validate:"required,min=8,max=100"`
		Role            string           `json:"role" validate:"required"`
		Phone           string           `json:"phone" validate:"required,e164"`
		Birthdate       customtypes.Date `json:"birthdate" validate:"required"`
		Gender          string           `json:"gender" validate:"required"`
		City            string           `json:"city" validate:"required"`
		Province        string           `json:"province" validate:"required"`
		Address         string           `json:"address" validate:"required"`
	}
)

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}
