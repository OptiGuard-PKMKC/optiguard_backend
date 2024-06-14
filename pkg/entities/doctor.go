package entities

import "time"

type DoctorProfile struct {
	ID             int64              `json:"id"`
	UserID         int64              `json:"user_id"`
	STRNo          string             `json:"str_number"`
	Specialization string             `json:"specialization"`
	BioDesc        string             `json:"bio_desc"`
	WorkYears      int                `json:"work_years"`
	Rating         int                `json:"rating"`
	Practices      []*DoctorPractice  `json:"practices,omitempty"`
	Educations     []*DoctorEducation `json:"educations,omitempty"`
	Schedules      []*DoctorSchedule  `json:"schedules,omitempty"`
}

type DoctorPractice struct {
	ID         int64     `json:"id"`
	DoctorID   int64     `json:"doctor_id"`
	City       string    `json:"city"`
	Province   string    `json:"province"`
	OfficeName string    `json:"office_name"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
}

type DoctorEducation struct {
	ID         int64     `json:"id"`
	DoctorID   int64     `json:"doctor_id"`
	University string    `json:"university"`
	Major      string    `json:"major"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
}

type DoctorSchedule struct {
	ID        int64     `json:"id"`
	DoctorID  int64     `json:"doctor_id"`
	Day       string    `json:"day"`
	StartHour string    `json:"start_hour"`
	EndHour   string    `json:"end_hour"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
