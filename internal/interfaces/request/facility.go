package request

import "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/helpers/customtypes"

type (
	CreateAdaptorSchedule struct {
		FacilityID int              `json:"facility_id" validate:"required"`
		ScheduleID int              `json:"schedule_id" validate:"required"`
		Date       customtypes.Date `json:"date" validate:"required"`
		AdaptorID  int              `json:"adaptor_id" validate:"omitempty"`
	}
)
