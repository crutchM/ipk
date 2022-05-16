package service

import (
	"ipk/pkg/data/stat"
	"ipk/pkg/repository"
)

type StatService struct {
	repo repository.StatInterface
}

func (s StatService) GetStat(chair int) ([]stat.ResponseStat, error) {
	return s.repo.GetStat(chair)
}

func (s StatService) GetStatByTeacher(id int) ([]stat.IndividualResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s StatService) AddRow(stat stat.Stat) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s StatService) AddResult(result stat.Result, rowId int) error {
	//TODO implement me
	panic("implement me")
}

func NewStatService(repo repository.StatInterface) *StatService {
	return &StatService{repo: repo}
}
