package data

type Test struct {
	Id     int     `json:"id" db:"id"`
	Name   string  `json:"name" db:"name"`
	Blocks []Block `json:"blocks"`
}

type Block struct {
	Id        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Questions []Question `json:"questions"`
}

type Question struct {
	Id     int    `json:"id" db:"id"`
	Number string `json:"number" db:"number"`
	Text   string `json:"text" db:"text"`
	Answer int    `json:"answer" db:"answer"`
}
