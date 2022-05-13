package repository

import (
	"github.com/jmoiron/sqlx"
	"ipk/pkg/data"
	"ipk/pkg/data/stat"
)

type StatPostgres struct {
	db *sqlx.DB
}

func NewStatPostgres(db *sqlx.DB) *StatPostgres {
	return &StatPostgres{db: db}
}

func (s StatPostgres) getBlocks(teacher int) ([]data.Block, error) {
	return nil, nil
}

func (s StatPostgres) getAllAnketForUser(id int) ([]data.Block, error) {
	return nil, nil
}
func (s StatPostgres) formResponse(entries []stat.Stat) ([]stat.ResponseStat, error) {
	var result []stat.ResponseStat
	for _, v := range entries {
		var item stat.ResponseStat
		teacher, err := s.getTeacher(v.UserId)
		if err != nil {
			return nil, err
		}
		item.Teacher = teacher
		emp, err := s.getEmployment(v.Employment)
		if err != nil {
			return nil, err
		}
		item.Employment = emp

	}
}
func (s StatPostgres) getBlocks() {

}

func (s StatPostgres) getEmployment(id int) (string, error) {
	var emp string
	err := s.db.Get(&emp, "select name from employment where id=$1", id)
	if err != nil {
		return "", err
	}
	return emp, nil
}

func (s StatPostgres) getTeacher(id string) (data.User, error) {
	var teacher data.User
	err := s.db.Get(&teacher, "select * from users where id=$1", id)
	if err != nil {
		return data.User{}, err
	}
	return teacher, nil
}

func (s StatPostgres) GetStat(chair int) ([]stat.ResponseStat, error) {
	var static []stat.Stat
	err := s.db.Select(&static, "select * from stat where chair=$1", chair)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s StatPostgres) GetStatByTeacher(id int) ([]stat.IndividualResult, error) {
	var static []stat.Stat
	err := s.db.Select(&static, "select * form Stat where userI=$1", id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s StatPostgres) AddRow(stat stat.Stat) error {
	return nil
}
