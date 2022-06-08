package service

import (
	"ipk/domain/model"
	"ipk/domain/repos"
	"ipk/domain/services"
)

type ChairService struct {
	repo services.ChairInterface
}

func NewChairService(repo repos.ChairInterface) *ChairService {
	return &ChairService{repo: repo}
}

func (s *ChairService) GetAllChairs() ([]model.Chair, error) {
	chairs, err := s.repo.GetAllChairs()
	if err != nil {
		return nil, err
	}
	return chairs, nil
}

func (s *ChairService) CreateChair(chair model.Chair) (int, error) {
	return s.repo.CreateChair(chair)
}
