package repos

import (
	"ipk/domain/model"
	"ipk/domain/model/stat"
)

type StatInterface interface {
	GetStat(chair int) ([]stat.ResponseStat, error)
	GetStatByTeacher(id string) ([]stat.ResponseStat, error)
	AddRow(stat stat.Stat) (int, error)
	AddResult(result []model.Block, rowId int) error
	RemoveUser(id string) error
}
