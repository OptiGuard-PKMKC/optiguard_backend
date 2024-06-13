package usecases

import (
	"errors"

	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/request"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"
	repo_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/repositories/interfaces"
	usecase_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/usecases/interfaces"
)

type AppointmentUsecase struct {
	aptRepo repo_intf.AppointmentRepository
}

func NewAppointmentUsecase(aptRepo repo_intf.AppointmentRepository) usecase_intf.AppointmentUsecase {
	return &AppointmentUsecase{
		aptRepo: aptRepo,
	}
}

func (u *AppointmentUsecase) Create(p *request.CreateAppointment) error {
	apt := &entities.Appointment{
		PatientID: p.PatientID,
		DoctorID:  p.DoctorID,
		Date:      p.Date,
		StartHour: p.StartHour,
		EndHour:   p.EndHour,
		Status:    "pending",
	}
	if err := u.aptRepo.Create(apt); err != nil {
		return err
	}

	return nil
}

func (u *AppointmentUsecase) FindAll(p *request.ViewAppointment) ([]*entities.Appointment, error) {
	var apts []*entities.Appointment
	var err error

	if p.UserRole == "doctor" {
		apts, err = u.aptRepo.FindAll(&p.UserID, nil)
	} else if p.UserRole == "patient" {
		apts, err = u.aptRepo.FindAll(nil, &p.UserID)
	}
	if err != nil {
		return nil, errors.New("failed to get appointments")
	}

	return apts, nil
}
