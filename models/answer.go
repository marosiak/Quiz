package models

type Answer struct {
	ID 		int 	`gorm:"primary_key;auto_increment"`
	QuestionFk 	int 	`json:"question_fk"`
}