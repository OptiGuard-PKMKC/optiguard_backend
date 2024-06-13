package request

import "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/helpers/customtypes"

type (
	CreateAppointment struct {
		PatientID int64            `json:"patient_id" validate:"required"`
		DoctorID  int64            `json:"doctor_id" validate:"required"`
		Date      customtypes.Date `json:"date" validate:"required"`
		StartHour string           `json:"start_hour" validate:"required"`
		EndHour   string           `json:"end_hour" validate:"required"`
	}

	ViewAppointment struct {
		UserID   int64
		UserRole string
	}
)
