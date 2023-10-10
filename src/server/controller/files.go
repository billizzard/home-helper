package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kataras/iris/v12"
	handler "homeHelper/src/server/handler/file"
	"homeHelper/src/server/services"
	"homeHelper/src/server/services/http"
)

func FileList(ctx iris.Context) {
	realPath, err := services.GetFilePathByUrl(ctx.Params().Get("path"))
	if err != nil {
		http.HttpNotFound(ctx, errors.New("Page not found"))

		return
	}

	isDir, err := services.IsDir(realPath)
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
	dto, err := handler.ShowFileHandler(ctx.Params().Get("path"))
	if err != nil {
		http.HttpNotFound(ctx, err)

		return
	}

	ctx.JSON(dto)
}

func FileUpload(ctx iris.Context) {
	path := ctx.Request().FormValue("path")
	err := handler.FileUploadHandler(path, ctx)
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

func FileRename(ctx iris.Context) {
	path := ctx.Request().FormValue("path")
	name := ctx.Request().FormValue("name")
	oldName := ctx.Request().FormValue("oldName")
	err := handler.FileRenameHandler(path, name, oldName)

	if err != nil {
		http.HttpBadRequest(ctx, err)

		return
	}

	ctx.JSON(nil)
}

func FileDelete(ctx iris.Context) {
	path := ctx.Request().FormValue("path")
	items := ctx.Request().FormValue("items")
	sl := []string{}
	err := json.Unmarshal([]byte(items), &sl)
	if err != nil {
		http.HttpBadRequest(ctx, errors.New("Invalid file list"))

		return
	}

	err = handler.FileDeleteHandler(path, sl)

	if err != nil {
		http.HttpBadRequest(ctx, err)

		return
	}

	ctx.JSON(nil)
}

func FileDownload(ctx iris.Context) {
	path := ctx.Request().FormValue("path")
	items := ctx.Request().FormValue("items")
	sl := []string{}
	err := json.Unmarshal([]byte(items), &sl)
	if err != nil {
		http.HttpBadRequest(ctx, errors.New("Invalid file list"))

		return
	}

	path, err = handler.FileDownloadHandler(path, sl)

	if err != nil {
		http.HttpBadRequest(ctx, err)

		return
	}

	//ctx.ResponseWriter().Header().Set("Content-Disposition", "attachment; filename=WHATEVER_YOU_WANT")
	//ctx.ResponseWriter().Header().Set("Content-Type", "application/octet-stream")
	//fmt.Println("==============================")
	//fmt.Println(path)
	////f, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	////f, _ := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	//f, err := os.Open(path)
	//if err != nil {
	//	fmt.Println("+++++++++++++++++++")
	//	fmt.Println(err)
	//}
	//defer func(f *os.File) {
	//	err := f.Close()
	//	if err != nil {
	//		fmt.Println("ERRRR======")
	//		fmt.Println(err)
	//	}
	//}(f)
	//_, err = io.Copy(ctx.ResponseWriter(), f)
	//if err != nil {
	//	fmt.Println(err)
	//}

	err = ctx.SendFile(path, "my_file.png")
	if err != nil {
		fmt.Println(err)
	}
	//http.ServeFile(w, r, filePath)
	//ctx.ResponseWriter().Header().Set("Content-Type", ctx.Request().Header.Get("Content-Type"))
	//
	//ctx.JSON(nil)
}
