package service

import "ipk/pkg/data"

type GetterService struct {
	repo GetterInterface
}

func (g GetterService) GetAllTeachers() []data.User {
	return g.repo.GetAllTeachers()
}

func (g GetterService) GetAllExperts() []data.Expert {
	return g.repo.GetAllExperts()
}

func (g GetterService) GetEmployments() []data.LessonType {
	return g.repo.GetEmployments()
}

func (g GetterService) GetUser(id string) data.User {
	return g.repo.GetUser(id)
}

func NewGetterService(repo GetterInterface) *GetterService {
	return &GetterService{repo: repo}
}
