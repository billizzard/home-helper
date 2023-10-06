package routes

import (
	"github.com/kataras/iris/v12/core/router"
	"homeHelper/src/server/controller"
)

//var BookRoutes = func(app *iris.Application) {
//	booksAPI := app.Party("/books")
//	booksAPI.Get("/", controller.BookList)
//	// POST: http://localhost:8080/books
//	//booksAPI.Post("/", create)
//}

var FileRoutes = func(api *router.Party) {

	booksAPI := (*api).Party("/files")
	booksAPI.Get("/{path:path}", controller.FileList)
	booksAPI.Post("/upload", controller.FileUpload)
	booksAPI.Post("/folder/create", controller.FolderCreate)
	// POST: http://localhost:8080/books
	//booksAPI.Post("/", create)
}

var MenuRoutes = func(api *router.Party) {
	(*api).Get("/", controller.MenuList)
	//booksAPI.Post("/upload", controller.FileUpload)
	// POST: http://localhost:8080/books
	//booksAPI.Post("/", create)
}
