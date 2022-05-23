package repository

import (
	"github.com/jmoiron/sqlx"
	"ipk/pkg/data"
)

type GetterPostgres struct {
	db *sqlx.DB
}

func (g GetterPostgres) GetAllTeachers() []data.User {
	var teachers []data.User
	err := g.db.Select(&teachers, "select * from users where post=3")
	if err != nil {
		return nil
	}
	return teachers
}

func (g GetterPostgres) GetAllExperts() []data.Expert {
	var experts []data.Expert
	err := g.db.Select(&experts, "select * from expert")
	if err != nil {
		return nil
	}
	return experts
}

func (g GetterPostgres) GetEmployments() []data.LessonType {
	var employments []data.LessonType
	err := g.db.Select(&employments, "select * from lessontype")
	if err != nil {
		return nil
	}
	return employments
}

func (g GetterPostgres) GetUser(id string) data.User {
	var user data.User
	err := g.db.Get(&user, "select * from users where id=$1", id)
	if err != nil {
		return data.User{}
	}
	return user
}

func NewGetterPostgres(db *sqlx.DB) *GetterPostgres {
	return &GetterPostgres{db: db}
}
