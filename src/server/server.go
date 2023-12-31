package server

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"homeHelper/src/routes"
	"homeHelper/src/server/services/front"
)

func RunApi(port string) error {
	fmt.Printf("Service 'api' on address localhost:%s is running... \n", port)

	err := front.CreateConfigFile()
	if err != nil {
		return err
	}

	app := initApp()
	initRoutes(app)
	initCors(app)

	err = app.Listen(":" + port)
	if err != nil {
		return err
	}

	return nil
}

func initApp() *iris.Application {
	app := iris.New()
	app.HandleDir("/assets", "./src/public/assets")
	app.HandleDir("/public", "./user/files")
	app.Use(iris.Compression)

	tmpl := iris.HTML("./src/public/template", ".html")
	tmpl.Reload(true) // reload templates on each request (development mode)
	app.RegisterView(tmpl)
	return app
}

func initRoutes(app *iris.Application) {
	api := app.Party("/api/v1")
	routes.FileRoutes(&api)
	routes.MenuRoutes(&api)
	routes.IndexRoutes(app)
}

func initCors(app *iris.Application) {
	crs := func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		if ctx.Method() == iris.MethodOptions {
			ctx.Header("Access-Control-Methods", "POST, PUT, PATCH, DELETE")
			ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Content-Type")
			ctx.Header("Access-Control-Max-Age", "86400")
			ctx.StatusCode(iris.StatusNoContent)
			return
		}

		ctx.Next()
	}
	app.UseRouter(crs)
}
