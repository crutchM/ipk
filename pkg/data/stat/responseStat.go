package stat

import (
	"ipk/pkg/data"
)

type ResponseStat struct {
	Teacher    data.User    `json:"teacher"`
	Employment string       `json:"employment"`
	Expert     data.Expert  `json:"expert"`
	LessonDate string       `json:"lessonDate"`
	AnketDate  string       `json:"anketDate"`
	Blocks     []data.Block `json:"blocks"`
}
