package repositories

import (
	"database/sql"

	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/entities"
	repo_intf "github.com/OptiGuard-PKMKC/optiguard_backend/pkg/repositories/interfaces"
)

type DbUserRepository struct {
	DB *sql.DB
}

func NewDbUserRepository(db *sql.DB) repo_intf.UserRepository {
	return &DbUserRepository{
		DB: db,
	}
}

func (r *DbUserRepository) Create(user *entities.User) (int64, error) {
	query := `INSERT INTO users (name, email, phone, password, role_id, birthdate, gender, city, province, address) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

	var userID int64
	err := r.DB.QueryRow(query, user.Name, user.Email, user.Phone, user.Password, user.RoleID, user.Birthdate.Time, user.Gender, user.City, user.Province, user.Address).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *DbUserRepository) FindAll() ([]entities.User, error) {
	return nil, nil
}

func (r *DbUserRepository) FindByID(id int64) (*entities.User, error) {
	query := `SELECT u.*, ur.role_name FROM users u JOIN user_roles ur ON u.role_id = ur.id WHERE u.id = $1`

	var user entities.User
	err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.RoleID, &user.Birthdate.Time, &user.Gender, &user.City, &user.Province, &user.Address, &user.CreatedAt, &user.UpdatedAt, &user.RoleName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *DbUserRepository) FindByEmail(email string) (*entities.User, error) {
	query := `SELECT u.*, ur.role_name FROM users u JOIN user_roles ur ON u.role_id = ur.id WHERE email = $1`

	var user entities.User
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.RoleID, &user.Birthdate.Time, &user.Gender, &user.City, &user.Province, &user.Address, &user.CreatedAt, &user.UpdatedAt, &user.RoleName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *DbUserRepository) GetRoleID(roleName string) (*entities.UserRole, error) {
	query := `SELECT * FROM user_roles WHERE role_name = $1`

	var userRole entities.UserRole
	err := r.DB.QueryRow(query, roleName).Scan(&userRole.ID, &userRole.RoleName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &userRole, nil
}

func (r *DbUserRepository) Update(user *entities.User) (*entities.User, error) {
	return nil, nil
}

func (r *DbUserRepository) Delete(id int) error {
	return nil
}
