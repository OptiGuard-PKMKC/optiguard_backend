package entities

import "time"

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	RoleID    int       `json:"role_id"`
	Birthdate time.Time `json:"birthdate"`
	Gender    string    `json:"gender"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRole struct {
	ID       int    `json:"id"`
	RoleName string `json:"role_name"`
}
