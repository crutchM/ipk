package service

import (
	"ipk/domain/model"
	"ipk/domain/repos"
)

type GetterService struct {
	repo repos.GetterInterface
}

func (g GetterService) GetAllTeachers() []model.User {
	return g.repo.GetAllTeachers()
}

func (g GetterService) GetAllExperts() []model.Expert {
	return g.repo.GetAllExperts()
}

func (g GetterService) GetEmployments() []model.LessonType {
	return g.repo.GetEmployments()
}

func (g GetterService) GetUser(id string) model.User {
	return g.repo.GetUser(id)
}

func NewGetterService(repo repos.GetterInterface) *GetterService {
	return &GetterService{repo: repo}
}
