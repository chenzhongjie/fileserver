package handler

import (
	"fileserver/frame"
	. "fileserver/log"
	"github.com/kataras/iris/v12"
)

func init() {
	frame.RegisterHandler("Post", "/page", checkPageToken, upload)
}

func checkPageToken(ctx iris.Context) {
	valid, tokenParam := isValidToken(ctx, pageToken)
	if valid {
		ctx.Next()
		return
	}
	Log.Warn("%s %s page token %s is wrong.", ctx.RemoteAddr(), ctx.FullRequestURI(), tokenParam)
	ctx.Writef("Authentication failed.")
}
