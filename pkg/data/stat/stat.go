package stat

import "time"

type Stat struct {
	Id         int       `json:"id" db:"id"`
	UserId     string    `json:"userId" db:"userI"`
	PostId     int       `json:"postId" db:"post"`
	Employment int       `json:"employment" db:"employment"`
	Block      int       `json:"block" db:"block"`
	Question   int       `json:"question" db:"question"`
	Answer     string    `json:"answer" db:"answer"`
	Expert     int       `json:"expert" db:"expert"`
	LessonDate time.Time `json:"lessonDate" db:"lessonDate"`
	AnketDate  time.Time `json:"anketDate" db:"anketDate"`
}

func NewStat() *Stat {
	return &Stat{AnketDate: time.Now().UTC()}
}
