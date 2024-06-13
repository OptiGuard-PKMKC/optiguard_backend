package repositories

import (
	"database/sql"

	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"
	repo_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/repositories/interfaces"
)

type DbAppointmentRepository struct {
	DB *sql.DB
}

func NewDbAppointmentRepository(db *sql.DB) repo_intf.AppointmentRepository {
	return &DbAppointmentRepository{DB: db}
}

func (r *DbAppointmentRepository) Create(apt *entities.Appointment) error {
	query := `INSERT INTO appointments (patient_id, doctor_id, date, start_hour, end_hour, status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	var id int64
	err := r.DB.QueryRow(query, apt.PatientID, apt.DoctorID, apt.Date, apt.StartHour, apt.EndHour, apt.Status).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func (r *DbAppointmentRepository) FindAll(doctorID *int64, patientID *int64) ([]*entities.Appointment, error) {
	var query string

	if patientID != nil && doctorID == nil {
		query = `SELECT * FROM appointments WHERE patient_id = $1`
	} else if patientID == nil && doctorID != nil {
		query = `SELECT * FROM appointments WHERE doctor_id = $1`
	}

	rows, err := r.DB.Query(query, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apts []*entities.Appointment
	for rows.Next() {
		apt := &entities.Appointment{}
		if err = rows.Scan(&apt.ID, &apt.PatientID, &apt.DoctorID, &apt.Date, &apt.StartHour, &apt.EndHour, &apt.Status, &apt.CreatedAt, &apt.UpdatedAt); err != nil {
			return nil, err
		}

		apts = append(apts, apt)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return apts, nil
}

func (r *DbAppointmentRepository) UpdateStatus(aptID int64, status string) error {
	query := `UPDATE appointments SET status = $1 WHERE id = $2`

	_, err := r.DB.Exec(query, status, aptID)
	if err != nil {
		return err
	}

	return nil
}

func (r *DbAppointmentRepository) Delete(id int64) error {
	query := `DELETE FROM appointments WHERE id = $1`

	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
