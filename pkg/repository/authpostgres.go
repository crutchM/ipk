package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"ipk/pkg/data"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

//добавляем юзера в базу
func (r *AuthPostgres) CreateUser(user data.User) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (id, fullname, login, chair, post, password) values($1,$2,$3,$4,$5,$6) RETURNING id", usersTable)
	row := r.db.QueryRow(query, uuid.New().String()[:8], user.FullName, user.Login, user.Chair, user.Post, user.Password)
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

//получаем юзера из базы
func (r *AuthPostgres) GetUser(username, password string) (data.User, error) {
	var user data.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE login=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}

func (r *AuthPostgres) GetAll() []data.User {
	var users []data.User
	err := r.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil
	}
	return users
}
