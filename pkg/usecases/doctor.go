package usecases

import (
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/request"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"
	repo_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/repositories/interfaces"
	usecase_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/usecases/interfaces"
)

type DoctorUsecase struct {
	doctorRepo repo_intf.DoctorRepository
}

func NewDoctorUsecase(doctorRepo repo_intf.DoctorRepository) usecase_intf.DoctorUsecase {
	return &DoctorUsecase{
		doctorRepo: doctorRepo,
	}
}

func (u *DoctorUsecase) CreateSchedule(userID int64, params []*request.CreateDoctorSchedule) error {
	var schedules []*entities.DoctorSchedule

	for _, p := range params {
		schedule := &entities.DoctorSchedule{
			ProfileID: userID,
			Day:       p.Day,
			StartHour: p.StartHour,
			EndHour:   p.EndHour,
		}
		schedules = append(schedules, schedule)
	}

	err := u.doctorRepo.CreateSchedule(schedules)
	if err != nil {
		return err
	}

	return nil
}
