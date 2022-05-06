package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"ipk/pkg/data"
)

type TestPostgres struct {
	db *sqlx.DB
}

//get test functional-------------------------------------------------------------------------------------------
func (t TestPostgres) getBlocks(tId int) ([]data.Block, error) {
	var blocks []data.Block
	var bId []int
	if err := t.db.Select(&bId, "select block_id from testBlocks where tes_id = $1", tId); err != nil {
		return nil, err
	}
	for _, v := range bId {
		if err := t.db.Select(&blocks, "select * from block where id=$1", v); err != nil {
			return nil, err
		}
	}

	return blocks, nil
}

func (t TestPostgres) getQuestions(bId int) ([]data.Question, error) {
	var qIds []int
	var result []data.Question
	if err := t.db.Select(&qIds, "select question_id from blockQuestions where block_id=$1", bId); err != nil {
		return nil, err
	}
	for _, v := range qIds {
		var question data.Question
		if err := t.db.Get(&question, "select * from question where id=$1", v); err != nil {
			return nil, err
		}
		result = append(result, question)
	}

	return result, nil
}

func (t TestPostgres) GetTest(id int) (data.Test, error) {
	var test data.Test
	if err := t.db.Get(&test, "select * from test where id=$1", id); err != nil {
		return data.Test{}, err
	}
	blocks, err := t.getBlocks(test.Id)
	if err != nil {
		return data.Test{}, err
	}
	var result []data.Block
	for _, block := range blocks {
		if questions, err := t.getQuestions(block.Id); err != nil {
			return data.Test{}, err
		} else {
			block.Questions = questions
		}
		result = append(result, block)
	}
	test.Blocks = result
	return test, nil
}

//-------------------------------------------------------------------------------------------------------------------

//create test functional---------------------------------------------------------------------------------------------
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
			_, err := t.insert("question", "text", val.Text, qm)
			if err != nil {
				return 0, err
			}
		}
	}
	return testId, nil
}

//-------------------------------------------------------------------------------------------------------------------

// NewTestPostgres constructor
func NewTestPostgres(db *sqlx.DB) *TestPostgres {
	return &TestPostgres{db: db}
}
