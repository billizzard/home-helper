package controller

import (
	"github.com/kataras/iris/v12"
	"homeHelper/src/server/dto/menuPage"
)

func MenuList(ctx iris.Context) {
	dto := menuPage.NewMenuPage()
	ctx.JSON(dto)
	//return
	//
	//if err := ctx.View("menu/menu.jet.html"); err != nil {
	//	ctx.HTML("<h3>%s</h3>", err.Error())
	//	return
	//}
}
