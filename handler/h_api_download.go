package handler

import (
	"fileserver/frame"
	. "fileserver/log"
	"github.com/kataras/iris/v12"
	"os"
	"path/filepath"
)

var localFilesDir string

func init() {
	frame.RegisterHandler("Get", "/{fileName:string}", checkApiToken, download)
}

func checkApiToken(ctx iris.Context) {
	valid, tokenParam := isValidToken(ctx, apiToken)
	if valid {
		ctx.Next()
		return
	}
	Log.Warn("[%s] %s api token %s is wrong.", ctx.RemoteAddr(), ctx.FullRequestURI(), tokenParam)
	ctx.StatusCode(iris.StatusUnauthorized)
	ctx.Writef("Authentication failed.")
}

func download(ctx iris.Context) {
	fileName := ctx.Params().Get("fileName")
	Log.Debug("to find download file %s", fileName)
	filePath := filepath.Join(localFilesDir, fileName)
	if !isExist(filePath) {
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

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func SetLocalFilesDir(path string) {
	localFilesDir = path
}