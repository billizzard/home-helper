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
}

// если разрастется перенести в отдельный контроллер file
func ShowFile(ctx iris.Context) {
	dto := handler.ShowFileHandler(ctx.Params().Get("path"))

	ctx.JSON(dto)
}

func FileUpload(ctx iris.Context) {
	path := ctx.Request().FormValue("path")
	err := handler.FileUploadHandler(path, &ctx)
	if err != nil {
		http.HttpBadRequest(ctx, err)

		return
	}

	ctx.JSON(nil)
}

func FolderCreate(ctx iris.Context) {
	path := ctx.Request().FormValue("path")
	name := ctx.Request().FormValue("name")
	err := handler.FolderCreateHandler(path, name)

	if err != nil {
		http.HttpBadRequest(ctx, err)

		return
	}

	ctx.JSON(nil)
}
