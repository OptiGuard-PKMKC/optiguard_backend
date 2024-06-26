package usecases

import (
	"errors"
	"log"

	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/request"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/helpers"
	repo_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/repositories/interfaces"
	usecase_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/usecases/interfaces"
)

type FundusUsecase struct {
	mlApiKey   string
	fundusRepo repo_intf.FundusRepository
	userRepo   repo_intf.UserRepository
}

func NewFundusUsecase(mlApiKey string, fundusRepo repo_intf.FundusRepository, userRepo repo_intf.UserRepository) usecase_intf.FundusUsecase {
	return &FundusUsecase{
		mlApiKey:   mlApiKey,
		fundusRepo: fundusRepo,
		userRepo:   userRepo,
	}
}

func (u *FundusUsecase) DetectImage(p *request.DetectFundusImage) (int64, error) {
	validPatient, err := u.userRepo.FindByIDAndRole(p.PatientID, "patient")
	if err != nil {
		log.Printf("Error finding patient: %v", err)
		return 0, err
	}

	if validPatient == nil {
		return 0, errors.New("user id is not a patient")
	}

	// Call machine learning API to detect fundus image
	// API auth using API key
	// API endpoint: /detect
	// API method: POST
	// API body: { "image_blob": "xxx" }
	// API response: []{ "disease": "xxx", "confidence_score": 0.0, "description": "xxx" }
	fundusDetails := []*entities.FundusDetail{
		{
			Disease:         "DR",
			ConfidenceScore: 33.2,
			Description:     "",
		}, {
			Disease:         "CT",
			ConfidenceScore: 20.1,
			Description:     "",
		}, {
			Disease:         "GL",
			ConfidenceScore: 10.5,
			Description:     "",
		},
	}

	// If error, return error

	// Store image in VM
	imageURL, err := helpers.StoreImage(p.ImageBlob)
	if err != nil {
		return 0, errors.New("failed to store image")
	}

	condition := helpers.GetFundusCondition(fundusDetails)

	// Create fundus record in database
	fundus := &entities.Fundus{
		PatientID: p.PatientID,
		ImageURL:  imageURL,
		Verified:  false,
		Condition: condition,
	}

	fundusID, err := u.fundusRepo.Create(fundus, fundusDetails)
	if err != nil {
		return 0, errors.New("failed to create fundus record")
	}

	return fundusID, nil
}

func (u *FundusUsecase) ViewFundus(fundusID int64) (*entities.Fundus, error) {
	fundus, err := u.fundusRepo.FindByID(fundusID)
	if err != nil {
		return nil, errors.New("failed to find fundus record")
	}

	return fundus, nil
}

func (u *FundusUsecase) FundusHistory(userID int64) ([]*entities.Fundus, error) {
	return nil, nil
}
func (u *FundusUsecase) RequestVerifyFundusByPatient() error { return nil }

func (u *FundusUsecase) VerifyFundusByDoctor(fundusID, doctorID int, status string, feedbacks []string) error {
	feedbacksEntity := []entities.FundusFeedback{}
	for _, fb := range feedbacks {
		feedback := &entities.FundusFeedback{
			FundusID: int64(fundusID),
			DoctorID: int64(doctorID),
			Notes:    fb,
		}

		feedbacksEntity = append(feedbacksEntity, *feedback)
	}

	if err := u.fundusRepo.CreateFeedback(feedbacksEntity); err != nil {
		return errors.New("failed storing feedbacks")
	}

	if err := u.fundusRepo.UpdateVerify(fundusID, doctorID, status); err != nil {
		return errors.New("failed to verify fundus")
	}

	return nil
}

func (u *FundusUsecase) DeleteFundus(fundusID int64) error {
	if err := u.fundusRepo.Delete(fundusID); err != nil {
		return err
	}
	return nil
}
