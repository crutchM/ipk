package data

type Stat struct {
	Id         int `json:"id" db:"id"`
	UserId     int `json:"userId" db:"user"`
	PostId     int `json:"postId" db:"post"`
	Employment int `json:"employment" db:"employment"`
	Block      int `json:"block" db:"block"`
	Question   int `json:"question" db:"question"`
	Answer     int `json:"answer" db:"answer"`
}
