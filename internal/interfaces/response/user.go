package response

import "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/helpers/customtypes"

type GetProfile struct {
	Name      string           `json:"name"`
	Role      string           `json:"role"`
	Birthdate customtypes.Date `json:"birthdate"`
	Gender    string           `json:"gender"`
	City      string           `json:"city"`
	Province  string           `json:"province"`
	Address   string           `json:"address"`
}

type (
	CurrentUser struct {
		ID   int64  `json:"id"`
		Role string `json:"role"`
	}
)
