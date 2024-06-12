package request

type (
	DetectFundusImage struct {
		PatientID int64  `json:"patient_id" validate:"required"`
		ImageBlob string `json:"image_blob" validate:"required"`
	}

	VerifyFundus struct {
		DoctorID  int64    `json:"doctor_id"`
		Status    string   `json:"status"`
		Feedbacks []string `json:"feedbacks"`
	}
)
