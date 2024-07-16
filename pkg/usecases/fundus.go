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

type RequestBody struct {
	FundusImage string `json:"fundus_image"`
}

type ResponseData struct {
	PredictedClass string `json:"predicted_class"`
	CroppedImage   string `json:"cropped_image"`
}

type ResponseBody struct {
	Success bool          `json:"success"`
	Message string        `json:"message,omitempty"`
	Error   string        `json:"error,omitempty"`
	Data    *ResponseData `json:"data,omitempty"`
}

type FundusUsecase struct {
	mlApi      string
	mlApiKey   string
	fundusRepo repo_intf.FundusRepository
	userRepo   repo_intf.UserRepository
}

func NewFundusUsecase(mlApi string, mlApiKey string, fundusRepo repo_intf.FundusRepository, userRepo repo_intf.UserRepository) usecase_intf.FundusUsecase {
	return &FundusUsecase{
		mlApi:      mlApi,
		mlApiKey:   mlApiKey,
		fundusRepo: fundusRepo,
		userRepo:   userRepo,
	}
}

func detectFundusImageAPI(mlApi string, mlApiKey string, imageBlob string) (*ResponseBody, error) {
	// Create the request body
	requestBody, err := json.Marshal(RequestBody{FundusImage: imageBlob})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/predict", mlApi), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", mlApiKey)

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the status code is not 200
	if resp.StatusCode > 299 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("non-200 status code: %d, response: %s", resp.StatusCode, string(body))
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the response body
	var responseBody ResponseBody
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	log.Printf("Predicted class: %s", responseBody.Data.PredictedClass)

	return &responseBody, nil
}

func (u *FundusUsecase) DetectImage(p *request.DetectFundusImage) (*entities.Fundus, *string, error) {
	validPatient, err := u.userRepo.FindByIDAndRole(p.PatientID, "patient")
	if err != nil {
		log.Printf("Error finding patient: %v", err)
		return nil, nil, err
	}

	if validPatient == nil {
		return nil, nil, errors.New("user id is not a patient")
	}

	// Call machine learning API to detect fundus image
	mlResponse, err := detectFundusImageAPI(u.mlApi, u.mlApiKey, p.FundusImage)
	if err != nil {
		return nil, nil, err
	}

	// Check if the predicted class is not fundus
	if mlResponse.Data == nil {
		return nil, &mlResponse.Message, nil
	}

	// Store image in VM
	imagePath, err := helpers.StoreImage(mlResponse.Data.CroppedImage)
	if err != nil {
		return nil, nil, errors.New("failed to store image")
	}

	// Create fundus record in database
	fundus := &entities.Fundus{
		PatientID: p.PatientID,
		ImagePath: imagePath,
		Verified:  false,
		Condition: mlResponse.Data.PredictedClass,
	}

	newFundus, err := u.fundusRepo.Create(fundus)
	if err != nil {
		return nil, nil, err
	}
	newFundus.ImageBlob = mlResponse.Data.CroppedImage

	return newFundus, nil, nil
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
