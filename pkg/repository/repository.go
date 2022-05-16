package repository

import (
	"github.com/jmoiron/sqlx"
	"ipk/pkg/data"
	"ipk/pkg/data/stat"
)

type Authorisation interface {
	CreateUser(user data.User) (string, error)
	GetUser(username, password string) (data.User, error)
	GetAll() []data.User
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
	GetStat(chair int) ([]stat.ResponseStat, error)
	GetStatByTeacher(id string) ([]stat.ResponseStat, error)
	AddRow(stat stat.Stat) (int, error)
	AddResult(result []data.Block, rowId int) error
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
