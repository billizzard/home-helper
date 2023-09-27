package routes

import (
	"github.com/kataras/iris/v12"
	"homeHelper/src/server/controller"
)

//var BookRoutes = func(app *iris.Application) {
//	booksAPI := app.Party("/books")
//	booksAPI.Get("/", controller.BookList)
//	// POST: http://localhost:8080/books
//	//booksAPI.Post("/", create)
//}

var FileRoutes = func(app *iris.Application) {
	booksAPI := app.Party("/files")
	booksAPI.Get("/{path:path}", controller.FileList).Name = "fileList"
	booksAPI.Post("/upload", controller.FileUpload)
	// POST: http://localhost:8080/books
	//booksAPI.Post("/", create)
}

var MenuRoutes = func(app *iris.Application) {
	app.Get("/", controller.MenuList)
	//booksAPI.Post("/upload", controller.FileUpload)
	// POST: http://localhost:8080/books
	//booksAPI.Post("/", create)
}
