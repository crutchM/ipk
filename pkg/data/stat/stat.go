package stat

import "time"

type Stat struct {
	Id         int       `json:"id" db:"id"`
	UserId     string    `json:"userId" db:"useri"`
	PostId     int       `json:"postId" db:"post"`
	ChairId    int       `json:"chairId" db:"chair"`
	Employment int       `json:"employment" db:"employment"`
	Expert     int       `json:"expert" db:"expert"`
	LessonDate time.Time `json:"lessonDate" db:"lessondate"`
	AnketDate  time.Time `json:"anketDate" db:"anketdate"`
}

func NewStat() *Stat {
	return &Stat{AnketDate: time.Now().UTC()}
}

type TestResult struct {
	Id       int `json:"id" db:"id"`
	Block    int `json:"block" db:"block"`
	Question int `json:"question" db:"question"`
	Answer   int `json:"answer" db:"answer"`
}
