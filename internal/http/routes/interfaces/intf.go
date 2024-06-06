package route_intf

import (
	controller_intf "github.com/OptiGuard-PKMKC/optiguard_backend/internal/http/controllers/interfaces"
)

type Controllers struct {
	Auth controller_intf.AuthController
	User controller_intf.UserController
}
