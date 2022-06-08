package service

import (
	"ipk/domain/model"
	"ipk/domain/model/stat"
	"ipk/domain/repos"
)

type StatService struct {
	repo repos.StatInterface
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

func (s StatService) AddResult(result []model.Block, rowId int) error {
	return s.repo.AddResult(result, rowId)
}

func (s StatService) RemoveUser(id string) error {
	return s.repo.RemoveUser(id)
}

func NewStatService(repo repos.StatInterface) *StatService {
	return &StatService{repo: repo}
}
