package handler

import (
	"fileserver/frame"
	. "fileserver/log"
	"fileserver/utils"
	"github.com/kataras/iris/v12"
	"path/filepath"
)

var localFilesDir string

func init() {
	frame.RegisterHandler("Get", "/{fileName:string}", checkApiToken, download)
}

func download(ctx iris.Context) {
	fileName := ctx.Params().Get("fileName")
	Log.Debug("to find download file %s", fileName)
	filePath := filepath.Join(localFilesDir, fileName)
	if !utils.IsFileExist(filePath) {
		Log.Info("%s is not exist.", fileName)
		ctx.Writef("%s is not exist.", fileName)
		return
	}
	err := ctx.SendFile(filePath, fileName)
	if err != nil {
		Log.Error("failed to send file %s. %s", fileName, err)
		return
	}
	Log.Info("%s was downloaded.", fileName)
}

func SetLocalFilesDir(path string) {
	localFilesDir = path
}