package entities

import "time"

type Fundus struct {
	ID        int64             `json:"id"`
	PatientID int64             `json:"patient_id"`
	ImagePath string            `json:"image_path"`
	ImageBlob string            `json:"image_blob,omitempty"`
	Verified  bool              `json:"verified"`
	Status    string            `json:"status"`
	Condition string            `json:"condition"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Detail    []*FundusDetail   `json:"details,omitempty"`
	Feedback  []*FundusFeedback `json:"feedbacks,omitempty"`
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
	DoctorID  int64     `json:"doctor_id"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
