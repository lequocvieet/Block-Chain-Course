package models

import "time"

type UserExam struct {
	ID         int       `json:"id" gorm:"column:id;"`
	Answer     string    `json:"answer" gorm:"column:answer;"`
	UserID     int       `json:"user_id" gorm:"column:user_id;"`
	ExamCode   int       `json:"exam_code" gorm:"column:exam_code;"`
	SubmitTime time.Time `json:"submit_time" gorm:"column:submit_time;"`
	Fee        int       `json:"fee" gorm:"column:fee;"`
}
