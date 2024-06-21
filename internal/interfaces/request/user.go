package request

import "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/helpers/customtypes"

type PatientID struct {
	ID int64 `json:"patient_id"`
}

type (
	CurrentUser struct {
		ID   int64  `json:"id"`
		Role string `json:"role"`
	}
)

type (
	CreateDoctorProfile struct {
		Specialization string                  `json:"specialization" validate:"required"`
		STRNumber      string                  `json:"str_number" validate:"required"`
		BioDesc        string                  `json:"bio_desc" validate:"required"`
		Practices      []CreateDoctorPractice  `json:"practices" validate:"required"`
		Educations     []CreateDoctorEducation `json:"educations" validate:"required"`
	}

	CreateDoctorPractice struct {
		City       string           `json:"city" validate:"required"`
		Province   string           `json:"province" validate:"required"`
		OfficeName string           `json:"office_name" validate:"required"`
		Address    string           `json:"address" validate:"required"`
		StartDate  customtypes.Date `json:"start_date" validate:"required"`
		EndDate    customtypes.Date `json:"end_date" validate:"required"`
	}

	CreateDoctorEducation struct {
		Degree     string           `json:"degree" validate:"required"`
		SchoolName string           `json:"school_name" validate:"required"`
		StartDate  customtypes.Date `json:"start_date" validate:"required"`
		EndDate    customtypes.Date `json:"end_date" validate:"required"`
	}
)

type (
	CreateDoctorSchedule struct {
		Day       string `json:"day" validate:"required"`
		StartHour string `json:"start_hour" validate:"required"`
		EndHour   string `json:"end_hour" validate:"required"`
	}
)
