package controller

import (
	"github.com/kataras/iris/v12"
)

func MenuList(ctx iris.Context) {
	if err := ctx.View("menu/menu.jet.html"); err != nil {
		ctx.HTML("<h3>%s</h3>", err.Error())
		return
	}
}
