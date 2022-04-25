package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"ipk/pkg/data"
)

type TestPostgres struct {
	db *sqlx.DB
}

func (t TestPostgres) GetTest(id int) (data.Test, error) {
	return data.Test{}, nil
}

type query struct {
	linkedTable string
	firstField  string
	secondField string
	firsId      int
}

func (t *TestPostgres) insert(tableName, valueName, param string, quer query) (int, error) {
	var id int
	queryBlock := fmt.Sprintf("insert into %s(%s) values($1) returning id", tableName, valueName)
	row := t.db.QueryRow(queryBlock, param)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	q := fmt.Sprintf("insert into %s(%s, %s) values($1, $2)", quer.linkedTable, quer.firstField, quer.secondField)
	row = t.db.QueryRow(q, quer.firsId, id)
	return id, nil
}
func (t TestPostgres) CreateTest(test data.Test) (int, error) {
	var testId int
	queryTest := "insert into test(name) values($1) returning id"
	row := t.db.QueryRow(queryTest, test.Name)
	if err := row.Scan(&testId); err != nil {
		return 0, err
	}
	for _, v := range test.Blocks {
		km := query{"testBlocks", "test_id", "block_id", testId}
		blockId, err := t.insert("block", "name", v.Name, km)
		if err != nil {
			return 0, err
		}
		for _, val := range v.Questions {
			qm := query{"blockQuestions", "block_id", "question_id", blockId}
			questionId, err := t.insert("question", "text", val.Text, qm)
			if err != nil {
				return 0, err
			}
			for _, value := range val.Answers {
				am := query{"questionAnswers", "question_id", "answer_id", questionId}
				_, err := t.insert("answer", "text", value.Text, am)
				if err != nil {
					return 0, err
				}
			}
		}
	}
	return testId, nil
}

func NewTestPostgres(db *sqlx.DB) *TestPostgres {
	return &TestPostgres{db: db}
}
