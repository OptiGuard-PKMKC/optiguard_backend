package usecases

import (
	"errors"

	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/response"
	repo_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/repositories/interfaces"
	usecase_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/usecases/interfaces"
)

type UserUsecae struct {
	userRepo repo_intf.UserRepository
}

func NewUserUsecase(userRepo repo_intf.UserRepository) usecase_intf.UserUsecase {
	return &UserUsecae{
		userRepo: userRepo,
	}
}

func (u *UserUsecae) GetProfile(userID int64) (*response.GetProfile, error) {
	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	return &response.GetProfile{
		Name:      user.Name,
		Role:      user.RoleName,
		Birthdate: user.Birthdate,
		Gender:    user.Gender,
		City:      user.City,
		Province:  user.Province,
		Address:   user.Address,
	}, nil
}

func (u *UserUsecae) UpdateProfile() error {
	return nil
}
