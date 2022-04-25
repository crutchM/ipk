package data

type Result struct {
	Teacher     User             `json:"teacher"`
	ListsCount  int              `json:"listsCount"`
	Estimations EstimationResult `json:"estimations"`
}

type EstimationResult struct {
	EstimationId        int `json:"estimationId"`
	ExpertsListsCount   int `json:"expertsListsCount"`
	ExpertsListsPercent int `json:"expertsListsPercent"`
	LessonsCount        int `json:"lessonsCount"`
	LessonsPercent      int `json:"lessonsPercent"`
}

type ListenerStat struct {
	Listener            int               `json:"listener"`
	LessonTypeCount     []LessonTypeCount `json:"lessonTypeCount"`
	OnlyListenerCount   int               `json:"onlyListenerCount"`
	OnlyListenerPercent float32           `json:"onlyListenerPercent"`
}

type LessonTypeCount struct {
	LessonType int `json:"lessonType"`
	Count      int `json:"count"`
}
