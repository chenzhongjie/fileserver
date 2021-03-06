package frame

import (
	"context"
	. "fileserver/log"
	"fileserver/version"
	"github.com/kataras/iris/v12"
	"time"
)

type handlerInfo struct {
	Method  string
	Uri     string
	Handler []iris.Handler
}

var _FuncInfo = make([]func(*iris.Application), 0, 20)
var _Middleware = make([]iris.Handler, 0, 10)
var _RegInfo = make([]handlerInfo, 0, 50)

func RegisterFunc(f func(*iris.Application)) {
	_FuncInfo = append(_FuncInfo, f)
}
func RegisterMiddleware(m iris.Handler) {
	_Middleware = append(_Middleware, m)
}
func RegisterHandler(method, uri string, handler ...iris.Handler) {
	info := handlerInfo{method, uri, handler}
	_RegInfo = append(_RegInfo, info)
}


func Run(port string) {
	Log.Info("start %s %s ...", version.NAME, version.VERSION)

	app := iris.New()
	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		CloseAllWork()
		Log.Info("%s %s terminated.", version.NAME, version.VERSION)
		_ = app.Shutdown(ctx)
	})

	for _, f := range _FuncInfo {
		f(app)
	}

	for _, mid := range _Middleware {
		app.Use(mid)
	}

	for _, info := range _RegInfo {
		switch info.Method {
		case "Get":
			app.Get(info.Uri, info.Handler...)
		case "Post":
			app.Post(info.Uri, info.Handler...)
		}
	}

	Log.Info("%s is running on %s port ...", version.NAME, port)
	_ = app.Listen(":" + port, iris.WithoutInterruptHandler)
}

func CloseAllWork() {

}
