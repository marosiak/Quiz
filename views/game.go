package views

import (
	"Quiz/models"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

func (srv* InternalViewsService) GetGamesList(ctx *fasthttp.RequestCtx) {
	var games []models.Game
	srv.db.Find(&games)
	output, err := json.Marshal(games)

	if err != nil {
		ctx.Response.Header.SetStatusCode(404)
		fmt.Fprint(ctx, "Not Found")
		return
	}
	fmt.Fprint(ctx, string(output))
}

func (srv* InternalViewsService) GetGame(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
	var game models.Game
	srv.db.Where("id = ?", id).First(&game)
	output, err := json.Marshal(game)

	if err != nil || game == (models.Game{}){
		ctx.Response.Header.SetStatusCode(404)
		fmt.Fprint(ctx, "Not Found")
		return
	}
	fmt.Fprint(ctx, string(output))
}