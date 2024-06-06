package repo_intf

import (
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"
)

type UserRepository interface {
	Create(*entities.User) (int64, error)
	FindAll() ([]entities.User, error)
	FindByID(int64) (*entities.User, error)
	FindByIDAndRole(user_id int64, role string) (*entities.User, error)
	FindByEmail(string) (*entities.User, error)
	GetRoleID(string) (*entities.UserRole, error)
	Update(*entities.User) (*entities.User, error)
	Delete(int) error
}
