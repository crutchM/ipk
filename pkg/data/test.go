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
	Id      int      `json:"id" db:"id"`
	Text    string   `json:"text" db:"text"`
	Answers []Answer `json:"answers"`
}

type Answer struct {
	Id   int    `json:"id" db:"id"`
	Text string `json:"text" db:"text"`
}
