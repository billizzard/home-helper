package controller

import (
	"errors"
	"github.com/kataras/iris/v12"
	"homeHelper/config"
	"homeHelper/src/server/dto/file"
	"log"
	"os"
)

func FileList(ctx iris.Context) {
	//ctx.Bad
	if !checkPath(ctx.URLParam("path")) {
		ctx.StopWithError(404, errors.New("Bdfddfdfd"))
		return
	}
	// проверить будет ли идти на уровень выше если передать ..
	files, err := os.ReadDir(config.APP["USER_FILES_FOLDER"])
	if err != nil {
		log.Fatal(err)
	}

	dto := file.NewFileList("Files", ctx.URLParam("path"))

	for _, fileInfo := range files {
		dto.AddFile(fileInfo)
	}

	ctx.ViewData("data", dto)

	if err := ctx.View("files/list.jet.html"); err != nil {
		ctx.HTML("<h3>%s</h3>", err.Error())
		return
	}
}

func checkPath(path string) bool {
	return true
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
