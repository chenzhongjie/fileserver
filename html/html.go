package html

import (
	"fileserver/frame"
	"github.com/kataras/iris/v12"
)

func init() {
	frame.RegisterFunc(htmlFunc)
}

func htmlFunc(app *iris.Application) {
	app.RegisterView(iris.HTML("html", ".html"))
}