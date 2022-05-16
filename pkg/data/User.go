package data

type User struct {
	Id       string `json:"id" db:"id"`
	FullName string `json:"fullname" binding:"required" db:"fullname"`
	Login    string `json:"login"  binding:"required" db:"login"`
	Chair    int    `json:"chair"  binding:"required" db:"chair"`
	Post     int    `json:"post"  binding:"required" db:"post"`
	Password string `json:"password"  binding:"required" db:"password"`
}

func NewUser(id string, fullName string, login string, chair int, post int, password string) *User {
	return &User{Id: id, FullName: fullName, Login: login, Chair: chair, Post: post, Password: password}
}
