package service

import (
	"ipk/data/repository"
	"ipk/domain/services"
)

type Service struct {
	services.Authorisation
	services.ChairInterface
	services.TestInterface
	services.StatInterface
	services.GetterInterface
}

//такая реализация di, чего не поделаешь ради пародии на клин код
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorisation:   NewAuthService(repos.Authorisation),
		ChairInterface:  NewChairService(repos.ChairInterface),
		TestInterface:   NewTestService(repos.TestInterface),
		StatInterface:   NewStatService(repos.StatInterface),
		GetterInterface: NewGetterService(repos.GetterInterface),
	}

}
