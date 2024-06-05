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
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`

	var userID int64
	err := r.DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *DbUserRepository) FindAll() ([]entities.User, error) {
	return nil, nil
}

func (r *DbUserRepository) FindByID(id int) (*entities.User, error) {
	return nil, nil
}

func (r *DbUserRepository) FindByEmail(email string) (*entities.User, error) {
	query := `SELECT * FROM users WHERE email = $1`

	var user entities.User
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *DbUserRepository) Update(user *entities.User) (*entities.User, error) {
	return nil, nil
}

func (r *DbUserRepository) Delete(id int) error {
	return nil
}
