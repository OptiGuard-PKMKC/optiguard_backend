package repo_intf

import "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"

type FundusRepository interface {
	Create(fundus *entities.Fundus, details []*entities.FundusDetail) (int64, error)
	CreateFeedback() error
	FindAll() error
	FindByID() error
	FindByIDVerified() error
	Delete() error
	DeleteFeedback() error
	UpdateVerifyDoctor() error
}
