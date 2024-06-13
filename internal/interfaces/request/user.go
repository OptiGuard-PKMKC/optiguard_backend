package request

type PatientID struct {
	ID int64 `json:"patient_id"`
}

type (
	CreateDoctorSchedule struct {
		Day       string `json:"day" validate:"required"`
		StartHour string `json:"start_hour" validate:"required"`
		EndHour   string `json:"end_hour" validate:"required"`
	}
)
