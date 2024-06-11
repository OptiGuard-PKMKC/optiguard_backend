package repo_intf

import "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"

type FundusRepository interface {
	Create(fundus *entities.Fundus, details []*entities.FundusDetail) (int64, error)
	CreateFeedback(feedback []*entities.FundusFeedback) error
	FindAll() error
	FindByID(id int64) (*entities.Fundus, error)
	FindFundusDetails(fundus_id int64) ([]*entities.FundusDetail, error)
	FindFundusFeedbacks(fundus_id int64) ([]*entities.FundusFeedback, error)
	FindByIDVerified() error
	Delete() error
	DeleteFeedback() error
	UpdateVerify() error
}
