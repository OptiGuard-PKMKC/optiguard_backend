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

func (r *DbFundusRepository) Create(fundus *entities.Fundus, details []*entities.FundusDetail) (int64, error) {
	// Begin a transaction
	tx, err := r.DB.Begin()
	if err != nil {
		return 0, err
	}

	// Ensure the transaction is rolled back in case of error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Insert the fundus record and get the generated ID
	query := `INSERT INTO funduses (patient_id, image_url, verified, condition) VALUES ($1, $2, $3, $4) RETURNING id`
	var fundusID int64
	err = tx.QueryRow(query, fundus.PatientID, fundus.ImageURL, fundus.Verified, fundus.Condition).Scan(&fundusID)
	if err != nil {
		return 0, err
	}

	// Prepare the query for inserting fundus details
	detailQuery := `INSERT INTO fundus_details (fundus_id, disease, confidence_score, description) VALUES ($1, $2, $3, $4)`
	stmt, err := tx.Prepare(detailQuery)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Insert each fundus detail
	for _, detail := range details {
		_, err = stmt.Exec(fundusID, detail.Disease, detail.ConfidenceScore, detail.Description)
		if err != nil {
			return 0, err
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return fundusID, nil
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
	err := r.DB.QueryRow(query, id).Scan(&fundus.ID, &fundus.PatientID, &fundus.ImageURL, &fundus.Verified, &fundus.Status, &fundus.Condition, &fundus.CreatedAt, &fundus.UpdatedAt)
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
