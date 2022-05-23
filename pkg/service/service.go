package service

import (
	"ipk/pkg/data"
	"ipk/pkg/data/stat"
	"ipk/pkg/repository"
)

type Authorisation interface {
	CreateUser(user data.User) (string, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(accessToken string) (string, error)
	GetAll() []data.User
}

type ChairInterface interface {
	GetAllChairs() ([]data.Chair, error)
	CreateChair(chair data.Chair) (int, error)
}

type TestInterface interface {
	CreateTest(test data.Test) (int, error)
	GetTest(id int) (test data.Test, err error)
}
type StatInterface interface {
	GetStat(chair int) ([]stat.ResponseStat, error)
	GetStatByTeacher(id string) ([]stat.ResponseStat, error)
	AddRow(stat stat.Stat) (int, error)
	AddResult(result []data.Block, rowId int) error
}

type GetterInterface interface {
	GetAllTeachers() []data.User
	GetAllExperts() []data.Expert
	GetEmployments() []data.LessonType
	GetUser(id string) data.User
}

type Service struct {
	Authorisation
	ChairInterface
	TestInterface
	StatInterface
	GetterInterface
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
