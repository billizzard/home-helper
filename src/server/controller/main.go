package controller

import (
	"github.com/kataras/iris/v12"
)

func MainPage(ctx iris.Context) {
	ctx.View("index/index.html", iris.Map{})
}
