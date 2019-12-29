package models

type Game struct {
	ID 		int 	`gorm:"primary_key;auto_increment"json:"id"`
	Name 	string 	`json:"name"`
}