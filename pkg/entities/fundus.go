package entities

import "time"

type PatientFundus struct {
	ID        int64     `json:"id"`
	PatientID int64     `json:"patient_id"`
	ImageURL  string    `json:"image_url"`
	Condition string    `json:"condition"`
	CreatedAt time.Time `json:"created_at"`
}

type FundusDetail struct {
	ID              int64   `json:"id"`
	FundusID        int64   `json:"fundus_id"`
	Disease         string  `json:"disease"`
	ConfidenceScore float64 `json:"confidence_score"`
	Description     string  `json:"description"`
}

type FundusFeedback struct {
	ID        int64     `json:"id"`
	FundusID  int64     `json:"fundus_id"`
	DoctorID  int64     `json:"doctor_id`
	Verified  bool      `json:"verified"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
