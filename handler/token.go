package handler

import (
	"fileserver/frame"
	. "fileserver/log"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"math/rand"
	"path/filepath"
)

const tokenFileName = ".token"

var pageToken string
var apiToken string
var isAuth = false

func InitToken() {
	apiToken = newToken(10)
	filePath := filepath.Join(localFilesDir, tokenFileName)
	err := saveToken(apiToken, filePath)
	if err != nil {
		Log.Error("failed to save apiToken: %s", err)
		panic(err)
	}
	Log.Info("saved api token %s in %s", apiToken, filePath)
	pageToken = newToken(10)
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

const letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
func newToken(n uint) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63() % int64(len(letters))]
	}
	return string(b)
}