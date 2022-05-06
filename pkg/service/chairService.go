package service

import (
	"ipk/pkg/data"
	"ipk/pkg/repository"
)

type ChairService struct {
	repo ChairInterface
}

func NewChairService(repo repository.ChairInterface) *ChairService {
	return &ChairService{repo: repo}
}

func (s *ChairService) GetAllChairs() ([]data.Chair, error) {
	chairs, err := s.repo.GetAllChairs()
	if err != nil {
		return nil, err
	}
	return chairs, nil
}

func (s *ChairService) CreateChair(chair data.Chair) (int, error) {
	return s.repo.CreateChair(chair)
}
