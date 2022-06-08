package model

type Chair struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
