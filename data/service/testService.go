package service

import (
	"ipk/domain/model"
	"ipk/domain/repos"
)

type TestService struct {
	repo repos.TestInterface
}

func (t TestService) CreateTest(test model.Test) (int, error) {
	return t.repo.CreateTest(test)
}

func (t TestService) GetTest(id int) (test model.Test, err error) {
	return t.repo.GetTest(id)
}

func NewTestService(repo repos.TestInterface) *TestService {
	return &TestService{repo: repo}
}
