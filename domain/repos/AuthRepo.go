package repos

import "ipk/domain/model"

type Authorisation interface {
	CreateUser(user model.User) (string, error)
	GetUser(username, password string) (model.User, error)
	GetAll() []model.User
}
