package repo_intf

import (
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"
)

type UserRepository interface {
	Create(*entities.User) (int64, error)
	FindAll() ([]entities.User, error)
	FindByID(int) (*entities.User, error)
	FindByEmail(string) (*entities.User, error)
	Update(*entities.User) (*entities.User, error)
	Delete(int) error
}
