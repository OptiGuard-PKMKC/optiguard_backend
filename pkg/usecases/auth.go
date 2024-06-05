package usecases

import (
	"errors"
	"log"
	"time"

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

func (u *AuthUsecase) Register(p *usecase_intf.ParamsRegister) (int64, error) {
	// Check if user already exists
	existUser, err := u.userRepo.FindAll()
	if err != nil {
		log.Printf("Error finding user: %v", err)
		return 0, err
	}

	if existUser != nil {
		return 0, errors.New("user already exists")
	}

	// Check if password and confirm password match
	if p.Password != p.ConfirmPassword {
		return 0, errors.New("password and confirm password do not match")
	}

	// Hash password
	hashedPassword, err := helpers.HashPassword(u.secretKey, p.Password)
	if err != nil {
		return 0, err
	}

	currentTime := time.Now()

	// Create user
	user := &entities.User{
		Name:      p.Name,
		Email:     p.Email,
		Password:  hashedPassword,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	userID, err := u.userRepo.Create(user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return 0, err
	}

	return userID, nil
}

func (u *AuthUsecase) Login(p *usecase_intf.ParamsLogin) (*usecase_intf.ResultLogin, error) {
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
	isValid, err := helpers.CheckPasswordHash(u.secretKey, p.Password, user.Password)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, errors.New("invalid password")
	}

	// Generate JWT
	paramsJWT := helpers.ParamsGenerateJWT{
		ExpiredInMinute: 60,
		SecretKey:       u.secretKey,
		UserID:          user.ID,
	}

	resultJWT, err := helpers.GenerateJWT(&paramsJWT)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		return nil, err
	}

	return &usecase_intf.ResultLogin{
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: resultJWT.Token,
	}, nil
}
