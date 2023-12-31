package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"homeHelper/src/server/controller"
)

var FileRoutes = func(api *router.Party) {
	booksAPI := (*api).Party("/files")
	booksAPI.Get("/{path:path}", controller.FileList)
	booksAPI.Post("/upload", controller.FileUpload)
	booksAPI.Post("/download", controller.FileDownload)
	booksAPI.Post("/rename", controller.FileRename)
	booksAPI.Post("/delete", controller.FileDelete)
	booksAPI.Post("/folder/create", controller.FolderCreate)
}

var MenuRoutes = func(api *router.Party) {
	(*api).Get("/", controller.MenuList)
}

var IndexRoutes = func(api *iris.Application) {
	(*api).Get("/{path:path}", controller.MainPage)
}
