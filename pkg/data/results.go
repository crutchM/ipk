package data

type Results struct {
	User        User `json:"user"`
	PapersCount int  `json:"papersCount"`
	FormTwo
	LessonType LessonType `json:"lessonType"`
}

type FormOne struct {
}

type FormTwo struct {
	Category          Post               `json:"category"`
	LessonTypeResults []LessonTypeResult `json:"lessonTypeResults"`
}

type LessonTypeResult struct {
	LessonType       LessonType `json:"lessonTypes"`
	EstimationsCount int        `json:"estimationsCount"`
}

type FormComponent struct {
}

type ComponentResult struct {
}
