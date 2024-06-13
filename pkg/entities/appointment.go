package entities

import (
	"time"

	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/helpers/customtypes"
)

type Appointment struct {
	ID        int64            `json:"id"`
	PatientID int64            `json:"patient_id"`
	DoctorID  int64            `json:"doctor_id"`
	Date      customtypes.Date `json:"date"`
	StartHour string           `json:"start_hour"`
	EndHour   string           `json:"end_hour"`
	Status    string           `json:"status"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}
