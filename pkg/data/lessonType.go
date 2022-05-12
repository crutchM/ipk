package data

type LessonType struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
