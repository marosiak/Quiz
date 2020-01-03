package views

import "github.com/jinzhu/gorm"

type ViewsService interface {
	GetGames()
	GetGame()

	GetQuestions()
	GetQuestionsByID()
}

type InternalViewsService struct {
	db *gorm.DB
}

func NewViewsService(db *gorm.DB) *InternalViewsService {
	return &InternalViewsService{db: db}
}

