package handler

import (
	"fileserver/frame"
	. "fileserver/log"
	"fileserver/utils"
	"github.com/kataras/iris/v12"
)

var pageUploadIp = "127.0.0.1"

func init() {
	internalIp, err := utils.GetInternalIp()
	if err == nil {
		pageUploadIp = internalIp
		Log.Debug("get internal ip: %s", pageUploadIp)
	}
	frame.RegisterFunc(htmlFunc)
	frame.RegisterHandler("Get", "/", pageGet)
}

func htmlFunc(app *iris.Application) {
	//app.RegisterView(iris.HTML("html", ".html"))
	app.RegisterView(iris.HTML("html", ".html").Binary(Asset, AssetNames))
}

func pageGet(ctx iris.Context) {
	err := ctx.View("upload_form.html", iris.Map{"ip": pageUploadIp, "token": pageToken})
	if err != nil {
		Log.Error("failed to view upload page. %s", err)
	}
}

func SetPageUploadIp(pageRemote bool) {
	if pageRemote {
		pageUploadIp, _ = utils.GetPublicIp()
		Log.Debug("get public ip: %s", pageUploadIp)
	}
	Log.Info("page upload ip: %s", pageUploadIp)
}