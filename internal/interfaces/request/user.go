package request

type PatientID struct {
	ID int64 `json:"patient_id"`
}

type (
	CreateDoctorProfile struct {
		Specialization string `json:"specialization" validate:"required"`
		STRNumber      string `json:"str_number" validate:"required"`
		BioDesc        string `json:"bio_desc" validate:"required"`
	}

	CreateDoctorSchedule struct {
		Day       string `json:"day" validate:"required"`
		StartHour string `json:"start_hour" validate:"required"`
		EndHour   string `json:"end_hour" validate:"required"`
	}
)
