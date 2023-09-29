package controller

import (
	"errors"
	"github.com/kataras/iris/v12"
	handler "homeHelper/src/server/handler/file"
	"homeHelper/src/server/services"
	"homeHelper/src/server/services/http"
)

func FileList(ctx iris.Context) {
	isDir, err := services.IsDir(ctx.Params().Get("path"))
	if err != nil {
		http.HttpNotFound(ctx, errors.New("Page not found"))

		return
	}

	if !isDir {
		ShowFile(ctx)

		return
	}

	dto, err := handler.FileListHandler(ctx.Params().Get("path"))
	if err != nil {
		http.HttpNotFound(ctx, errors.New("Page not found"))

		return
	}

	ctx.JSON(dto)

	//ctx.ViewData("data", dto)
	//if err := ctx.View("files/list.jet.html"); err != nil {
	//	ctx.HTML("<h3>%s</h3>", err.Error())
	//	return
	//}
}

// если разрастется перенести в отдельный контроллер file
func ShowFile(ctx iris.Context) {
	ctx.HTML("<h3>%s</h3>", "File: "+ctx.Params().Get("path"))
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

	_, _, err := ctx.UploadFormFiles("./user/files")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
}
