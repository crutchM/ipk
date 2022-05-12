package repository

import (
	"github.com/jmoiron/sqlx"
	"ipk/pkg/data"
	"ipk/pkg/data/stat"
)

type Authorisation interface {
	CreateUser(user data.User) (string, error)
	GetUser(username, password string) (data.User, error)
}

type ChairInterface interface {
	GetAllChairs() ([]data.Chair, error)
	CreateChair(chair data.Chair) (int, error)
}

type TestInterface interface {
	GetTest(id int) (data.Test, error)
	CreateTest(test data.Test) (int, error)
}

type StatInterface interface {
	GetStat() ([]stat.ResponseStat, error)
	GetStatByTeacher(id int) ([]stat.ResponseStat, error)
	AddRow(stat stat.Stat) error
}

type Repository struct {
	Authorisation
	ChairInterface
	TestInterface
	StatInterface
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation:  NewAuthPostgres(db),
		ChairInterface: NewChairPostgres(db),
		TestInterface:  NewTestPostgres(db),
		StatInterface:  NewStatPostgres(db),
	}
}
