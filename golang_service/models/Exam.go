package models

type Exam struct {
	Reward   int    `json:"reward" gorm:"column:reward"`
	ExamCode int    `json:"exam_code" gorm:"primaryKey;column:exam_code"`
	Answer   string `json:"answer" gorm:"column:answer"`
	Fee      int    `json:"fee" gorm:"column:fee"`
}
