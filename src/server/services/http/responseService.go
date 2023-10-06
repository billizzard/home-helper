package http

import (
	"github.com/kataras/iris/v12"
)

func HttpNotFound(ctx iris.Context, err error) {
	if err != nil {
		ctx.StopWithError(404, err)

		//if err := ctx.View("errors/404.jet.html"); err != nil {
		//	ctx.HTML("<h3>%s</h3>", "Service not available, try later")
		//	return
		//}
	}
}

func HttpBadRequest(ctx iris.Context, err error) {
	if err != nil {
		ctx.StopWithError(400, err)
		//ctx.JSON({
		//
		//})
		//
		//if err := ctx.View("errors/404.jet.html"); err != nil {
		//	ctx.HTML("<h3>%s</h3>", "Service not available, try later")
		//	return
		//}
	}
}

//func httpError(ctx iris.Context, code int, err error) {
//	ctx.StopWithError(code, err)
//}
