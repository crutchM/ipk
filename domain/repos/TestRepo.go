package repos

import "ipk/domain/model"

type TestInterface interface {
	GetTest(id int) (model.Test, error)
	CreateTest(test model.Test) (int, error)
}
