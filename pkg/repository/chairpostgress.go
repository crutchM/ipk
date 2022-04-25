package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"ipk"
)

type ChairPostgres struct {
	db *sqlx.DB
}

func NewChairPostgres(db *sqlx.DB) *ChairPostgres {
	return &ChairPostgres{db: db}
}

func (c *ChairPostgres) GetAllChairs() ([]ipk.Chair, error) {
	var chairs []ipk.Chair
	err := c.db.Select(&chairs, "SELECT * FROM chairs")
	if err != nil {
		return nil, err
	}

	return chairs, nil
}

func (c *ChairPostgres) CreateChair(chair ipk.Chair) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO chairs(name) VALUES($1) RETURNING id")
	row := c.db.QueryRow(query, chair.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
