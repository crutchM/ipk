package repository

import (
	"github.com/jmoiron/sqlx"
	"ipk/domain/model"
)

type GetterPostgres struct {
	db *sqlx.DB
}

func (g GetterPostgres) GetAllTeachers() []model.User {
	var teachers []model.User
	err := g.db.Select(&teachers, "select * from users where post=3")
	if err != nil {
		return nil
	}
	return teachers
}

func (g GetterPostgres) GetAllExperts() []model.Expert {
	var experts []model.Expert
	err := g.db.Select(&experts, "select * from expert")
	if err != nil {
		return nil
	}
	return experts
}

func (g GetterPostgres) GetEmployments() []model.LessonType {
	var employments []model.LessonType
	err := g.db.Select(&employments, "select * from lessontype")
	if err != nil {
		return nil
	}
	return employments
}

func (g GetterPostgres) GetUser(id string) model.User {
	var user model.User
	err := g.db.Get(&user, "select * from users where id=$1", id)
	if err != nil {
		return model.User{}
	}
	return user
}

func NewGetterPostgres(db *sqlx.DB) *GetterPostgres {
	return &GetterPostgres{db: db}
}
