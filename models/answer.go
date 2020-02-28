package models

type Answer struct {
	ID 			int 	`gorm:"primary_key;auto_increment" json:"id"`
	QuestionFk 	int 	`json:"question_fk"`
	IsCorrect	bool	`json:"is_correct"`
}