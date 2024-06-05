package entities

import "time"

type DoctorSchedule struct {
	ID        int64     `json:"id"`
	ProfileID int64     `json:"profile_id"`
	StartDay  string    `json:"start_day"`
	EndDay    string    `json:"end_day"`
	StartHour time.Time `json:"start_hour"`
	EndHour   time.Time `json:"end_hour"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DoctorPatientSchedule struct {
	ID        int64     `json:"id"`
	PatientID int64     `json:"patient_id"`
	DoctorID  int64     `json:"doctor_id"`
	Date      time.Time `json:"date"`
	StartHour time.Time `json:"start_hour"`
	EndHour   time.Time `json:"end_hour"`
	Confirm   bool      `json:"confirm"`
	CreatedAt time.Time `json:"created_at"`
}
