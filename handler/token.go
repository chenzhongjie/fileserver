package handler

import (
	"fileserver/frame"
	. "fileserver/log"
	"fileserver/utils"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"path/filepath"
)

const tokenFileName = ".token"

var pageToken string
var apiToken string
var isAuth = false

func InitToken() {
	apiToken = utils.RandString(10)
	filePath := filepath.Join(localFilesDir, tokenFileName)
	err := saveToken(apiToken, filePath)
	if err != nil {
		Log.Error("failed to save apiToken: %s", err)
		panic(err)
	}
	Log.Info("saved api token %s in %s", apiToken, filePath)
	pageToken = utils.RandString(10)
	Log.Debug("new page token: %s", pageToken)

	frame.RegisterMiddleware(getTokenParam)
	isAuth = true
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

func checkPageToken(ctx iris.Context) {
	valid, tokenParam := isValidToken(ctx, pageToken)
	if valid {
		ctx.Next()
		return
	}
	Log.Warn("[%s] %s page token %s is wrong.", ctx.RemoteAddr(), ctx.FullRequestURI(), tokenParam)
	ctx.StatusCode(iris.StatusUnauthorized)
	ctx.Writef("Authentication failed.")
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

func isValidToken(ctx iris.Context, token string) (bool, string) {
	if !isAuth {
		return true, ""
	}
	tokenParam := ctx.Values().Get("token").(string)
	return tokenParam == token, tokenParam
}

func saveToken(token string, filePath string) error {
	return ioutil.WriteFile(filePath, []byte(token), 0666)
}