package views

import (
	"Quiz/models"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

func (srv* InternalViewsService) GetQuestionsList(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
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