package usecases

import (
	"errors"

	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/request"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/helpers"
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

func (u *DoctorUsecase) CreateProfile(user entities.User, p *request.CreateDoctorProfile) error {
	if user.RoleName != "doctor" {
		return errors.New("user is not a doctor")
	}

	isProfileExist, err := u.doctorRepo.FindProfileByUserID(user.ID)
	if err != nil {
		return err
	}
	if isProfileExist != nil {
		return errors.New("doctor profile already exist")
	}

	profile := &entities.DoctorProfile{
		UserID:         user.ID,
		Specialization: p.Specialization,
		STRNo:          p.STRNumber,
		BioDesc:        p.BioDesc,
	}

	practices := []*entities.DoctorPractice{}
	for _, pr := range p.Practices {
		practice := &entities.DoctorPractice{
			City:       pr.City,
			Province:   pr.Province,
			OfficeName: pr.OfficeName,
			StartDate:  pr.StartDate,
			EndDate:    pr.EndDate,
		}
		practices = append(practices, practice)
	}

	educations := []*entities.DoctorEducation{}
	for _, ed := range p.Educations {
		education := &entities.DoctorEducation{
			Degree:     ed.Degree,
			SchoolName: ed.SchoolName,
			StartDate:  ed.StartDate,
			EndDate:    ed.EndDate,
		}
		educations = append(educations, education)
	}

	_, err = u.doctorRepo.CreateProfile(profile, practices, educations)
	if err != nil {
		return err
	}

	return nil
}

func (u *DoctorUsecase) FindAll(filter *request.FilterAppointmentSchedule) ([]*entities.DoctorProfile, error) {
	daysOfWeek, err := helpers.GetDaysOfWeek(filter.StartDate, filter.EndDate)
	if err != nil {
		return nil, errors.New("failed to get days of week")
	}
	filter.DaysInt = daysOfWeek

	doctors, err := u.doctorRepo.FindAll(filter)
	if err != nil {
		return nil, err
	}

	return doctors, nil
}

func (u *DoctorUsecase) GetProfile(doctorID int64) (*entities.DoctorProfile, error) {
	doctor, err := u.doctorRepo.GetProfileByID(doctorID)
	if err != nil {
		return nil, err
	}

	return doctor, nil
}

func (u *DoctorUsecase) CreateSchedule(userID int64, params []*request.CreateDoctorSchedule) error {
	var schedules []*entities.DoctorSchedule

	for _, p := range params {
		schedule := &entities.DoctorSchedule{
			DoctorID:  userID,
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
