package handler

import (
	"crypto/md5"
	"fileserver/frame"
	. "fileserver/log"
	"fmt"
	"github.com/kataras/iris/v12"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var LocalDir = "files"

func init() {
	frame.RegisterHandler("Get", "/", root)
	frame.RegisterHandler("Get", "/{filename:string}", download)
}

func root(ctx iris.Context) {
	//创建一个令牌（可选）。
	now := time.Now().Unix()
	h := md5.New()
	_, _ = io.WriteString(h, strconv.FormatInt(now, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	//使用令牌渲染表单以供您使用。
	//ctx.ViewData(""，token)
	//或者在`View`方法中添加第二个参数。

	//令牌将作为{{.}}传递到模板中。

	_ = ctx.View("upload_form.html", token)
}

func download(ctx iris.Context) {
	filePath := filepath.Join(LocalDir, ctx.Params().Get("filename"))
	if !isExist(filePath) {
		Log.Info("%s is not exist.", ctx.Params().Get("filename"))
		ctx.Writef("%s is not exist.", ctx.Params().Get("filename"))
		return
	}
	err := ctx.SendFile(filePath, ctx.Params().Get("filename"))
	if err != nil {
		Log.Error("failed to send file %s. %s", ctx.Params().Get("filename"), err)
		return
	}
	Log.Info("%s was downloaded.", ctx.Params().Get("filename"))
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}