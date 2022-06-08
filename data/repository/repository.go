package repository

import (
	"github.com/jmoiron/sqlx"
	"ipk/domain/repos"
)

type Repository struct {
	repos.Authorisation
	repos.ChairInterface
	repos.TestInterface
	repos.StatInterface
	repos.GetterInterface
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation:   NewAuthPostgres(db),
		ChairInterface:  NewChairPostgres(db),
		TestInterface:   NewTestPostgres(db),
		StatInterface:   NewStatPostgres(db),
		GetterInterface: NewGetterPostgres(db),
	}
}
