package main

import (
	"Quiz/models"
	"Quiz/views"
	"github.com/buaazp/fasthttprouter"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/valyala/fasthttp"
	"log"
)
func main() {
	router := fasthttprouter.New()

	db, err := gorm.Open("mysql", "root:password@(db)/application")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	viewsSrv := views.NewViewsService(db)
	register(router, viewsSrv)

	db.AutoMigrate(&models.Game{})
	db.AutoMigrate(&models.Question{})
	db.AutoMigrate(&models.Answer{})
	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}


func register(router *fasthttprouter.Router, viewsSrv *views.InternalViewsService) {
	router.GET("/games", viewsSrv.GetGamesList)
	router.GET("/games/:id", viewsSrv.GetGame)
}