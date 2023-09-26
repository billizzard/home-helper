package controller

import (
	"github.com/kataras/iris/v12"
	handler "homeHelper/src/server/handler/file"
)

func FileList(ctx iris.Context) {
	dto, err := handler.FileListHandler(ctx.Params().Get("path"))
	if err != nil {
		ctx.StopWithError(404, err)

		return
	}

	ctx.ViewData("data", dto)

	if err := ctx.View("files/list.jet.html"); err != nil {
		ctx.HTML("<h3>%s</h3>", err.Error())
		return
	}
}

func FileUpload(ctx iris.Context) {

	// Get the file from the request.
	//fmt.Println("dfdfdfdfdfdfdfd")
	//f, fh, err := ctx.FormFile("uploadfile")
	//if err != nil {
	//	ctx.StatusCode(iris.StatusInternalServerError)
	//	ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
	//	return
	//}
	//defer f.Close()
	//
	//_, err = ctx.SaveFormFile(fh, filepath.Join("./uploads", fh.Filename))
	//if err != nil {
	//	ctx.StatusCode(iris.StatusInternalServerError)
	//	ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
	//	return
	//}

	_, _, err := ctx.UploadFormFiles("./uploads")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
}
