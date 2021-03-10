package handler

import (
	"errors"
	"fileserver/frame"
	. "fileserver/log"
	"fileserver/utils"
	"github.com/kataras/iris/v12"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func init() {
	frame.RegisterHandler("Post", "/", checkApiToken, upload)
}

func upload(ctx iris.Context) {
	// Get the max post value size in memory passed via iris.WithPostMaxMemory.
	// Too big file will be saved to tmp disk file and then to be copied to destDirectory.
	maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()
	err := ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}

	form := ctx.Request().MultipartForm
	totals := 0
	failures := 0
	for _, files := range form.File {
		totals += len(files)
		for _, fileHeader := range files {
			_, err = saveUploadedFile(fileHeader, localFilesDir)
			if err != nil {
				failures++
				Log.Info("failed to save %s: %s", fileHeader.Filename, err)
				ctx.Writef("failed to upload %s: %s\n", fileHeader.Filename, err)
				continue
			}
			Log.Info("%s has been saved.", fileHeader.Filename)
			ctx.Writef("%s has been uploaded.\n", fileHeader.Filename)
		}
	}
	Log.Info("[%s] %d files uploaded. %d failures.", ctx.RemoteAddr(), totals-failures, failures)
	ctx.Writef("%d files uploaded. %d failures.", totals-failures, failures)
}

func saveUploadedFile(fh *multipart.FileHeader, destDirectory string) (int64, error) {
	src, err := fh.Open()
	if err != nil {
		return 0, err
	}
	defer src.Close()

	var filePath = filepath.Join(destDirectory, fh.Filename)
	if utils.IsFileExist(filePath) {
		return 0, errors.New(fh.Filename + " is existing.")
	}
	out, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, os.FileMode(0666))
	if err != nil {
		return 0, err
	}
	defer out.Close()

	return io.Copy(out, src)
}