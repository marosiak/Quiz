package models

type Question struct {
	ID 		int 	`gorm:"primary_key;auto_increment"`
	GameFk 	int  	`json:"game_id"`
	Text 	string 	`json:"text"`
}