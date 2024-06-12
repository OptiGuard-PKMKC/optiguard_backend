package usecase_intf

import (
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/request"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"
)

type FundusUsecase interface {
	DetectImage(p *request.DetectFundusImage) (int64, error)
	ViewFundus(fundusID int64) (*entities.Fundus, error)
	FundusHistory(userID int64) ([]*entities.Fundus, error)
	RequestVerifyFundusByPatient() error
	VerifyFundusByDoctor(fundusID, doctorID int, status string, feedbacks []string) error
	DeleteFundus(fundusID int64) error
}
