package stat

import (
	"ipk/presentation/model"
)

type ResponseStat struct {
	Teacher    model.User    `json:"teacher"`
	Employment string        `json:"employment"`
	Expert     model.Expert  `json:"expert"`
	LessonDate string        `json:"lessonDate"`
	AnketDate  string        `json:"anketDate"`
	Blocks     []model.Block `json:"blocks"`
}
