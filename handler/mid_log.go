package handler

import (
	"fileserver/frame"
	. "fileserver/log"
	"github.com/kataras/iris/v12"
)

func init() {
	frame.RegisterMiddleware(midLog)
}

func midLog(ctx iris.Context) {
	Log.Info("%s %s %s", ctx.RemoteAddr(), ctx.Method(), ctx.FullRequestURI())
	ctx.Next()
}
