package usecase_intf

import (
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/request"
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/response"
)

type AuthUsecase interface {
	RegisterValidate(p *request.RegisterValidate) error
	Register(p *request.Register) (*response.Register, error)
	Login(p *request.Login) (*response.Login, error)
}
