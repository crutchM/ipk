package services

import "ipk/domain/model"

type TestInterface interface {
	CreateTest(test model.Test) (int, error)
	GetTest(id int) (test model.Test, err error)
}
