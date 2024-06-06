package response

import "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"

type (
	Fundus struct {
		ID        int64                     `json:"id"`
		ImageURL  string                    `json:"image_url"`
		Verified  bool                      `json:"verified"`
		Status    string                    `json:"status"`
		Condition string                    `json:"condition"`
		CreatedAt string                    `json:"created_at"`
		Details   []entities.FundusDetail   `json:"details"`
		Feedbacks []entities.FundusFeedback `json:"feedbacks,omitempty"`
	}

	FundusID struct {
		ID int64 `json:"fundus_id"`
	}
)
