package usecases

import (
	"errors"

	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/request"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"
	repo_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/repositories/interfaces"
	usecase_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/usecases/interfaces"
)

type HealthFacilityUsecase struct {
	facilityRepo repo_intf.HealthFacilityRepository
}

func NewHealthFacilityUsecase(facilityRepo repo_intf.HealthFacilityRepository) usecase_intf.HealthFacilityUsecase {
	return &HealthFacilityUsecase{
		facilityRepo: facilityRepo,
	}
}

func (u *HealthFacilityUsecase) CreateSchedule(user request.CurrentUser, p *request.CreateAdaptorSchedule) error {
	if user.Role != "patient" {
		return errors.New("user is not patient")
	}

	if err := u.facilityRepo.CreateSchedule(user.ID, p); err != nil {
		return err
	}

	return nil
}

func (u *HealthFacilityUsecase) FindAll() ([]*entities.HealthFacility, error) {
	facilities, err := u.facilityRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return facilities, nil
}

func (u *HealthFacilityUsecase) FindByID(id int64) (*entities.HealthFacility, error) {
	facility, err := u.facilityRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return facility, nil
}
