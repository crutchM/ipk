package service

import (
	"ipk/pkg/data"
	"ipk/pkg/repository"
)

type TestService struct {
	repo repository.TestInterface
}

func (t TestService) CreateTest(test data.Test) (int, error) {
	return t.repo.CreateTest(test)
}

func (t TestService) GetTest(id int) (test data.Test, err error) {
	return t.repo.GetTest(id)
}

func NewTestService(repo repository.TestInterface) *TestService {
	return &TestService{repo: repo}
}
