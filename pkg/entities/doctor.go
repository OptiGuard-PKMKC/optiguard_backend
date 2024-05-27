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
