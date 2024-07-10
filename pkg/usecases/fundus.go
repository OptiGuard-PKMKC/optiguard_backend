package usecases

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/request"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/helpers"
	repo_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/repositories/interfaces"
	usecase_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/usecases/interfaces"
)

const apiKey = "oEfC3PglPKoCg1jDa833awsnTLoCxWSjbumTypmSEbNgWAincCp00DkcFEw45JznC6Cou73GrU07VieU01ktlsckPyqlWoSU75Bf"
const apiEndpoint = "https://normal-utterly-raptor.ngrok-free.app/predict"

type RequestBody struct {
	FundusImage string `json:"fundus_image"`
}

type ResponseBody struct {
	PredictedClass string `json:"predicted_class"`
}

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

func detectFundusImageAPI(imageBlob string) (string, error) {
	// Create the request body
	requestBody, err := json.Marshal(RequestBody{FundusImage: imageBlob})
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the status code is not 200
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("non-200 status code: %d, response: %s", resp.StatusCode, string(body))
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the response body
	var responseBody ResponseBody
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	log.Printf("Predicted class: %s", responseBody.PredictedClass)

	return responseBody.PredictedClass, nil
}

func (u *FundusUsecase) DetectImage(p *request.DetectFundusImage) (*entities.Fundus, error) {
	validPatient, err := u.userRepo.FindByIDAndRole(p.PatientID, "patient")
	if err != nil {
		log.Printf("Error finding patient: %v", err)
		return nil, err
	}

	if validPatient == nil {
		return nil, errors.New("user id is not a patient")
	}

	// Call machine learning API to detect fundus image
	predictedClass, err := detectFundusImageAPI(p.FundusImage)
	if err != nil {
		return nil, err
	}

	// Store image in VM
	imagePath, err := helpers.StoreImage(p.FundusImage)
	if err != nil {
		return nil, errors.New("failed to store image")
	}

	// Create fundus record in database
	fundus := &entities.Fundus{
		PatientID: p.PatientID,
		ImagePath: imagePath,
		Verified:  false,
		Condition: predictedClass,
	}

	newFundus, err := u.fundusRepo.Create(fundus)
	if err != nil {
		return nil, err
	}

	return newFundus, nil
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
