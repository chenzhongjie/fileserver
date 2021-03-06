package handler

import (
	"fileserver/frame"
	. "fileserver/log"
	"github.com/kataras/iris/v12"
)

func init() {
	frame.RegisterMiddleware(mid)
}

func mid(ctx iris.Context) {
	Log.Info("%s %s %s", ctx.RemoteAddr(), ctx.Method(), ctx.Path())
	ctx.Next()
}
