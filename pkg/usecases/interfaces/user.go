package usecase_intf

import "github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/response"

type UserUsecase interface {
	GetProfile(userID int64) (*response.GetProfile, error)
	UpdateProfile() error
}
