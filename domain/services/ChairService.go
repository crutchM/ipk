package services

import "ipk/domain/model"

type ChairInterface interface {
	GetAllChairs() ([]model.Chair, error)
	CreateChair(chair model.Chair) (int, error)
}
