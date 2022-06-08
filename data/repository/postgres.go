package repository

import (
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
)

//type Config struct {
//	Host     string
//	Port     string
//	Username string
//	Password string
//	DBName   string
//	SSLMode  string
//}

type Config struct {
	ConnectionRow string
}

//создание подключения к бд(трогать не надо)
//я знаю что есть gorm но sqlx для меня более гибкой оказалась
func NewPostgresDb(cfg Config) (*sqlx.DB, error) {
	//db, err := sqlx.Open("postgres", fmt.Sprintf("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"))
	//cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password))
	db, err := sqlx.Open("postgres", cfg.ConnectionRow)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
