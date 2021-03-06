package handler

import (
	"fileserver/frame"
	. "fileserver/log"
	"github.com/kataras/iris/v12"
)

func init() {
	frame.RegisterMiddleware(reqLog)
	frame.RegisterMiddleware(interceptFavicon)
}

func reqLog(ctx iris.Context) {
	Log.Info("[%s] %s %s", ctx.RemoteAddr(), ctx.Method(), ctx.FullRequestURI())
	ctx.Next()
}

func interceptFavicon(ctx iris.Context) {
	if ctx.Path() == "/favicon.ico" {
		return
	}
	ctx.Next()
}
