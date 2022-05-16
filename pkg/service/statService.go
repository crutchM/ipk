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

type FormData struct {
	Expert data.Expert
	Levels []Level
}

type Level struct {
	Name  string
	Value float32
}

func (s StatService) FormStat(id string) {
	res, err := s.repo.GetStatByTeacher(id)
	if err != nil {
		return
	}
	var zav []stat.ResponseStat
	var vzaimo []stat.ResponseStat
	var self []stat.ResponseStat
	for _, val := range res {
		if val.Expert.Id == 1 {
			zav = append(zav, val)
		} else if val.Expert.Id == 2 {
			vzaimo = append(vzaimo, val)
		} else {
			self = append(self, val)
		}
	}
}

func NewStatService(repo repository.StatInterface) *StatService {
	return &StatService{repo: repo}
}
