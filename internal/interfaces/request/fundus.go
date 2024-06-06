package request

type (
	DetectFundusImage struct {
		PatientID int64  `json:"patient_id" validate:"required"`
		ImageBlob string `json:"image_blob" validate:"required"`
	}
)
