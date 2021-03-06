package handler

import (
	"fileserver/frame"
	. "fileserver/log"
	"github.com/kataras/iris/v12"
	"github.com/xjh22222228/ip"
)

var pageUploadIp = "127.0.0.1"

func init() {
	frame.RegisterHandler("Get", "/", pageGet)
}

func pageGet(ctx iris.Context) {
	err := ctx.View("upload_form.html", iris.Map{"ip": pageUploadIp, "token": pageToken})
	if err != nil {
		Log.Error("failed to view upload page. %s", err)
	}
}

func SetPageRemoteUpload(pageRemote bool) {
	if pageRemote {
		pageUploadIp, _ = getPublicIp()
		Log.Debug("get public ip: %s", pageUploadIp)
	}
	Log.Info("page upload ip: %s", pageUploadIp)
}

func getPublicIp() (string, error) {
	ipv4, err := ip.V4()
	if err != nil {
		return "", err
	}
	return ipv4, nil
}