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
func (s StatPostgres) formResponse(statFromDb []stat.Stat) ([]stat.ResponseStat, error) {
	var item stat.ResponseStat
	var result []stat.ResponseStat
	for _, v := range statFromDb {
		var user data.User
		err := s.db.Get(&user, "select * from users where id=$1", v.UserId)
		if err != nil {
			return nil, err
		}
		item.Teacher = user
		var emp string
		err = s.db.Get(&emp, "selet name from employment where id=$1", v.Employment)
		if err != nil {
			return nil, err
		}
		item.Employment = emp
	}
}

func (s StatPostgres) getAllAnketForUser(id int) ([]data.Block, error) {
	var result []data.Block
	var ids []int
	err := s.db.Select(&ids, "select block from stat where userI=$1 group by id order by id", id)
	if err != nil {
		return nil, err
	}
	for _, v := range ids {

		err = s.db.Select(&qIds, "select question, answer from stat where block=$1", v)
	}
}

func (s StatPostgres) GetStat() ([]stat.ResponseStat, error) {
	var static []stat.Stat
	err := s.db.Select(&static, "select * from Stat")
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s StatPostgres) GetStatByTeacher(id int) ([]stat.ResponseStat, error) {
	var stat []stat.Stat
	err := s.db.Select(&stat, "select * form Stat where userI=$1", id)
	if err != nil {
		return nil, err
	}
	return stat, nil
}

func (s StatPostgres) AddRow(stat stat.Stat) error {
	return nil
}
