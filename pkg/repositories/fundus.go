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

func (r *DbFundusRepository) CreateFeedback() error     { return nil }
func (r *DbFundusRepository) FindAll() error            { return nil }
func (r *DbFundusRepository) FindByID() error           { return nil }
func (r *DbFundusRepository) FindByIDVerified() error   { return nil }
func (r *DbFundusRepository) Delete() error             { return nil }
func (r *DbFundusRepository) DeleteFeedback() error     { return nil }
func (r *DbFundusRepository) UpdateVerifyDoctor() error { return nil }
