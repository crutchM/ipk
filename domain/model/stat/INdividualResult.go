package stat

import "ipk/presentation/model"

//адская мазня говном, попытка собрать мало-мальски адекватную сущность для заполнения индивидуальной карточки преподавателя
type IndividualResult struct {
	Teacher model.User             `json:"teacher"`
	Items   []IndividualResultItem `json:"items"`
	General []Result               `json:"general"`
}

func NewIndividualResult(teacher model.User) *IndividualResult {
	return &IndividualResult{Teacher: teacher, Items: []IndividualResultItem{}, General: []Result{}}
}

type IndividualResultItem struct {
	Expert      model.Expert `json:"expert"`
	Estimations []Estimation `json:"estimations"`
}

type Estimation struct {
	BlockName string           `json:"blockName"`
	Questions []model.Question `json:"questions"`
}

type Result struct {
	Expert model.Expert `json:"expert"`
	Levels []Level      `json:"levels"`
}
type Level struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
