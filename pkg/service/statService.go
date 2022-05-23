package service

import (
	"ipk/pkg/data"
	"ipk/pkg/data/stat"
	"ipk/pkg/repository"
)

type StatService struct {
	repo repository.StatInterface
}

func (s StatService) GetStat(chair int) ([]stat.ResponseStat, error) {
	return s.repo.GetStat(chair)
}

func (s StatService) GetStatByTeacher(id string) ([]stat.ResponseStat, error) {
	return s.repo.GetStatByTeacher(id)
}

func (s StatService) AddRow(stat stat.Stat) (int, error) {
	return s.repo.AddRow(stat)
}

func (s StatService) AddResult(result []data.Block, rowId int) error {
	return s.repo.AddResult(result, rowId)
}

func NewStatService(repo repository.StatInterface) *StatService {
	return &StatService{repo: repo}
}
