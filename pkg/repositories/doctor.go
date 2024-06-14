package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/OptiGuard-PKMKC/optiguard_backend/internal/interfaces/request"
	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"
	repo_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/repositories/interfaces"
)

type DbDoctorRepository struct {
	DB *sql.DB
}

func NewDbDoctorRepository(db *sql.DB) repo_intf.DoctorRepository {
	return &DbDoctorRepository{DB: db}
}

func (r *DbDoctorRepository) FindAll(filter *request.FilterAppointmentSchedule) ([]*entities.DoctorProfile, error) {
	var query string

	query = `SELECT pr.*, sc.* FROM doctor_profiles pr JOIN doctor_schedules sc ON pr.id = sc.doctor_id`

	conditions := []string{}
	paramLen := 1
	params := []interface{}{}

	if filter != nil {
		if filter.DaysInt != nil && len(filter.DaysInt) > 0 {
			placeholders := []string{}
			for _, day := range filter.DaysInt {
				placeholders = append(placeholders, fmt.Sprintf("$%d", paramLen))
				params = append(params, day)
			}
			conditions = append(conditions, fmt.Sprintf("sc.day IN (%s)", strings.Join(placeholders, ",")))
		}

		if filter.StartHour != "" {
			paramLen++
			conditions = append(conditions, fmt.Sprintf("sc.start_hour >= $%d", paramLen))
			params = append(params, filter.StartHour)
		}

		if filter.EndHour != "" {
			paramLen++
			conditions = append(conditions, fmt.Sprintf("sc.end_hour <= $%d", paramLen))
			params = append(params, filter.EndHour)
		}

		if len(conditions) > 0 {
			query = fmt.Sprintf("%s WHERE %s", query, strings.Join(conditions, " AND "))
		}
	}

	rows, err := r.DB.Query(query, params...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	profiles := []*entities.DoctorProfile{}
	for rows.Next() {
		profile := &entities.DoctorProfile{}
		if err := rows.Scan(&profile.ID, &profile.UserID, &profile.STRNo, &profile.BioDesc); err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

func (r *DbDoctorRepository) GetProfileByID(profileID int64) (*entities.DoctorProfile, error) {
	query := `SELECT * FROM doctor_profiles WHERE id = $1`

	profile := &entities.DoctorProfile{}
	if err := r.DB.QueryRow(query, profileID).Scan(&profile.ID, &profile.UserID, &profile.STRNo, &profile.BioDesc); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	practices, err := r.GetPractice(profileID)
	if err != nil {
		return nil, err
	}
	if practices != nil {
		profile.Practices = practices
	}

	schedules, err := r.GetSchedule(profileID)
	if err != nil {
		return nil, err
	}
	if schedules != nil {
		profile.Schedules = schedules
	}

	return profile, nil
}

func (r *DbDoctorRepository) GetPractice(profileID int64) ([]*entities.DoctorPractice, error) {
	query := `SELECT * FROM doctor_practices WHERE doctor_id = $1`

	rows, err := r.DB.Query(query, profileID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	var practices []*entities.DoctorPractice
	for rows.Next() {
		practice := &entities.DoctorPractice{}
		if err := rows.Scan(&practice.ID, &practice.DoctorID, &practice.City, &practice.Province, &practice.OfficeName, &practice.StartDate, &practice.EndDate); err != nil {
			return nil, err
		}
		practices = append(practices, practice)
	}

	return practices, nil
}

func (r *DbDoctorRepository) GetSchedule(profileID int64) ([]*entities.DoctorSchedule, error) {
	query := `SELECT * FROM doctor_schedules WHERE doctor_id = $1`

	rows, err := r.DB.Query(query, profileID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	var schedules []*entities.DoctorSchedule
	for rows.Next() {
		schedule := entities.DoctorSchedule{}
		if err := rows.Scan(&schedule.ID, &schedule.DoctorID, &schedule.Day, &schedule.StartHour, &schedule.EndHour, &schedule.CreatedAt, &schedule.UpdatedAt); err != nil {
			return nil, err
		}
		schedules = append(schedules, &schedule)
	}

	return schedules, nil
}

func (r *DbDoctorRepository) CreateSchedule(schedules []*entities.DoctorSchedule) error {
	query := `INSERT INTO doctor_schedules (doctor_id, day, start_hour, end_hour) VALUES ($1, $2, $3, $4)`

	for _, schedule := range schedules {
		_, err := r.DB.Exec(query, schedule.DoctorID, schedule.Day, schedule.StartHour, schedule.EndHour)
		if err != nil {
			return err
		}
	}

	return nil
}
