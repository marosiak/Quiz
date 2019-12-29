package views

import (
	"Quiz/models"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

func (srv* InternalViewsService) GetQuestionsList(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
	var game models.Game
	srv.db.Where("id = ?", id).First(&game)
	// If game doesnt exists, no point to look for questions
	if game == (models.Game{}) {
		fmt.Fprint(ctx, "Not Found")
		return
	}

	var questions []models.Question
	srv.db.Where("game_fk = ?", id).Find(&questions)
	output, err := json.Marshal(questions)

	if err != nil {
		ctx.Response.Header.SetStatusCode(404)
		fmt.Fprint(ctx, "Not Found")
		return
	}
	fmt.Fprint(ctx, string(output))
}