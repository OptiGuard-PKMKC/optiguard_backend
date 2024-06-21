package entities

import "time"

type HealthFacility struct {
	ID              int64     `json:"id"`
	FacilityName    string    `json:"facility_name"`
	City            string    `json:"city"`
	Province        string    `json:"province"`
	Address         string    `json:"address"`
	AdaptorQuantity int       `json:"adaptor_quantity"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Adaptors        []Adaptor `json:"adaptors"`
}

type Adaptor struct {
	ID         int64     `json:"id"`
	FacilityID int64     `json:"facility_id"`
	DeviceCode string    `json:"device_code"`
	Used       bool      `json:"used"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserAdaptor struct {
	ID        int64     `json:"id"`
	PatientID int64     `json:"patient_id"`
	AdaptorID int64     `json:"adaptor_id"`
	Date      time.Time `json:"date"`
	StartHour time.Time `json:"start_hour"`
	EndHour   time.Time `json:"end_hour"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
