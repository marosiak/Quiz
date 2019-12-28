package models

type Game struct {
	ID 		int 	`gorm:"primary_key;auto_increment"`
	Name 	string 	`json:"name"`
}