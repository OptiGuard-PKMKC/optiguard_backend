package repositories

import (
	"database/sql"

	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"
	repo_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/repositories/interfaces"
)

type DbDoctorRepository struct {
	DB *sql.DB
}

func NewDbDoctorRepository(db *sql.DB) repo_intf.DoctorRepository {
	return &DbDoctorRepository{DB: db}
}

func (r *DbDoctorRepository) CreateSchedule(schedules []*entities.DoctorSchedule) error {
	query := `INSERT INTO doctor_schedules (doctor_id, day, start_hour, end_hour) VALUES ($1, $2, $3, $4)`

	for _, schedule := range schedules {
		_, err := r.DB.Exec(query, schedule.ProfileID, schedule.Day, schedule.StartHour, schedule.EndHour)
		if err != nil {
			return err
		}
	}

	return nil
}
