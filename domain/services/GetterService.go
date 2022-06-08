package services

import "ipk/domain/model"

type GetterInterface interface {
	GetAllTeachers() []model.User
	GetAllExperts() []model.Expert
	GetEmployments() []model.LessonType
	GetUser(id string) model.User
}
