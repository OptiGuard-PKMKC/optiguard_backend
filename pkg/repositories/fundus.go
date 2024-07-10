package repositories

import (
	"database/sql"

	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"
	repo_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/repositories/interfaces"
)

type DbFundusRepository struct {
	DB *sql.DB
}

func NewDbFundusRepository(db *sql.DB) repo_intf.FundusRepository {
	return &DbFundusRepository{DB: db}
}

func (r *DbFundusRepository) Create(fundus *entities.Fundus) (*entities.Fundus, error) {
	query := `INSERT INTO funduses (patient_id, image_path, verified, status, condition) VALUES ($1, $2, $3, $4, $5) RETURNING id, patient_id, image_path, verified, status, condition, created_at`

	var newFundus entities.Fundus
	if err := r.DB.QueryRow(query, fundus.PatientID, fundus.ImagePath, fundus.Verified, fundus.Status, fundus.Condition).Scan(&newFundus.ID, &newFundus.PatientID, &newFundus.ImagePath, &newFundus.Verified, &newFundus.Status, &newFundus.Condition, &newFundus.CreatedAt); err != nil {
		return nil, err
	}

	return &newFundus, nil
}

func (r *DbFundusRepository) CreateFeedback(feedback []entities.FundusFeedback) error {
	query := `INSERT INTO fundus_feedbacks (fundus_id, doctor_id, notes) VALUES ($1, $2, $3)`

	for _, fb := range feedback {
		_, err := r.DB.Exec(query, fb.FundusID, fb.DoctorID, fb.Notes)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *DbFundusRepository) FindAll() error { return nil }

func (r *DbFundusRepository) FindByID(id int64) (*entities.Fundus, error) {
	query := `SELECT * FROM funduses WHERE id = $1`

	var fundus entities.Fundus
	err := r.DB.QueryRow(query, id).Scan(&fundus.ID, &fundus.PatientID, &fundus.ImagePath, &fundus.Verified, &fundus.Status, &fundus.Condition, &fundus.CreatedAt, &fundus.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	details, err := r.FindFundusDetails(fundus.ID)
	if err != nil {
		return nil, err
	}
	fundus.Detail = details

	feedbacks, err := r.FindFundusFeedbacks(fundus.ID)
	if err != nil {
		return nil, err
	}
	fundus.Feedback = feedbacks

	return &fundus, nil
}

func (r *DbFundusRepository) FindFundusDetails(fundusID int64) ([]*entities.FundusDetail, error) {
	query := `SELECT * FROM fundus_details WHERE fundus_id = $1`

	rows, err := r.DB.Query(query, fundusID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fundusDetails := []*entities.FundusDetail{}
	for rows.Next() {
		detail := entities.FundusDetail{}
		if err := rows.Scan(&detail.ID, &detail.FundusID, &detail.Disease, &detail.ConfidenceScore, &detail.Description); err != nil {
			return nil, err
		}
		fundusDetails = append(fundusDetails, &detail)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return fundusDetails, nil
}

func (r *DbFundusRepository) FindFundusFeedbacks(fundusID int64) ([]*entities.FundusFeedback, error) {
	query := `SELECT * FROM fundus_feedbacks WHERE fundus_id = $1`

	rows, err := r.DB.Query(query, fundusID)
	if err != nil {
		return nil, err
	}
	defer rows.Next()

	fundusFeedbacks := []*entities.FundusFeedback{}
	for rows.Next() {
		feedback := entities.FundusFeedback{}
		if err := rows.Scan(&feedback.ID, &feedback.FundusID, &feedback.DoctorID, &feedback.Notes, &feedback.CreatedAt, &feedback.UpdatedAt); err != nil {
			return nil, err
		}
		fundusFeedbacks = append(fundusFeedbacks, &feedback)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return fundusFeedbacks, nil
}

func (r *DbFundusRepository) FindByIDVerified() error { return nil }

func (r *DbFundusRepository) Delete(id int64) error {
	query := `DELETE FROM funduses WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *DbFundusRepository) DeleteFeedback() error { return nil }

func (r *DbFundusRepository) UpdateVerify(fundusID, doctorID int, status string) error {
	query := `UPDATE funduses SET verified = $1, verified_by = $2, status = $3 WHERE id = $4`

	_, err := r.DB.Exec(query, true, doctorID, status, fundusID)
	if err != nil {
		return err
	}

	return nil
}
