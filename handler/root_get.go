package handler

import (
	"crypto/md5"
	"fileserver/frame"
	. "fileserver/log"
	"fmt"
	"github.com/kataras/iris/v12"
	"io"
	"path/filepath"
	"strconv"
	"time"
)

func init() {
	//frame.RegisterHandler("Get", "/", page)
	frame.RegisterHandler("Get", "/{fileName:string}", download)
}

func page(ctx iris.Context) {
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
	fileName := ctx.Params().Get("fileName")
	Log.Debug("to find download file %s", fileName)
	filePath := filepath.Join(LocalDir, fileName)
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