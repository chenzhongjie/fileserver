package handler

import (
	. "fileserver/log"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
)

const localFilesDir = "files"
const tokenFileName = ".token"

var pageToken string
var apiToken string

func init() {
	initLocalDir()

	pageToken = newToken(10)
	Log.Debug("new page token: %s", pageToken)

	apiToken = newToken(10)
	Log.Info("new api token: %s", apiToken)
	filePath := filepath.Join(localFilesDir, tokenFileName)
	err := saveToken(apiToken, filePath)
	if err != nil {
		Log.Error("failed to save apiToken: %s", err)
		panic(err)
	}
	Log.Debug("saved api token in %s", filePath)
}

func isValidToken(ctx iris.Context, token string) (bool, string) {
	tokenParam := ctx.URLParam("token")
	Log.Debug("token param: %s", tokenParam)
	return tokenParam == token, tokenParam
}

func saveToken(token string, filePath string) error {
	return ioutil.WriteFile(filePath, []byte(token), 0666)
}

const letters = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
func newToken(n uint) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63() % int64(len(letters))]
	}
	return string(b)
}

func initLocalDir() {
	if isExist(localFilesDir) {
		return
	}
	err := os.Mkdir(localFilesDir, 0777)
	if err != nil {
		Log.Error("failed to make dir %s", localFilesDir)
		panic(err)
	}
	Log.Debug("make dir: %s", localFilesDir)
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}