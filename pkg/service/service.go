package service

import (
	"ipk/pkg/data"
	"ipk/pkg/repository"
)

type Authorisation interface {
	CreateUser(user data.User) (string, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(accessToken string) (string, error)
}

type ChairInterface interface {
	GetAllChairs() ([]data.Chair, error)
	CreateChair(chair data.Chair) (int, error)
}

type TestInterface interface {
	CreateTest(test data.Test) (int, error)
	GetTest(id int) (test data.Test, err error)
}

type Service struct {
	Authorisation
	ChairInterface
	TestInterface
}

//такая реализация di, чего не поделаешь ради пародии на клин код
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorisation:  NewAuthService(repos.Authorisation),
		ChairInterface: NewChairService(repos.ChairInterface),
		TestInterface:  NewTestService(repos.TestInterface),
	}

}
