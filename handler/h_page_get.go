package handler

import (
	"fileserver/frame"
	. "fileserver/log"
	"github.com/kataras/iris/v12"
	"github.com/xjh22222228/ip"
)

var publicIp string

func init() {
	frame.RegisterHandler("Get", "/", pageGet)
	//publicIp, _ = getPublicIp()
	//Log.Info("get public ip: %s", publicIp)
	publicIp = "localhost" // for local test
}

func pageGet(ctx iris.Context) {
	err := ctx.View("upload_form.html", iris.Map{"ip": publicIp, "token": pageToken})
	if err != nil {
		Log.Error("failed to view upload page. %s", err)
	}
}

func getPublicIp() (string, error) {
	ipv4, err := ip.V4()
	if err != nil {
		return "", err
	}
	return ipv4, nil
}