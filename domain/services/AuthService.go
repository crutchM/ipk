package services

import "ipk/domain/model"

type Authorisation interface {
	CreateUser(user model.User) (string, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(accessToken string) (string, error)
	GetAll() []model.User
}
