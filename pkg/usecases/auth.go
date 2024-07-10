package usecases

import (
	"errors"
	"log"

	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/request"
	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/response"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/helpers"
	repo_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/repositories/interfaces"
	usecase_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/usecases/interfaces"
)

type AuthUsecase struct {
	secretKey string
	userRepo  repo_intf.UserRepository
}

func NewAuthUsecase(secretKey string, userRepo repo_intf.UserRepository) usecase_intf.AuthUsecase {
	return &AuthUsecase{
		secretKey: secretKey,
		userRepo:  userRepo,
	}
}

func (u *AuthUsecase) RegisterValidate(p *request.RegisterValidate) error {
	existUser, err := u.userRepo.FindAll()
	if err != nil {
		log.Printf("Error finding user: %v", err)
		return err
	}

	if existUser != nil {
		return errors.New("user already exists")
	}

	if p.Password != p.ConfirmPassword {
		return errors.New("password and confirm password do not match")
	}

	return nil
}

func (u *AuthUsecase) Register(p *request.Register) (*response.Register, error) {
	// Check if password and confirm password match
	if p.Password != p.ConfirmPassword {
		return nil, errors.New("password and confirm password do not match")
	}

	// Hash password
	hashedPassword, err := helpers.HashPassword(u.secretKey, p.Password)
	if err != nil {
		return nil, err
	}

	if p.Role == "" {
		p.Role = "guest"
	}

	// Get user role id
	role, err := u.userRepo.GetRoleID(p.Role)
	if err != nil {
		return nil, errors.New("error getting role id")
	}
	if role == nil {
		return nil, errors.New("role not found")
	}

	// Create user
	user := &entities.User{
		Name:      p.Name,
		Phone:     p.Phone,
		Email:     p.Email,
		Password:  hashedPassword,
		RoleID:    role.ID,
		Birthdate: p.Birthdate,
		Gender:    p.Gender,
		City:      p.City,
		Province:  p.Province,
		Address:   p.Address,
	}

	_, err = u.userRepo.Create(user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	// Generate JWT
	paramsJWT := helpers.ParamsGenerateJWT{
		ExpiredInMinute: 60 * 24 * 30,
		SecretKey:       u.secretKey,
		UserID:          user.ID,
		UserRole:        role.RoleName,
	}

	resultJWT, err := helpers.GenerateJWT(&paramsJWT)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		return nil, err
	}

	return &response.Register{
		Name:        p.Name,
		Email:       p.Email,
		Role:        role.RoleName,
		AccessToken: resultJWT.Token,
	}, nil
}

func (u *AuthUsecase) Login(p *request.Login) (*response.Login, error) {
	// Find user by email
	user, err := u.userRepo.FindByEmail(p.Email)
	if err != nil {
		log.Printf("Error finding user by email: %v", err)
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	// Check password
	err = helpers.CheckPasswordHash(u.secretKey, p.Password, user.Password)
	if err != nil {
		return nil, errors.New("invalid password")
	}

	// Generate JWT
	paramsJWT := helpers.ParamsGenerateJWT{
		ExpiredInMinute: 60 * 24 * 30,
		SecretKey:       u.secretKey,
		UserID:          user.ID,
		UserRole:        user.RoleName,
	}

	resultJWT, err := helpers.GenerateJWT(&paramsJWT)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		return nil, err
	}

	return &response.Login{
		Name:        user.Name,
		Email:       user.Email,
		Role:        user.RoleName,
		AccessToken: resultJWT.Token,
	}, nil
}
