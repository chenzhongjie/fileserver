package handler

import (
	"fileserver/frame"
	. "fileserver/log"
	"github.com/kataras/iris/v12"
)

func init() {
	frame.RegisterMiddleware(reqLog)
	frame.RegisterMiddleware(getTokenParam)
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

func getTokenParam(ctx iris.Context) {
	tokenParam := ctx.GetHeader("token")
	if tokenParam == "" {
		tokenParam = ctx.URLParam("token")
	}
	Log.Debug("token param: %s", tokenParam)
	ctx.Values().Save("token", tokenParam, true)
	ctx.Next()
}
