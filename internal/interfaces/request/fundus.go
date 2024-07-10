package request

type (
	DetectFundusImage struct {
		PatientID   int64  `json:"patient_id"`
		FundusImage string `json:"fundus_image"`
	}

	VerifyFundus struct {
		DoctorID  int64    `json:"doctor_id"`
		Status    string   `json:"status"`
		Feedbacks []string `json:"feedbacks"`
	}
)
