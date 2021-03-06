package handler

import (
	"fileserver/frame"
	. "fileserver/log"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"math/rand"
	"path/filepath"
)

var Token string

func init() {
	Token = newToken(10)
	Log.Info("new token: %s", Token)
	filePath := filepath.Join(LocalDir, ".token")
	err := saveToken(Token, filePath)
	if err != nil {
		Log.Error("failed to save token: %s", err)
		panic(err)
	}
	Log.Info("saved token in %s", filePath)
	frame.RegisterMiddleware(checkToken)
}

func checkToken(ctx iris.Context) {
	tokenParam := ctx.URLParam("token")
	Log.Debug("token: %s", tokenParam)
	if tokenParam != Token {
		Log.Warn("%s %s token %s is wrong.", ctx.RemoteAddr(), ctx.FullRequestURI(), tokenParam)
		ctx.Writef("Authentication failed.")
		return
	}
	ctx.Next()
}

func saveToken(token string, filePath string) error {
	return ioutil.WriteFile(filePath, []byte(token), 0666)
}

const letters = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func newToken(len uint) string {
	return randSeq(len)
}

func randSeq(n uint) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63() % int64(len(letters))]
	}
	return string(b)
}