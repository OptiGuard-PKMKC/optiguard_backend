package repo_intf

import (
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/request"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"
)

type AppointmentRepository interface {
	Create(apt *entities.Appointment) error
	FindAll(doctorID *int64, patientID *int64) ([]*entities.Appointment, error)
	UpdateStatus(aptID int64, status string) error
	Delete(id int64) error
}

type FundusRepository interface {
	Create(fundus *entities.Fundus, details []*entities.FundusDetail) (int64, error)
	CreateFeedback(feedback []entities.FundusFeedback) error
	FindAll() error
	FindByID(id int64) (*entities.Fundus, error)
	FindFundusDetails(fundusID int64) ([]*entities.FundusDetail, error)
	FindFundusFeedbacks(fundusID int64) ([]*entities.FundusFeedback, error)
	FindByIDVerified() error
	Delete(id int64) error
	DeleteFeedback() error
	UpdateVerify(fundusID, doctorID int, status string) error
}

type UserRepository interface {
	Create(*entities.User) (int64, error)
	FindAll() ([]entities.User, error)
	FindByID(int64) (*entities.User, error)
	FindByIDAndRole(user_id int64, role string) (*entities.User, error)
	FindByEmail(string) (*entities.User, error)
	GetRoleID(string) (*entities.UserRole, error)
	Update(*entities.User) (*entities.User, error)
	Delete(int) error
}

type DoctorRepository interface {
	CreateProfile(profile *entities.DoctorProfile, practices []*entities.DoctorPractice, educations []*entities.DoctorEducation) (*int64, error)
	FindAll(filter *request.FilterAppointmentSchedule) ([]*entities.DoctorProfile, error)
	FindProfileByUserID(userID int64) (*int64, error)
	GetProfileByID(profileID int64) (*entities.DoctorProfile, error)
	GetPractice(profileID int64) ([]*entities.DoctorPractice, error)
	GetSchedule(profileID int64) ([]*entities.DoctorSchedule, error)
	CreateSchedule(schedules []*entities.DoctorSchedule) error
}
