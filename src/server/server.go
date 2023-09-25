package server

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"homeHelper/src/routes"
)

func RunApi(port string) error {
	fmt.Printf("Service 'api' on address localhost:%s is running... \n", port)

	app := initApp()
	initRoutes(app)

	err := app.Listen(":" + port)
	if err != nil {
		return err
	}

	return nil
}

func initApp() *iris.Application {
	app := iris.New()
	app.HandleDir("/assets", "./src/public/assets")
	app.Use(iris.Compression)
	tmpl := iris.Jet("./src/public/template", ".jet.html")
	tmpl.Reload(true) // reload templates on each request (development mode)
	app.RegisterView(tmpl)

	return app
}

func initRoutes(app *iris.Application) {
	//routes.BookRoutes(app)
	routes.FileRoutes(app)
	routes.MenuRoutes(app)
}
