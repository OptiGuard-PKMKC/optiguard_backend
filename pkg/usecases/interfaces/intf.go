package usecase_intf

import (
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/request"
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/response"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"
)

type AuthUsecase interface {
	RegisterValidate(p *request.RegisterValidate) error
	Register(p *request.Register) (*response.Register, error)
	Login(p *request.Login) (*response.Login, error)
}

type AppointmentUsecase interface {
	Create(p *request.CreateAppointment) error
	FindAll(p *request.ViewAppointment) ([]*entities.Appointment, error)
}

type FundusUsecase interface {
	DetectImage(p *request.DetectFundusImage) (int64, error)
	ViewFundus(fundusID int64) (*entities.Fundus, error)
	FundusHistory(userID int64) ([]*entities.Fundus, error)
	RequestVerifyFundusByPatient() error
	VerifyFundusByDoctor(fundusID, doctorID int, status string, feedbacks []string) error
	DeleteFundus(fundusID int64) error
}

type UserUsecase interface {
	GetProfile(userID int64) (*response.GetProfile, error)
	UpdateProfile() error
}