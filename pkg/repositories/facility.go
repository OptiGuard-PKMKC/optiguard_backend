package repositories

import (
	"database/sql"

	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/request"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"
	repo_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/repositories/interfaces"
)

type DbHealthFacilityRepository struct {
	DB *sql.DB
}

func NewDbHealthFacilityRepository(db *sql.DB) repo_intf.HealthFacilityRepository {
	return &DbHealthFacilityRepository{DB: db}
}

func (r *DbHealthFacilityRepository) CreateSchedule(userID int64, p *request.CreateAdaptorSchedule) error {
	query := `INSERT INTO user_adaptors (user_id, facility_id, adaptor_id, schedule_id, date) VALUES ($1, $2, $3, $4, $5)`

	_, err := r.DB.Exec(query, userID, p.FacilityID, p.AdaptorID, p.ScheduleID, p.Date.Time)
	if err != nil {
		return err
	}

	return nil
}

func (r *DbHealthFacilityRepository) FindAll() ([]*entities.HealthFacility, error) {
	query := `SELECT * FROM health_facilities`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var facilities []*entities.HealthFacility
	for rows.Next() {
		fc := &entities.HealthFacility{}

		if err := rows.Scan(&fc.ID, &fc.FacilityName, &fc.City, &fc.Province, &fc.Address, &fc.AdaptorQuantity); err != nil {
			return nil, err
		}
	}

	return facilities, nil
}

func (r *DbHealthFacilityRepository) FindByID(id int64) (*entities.HealthFacility, error) {
	query := `SELECT * FROM health_facilities WHERE id = $1 RETURNING id`

	var facility entities.HealthFacility
	if err := r.DB.QueryRow(query, id).Scan(&facility.ID, &facility.City, &facility.Province, &facility.Address, &facility.AdaptorQuantity); err != nil {
		return nil, err
	}

	adaptors, err := r.FindAdaptorsByFacilityID(facility.ID)
	if err != nil {
		return nil, err
	}
	facility.Adaptors = adaptors

	return &facility, nil
}

func (r *DbHealthFacilityRepository) FindAdaptorsByFacilityID(facilityID int64) ([]entities.Adaptor, error) {
	query := `SELECT * FROM adaptors WHERE facility_id = $1`

	rows, err := r.DB.Query(query, facilityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var adaptors []entities.Adaptor
	for rows.Next() {
		adaptor := &entities.Adaptor{}

		if err := rows.Scan(&adaptor.ID, &adaptor.FacilityID, &adaptor.DeviceCode, &adaptor.Used); err != nil {
			return nil, err
		}
		adaptors = append(adaptors, *adaptor)
	}

	return adaptors, nil
}
