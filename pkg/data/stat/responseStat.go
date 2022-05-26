package stat

import (
	"ipk/pkg/data"
	"time"
)

type ResponseStat struct {
	Teacher    data.User    `json:"teacher"`
	Employment string       `json:"employment"`
	Expert     data.Expert  `json:"expert"`
	LessonDate time.Time    `json:"lessonDate"`
	AnketDate  time.Time    `json:"anketDate"`
	Blocks     []data.Block `json:"blocks"`
}
