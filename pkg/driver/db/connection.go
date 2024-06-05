package driver_db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/OptiGuard-PKMKC/optiguard_backend/pkg/config"
	_ "github.com/lib/pq"
)

type ConfigDB struct {
	Dialect  string
	Host     string
	Port     int
	User     string
	Name     string
	Password string
	SslMode  string
}

func NewConnection(env *config.Env) (*sql.DB, error) {
	config := ConfigDB{
		Dialect:  env.DbDialect,
		Host:     env.DbHost,
		Port:     env.DbPort,
		User:     env.DbUser,
		Name:     env.DbName,
		Password: env.DbPassword,
		SslMode:  env.DbSslMode,
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", config.User, config.Password, config.Host, config.Port, config.Name, config.SslMode)

	db, err := sql.Open(config.Dialect, connString)
	if err != nil {
		log.Printf("Error on connecting to database: %v", err)
		return nil, err
	}

	// Test connection
	if err := db.Ping(); err != nil {
		log.Printf("Error on pinging to database: %v", err)
		return nil, err
	}

	log.Println("Connected to postgres database")
	return db, nil
}
