package stat

import "ipk/pkg/data"

//адская мазня говном, попытка собрать мало-мальски адекватную сущность для заполнения индивидуальной карточки преподавателя
type IndividualResult struct {
	Teacher data.User              `json:"teacher"`
	Items   []IndividualResultItem `json:"items"`
	General []Result               `json:"general"`
}

func NewIndividualResult(teacher data.User) *IndividualResult {
	return &IndividualResult{Teacher: teacher, Items: []IndividualResultItem{}, General: []Result{}}
}

type IndividualResultItem struct {
	Expert      data.Expert  `json:"expert"`
	Estimations []Estimation `json:"estimations"`
}

type Estimation struct {
	BlockName      string `json:"blockName"`
	QuestionNumber string `json:"questionNumber"`
	Answer         int    `json:"answer"`
}

type Result struct {
	Expert data.Expert `json:"expert"`
	Levels []Level     `json:"levels"`
}
type Level struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func (s *IndividualResult) getStat(expert data.Expert) {
	level1 := Level{Name: "низкий", Count: 0}
	level2 := Level{Name: "ниже среднего", Count: 0}
	level3 := Level{Name: "средний", Count: 0}
	level4 := Level{Name: "выше среднего", Count: 0}
	level5 := Level{Name: "высокий", Count: 0}
	for _, val := range s.Items {
		var tmp int
		var count int
		var res float32
		for _, v := range val.Estimations {
			if expert.Id == val.Expert.Id {
				tmp += v.Answer
				count++
			}
		}
		res = float32(tmp / count)
		if res <= 1.49 {
			level1.Count++
		} else if res <= 2.39 {
			level2.Count++
		} else if res <= 3.49 {
			level3.Count++
		} else if res <= 4.49 {
			level4.Count++
		} else if res <= 5 {
			level5.Count++
		}
	}
	var general Result
	general.Expert = expert
	general.Levels = []Level{level1, level2, level3, level4, level5}
	s.General = append(s.General, general)
}
