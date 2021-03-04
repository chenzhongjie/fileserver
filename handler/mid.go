package handler

import (
	"fileserver/frame"
	"github.com/kataras/iris/v12"
)

func init() {
	frame.RegisterMiddleware(mid)
}

func mid(ctx iris.Context) {
	ctx.Next()
}
