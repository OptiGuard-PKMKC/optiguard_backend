package entities

import "time"

type DoctorProfile struct {
	ID             int64  `json:"id"`
	UserID         int64  `json:"user_id"`
	STRNo          string `json:"str_no"`
	BioDescription string `json:"bio_description"`
}

type DoctorPractice struct {
	ID         int64     `json:"id"`
	ProfileID  int64     `json:"profile_id"`
	City       string    `json:"city"`
	OfficeName string    `json:"office_name"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
}

type DoctorEducation struct {
	ID         int64     `json:"id"`
	ProfileID  int64     `json:"profile_id"`
	University string    `json:"university"`
	Major      string    `json:"major"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
}

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