package repository

import (
	"github.com/jmoiron/sqlx"
	"ipk"
	"ipk/pkg/data"
)

type Authorisation interface {
	CreateUser(user data.User) (string, error)
	GetUser(username, password string) (data.User, error)
}

type ChairInterface interface {
	GetAllChairs() ([]ipk.Chair, error)
	CreateChair(chair ipk.Chair) (int, error)
}

type TestInterface interface {
	GetTest(id int) (data.Test, error)
	CreateTest(test data.Test) (int, error)
}

type Repository struct {
	Authorisation
	ChairInterface
	TestInterface
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation:  NewAuthPostgres(db),
		ChairInterface: NewChairPostgres(db),
		TestInterface:  NewTestPostgres(db),
	}
}
