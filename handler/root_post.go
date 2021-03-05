package handler

import (
	"fileserver/frame"
	. "fileserver/log"
	"github.com/kataras/iris/v12"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func init() {
	frame.RegisterHandler("Post", "/", rootPost)
}

func rootPost(ctx iris.Context) {
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
		for _, file := range files {
			_, err = saveUploadedFile(file, "files")
			if err != nil {
				failures++
				ctx.Writef("failed to upload: %s", file.Filename)
			}
			Log.Info("%s has been saved.", file.Filename)
		}
	}
	ctx.Writef("%d files uploaded. %d failures", totals-failures, failures)
}

func saveUploadedFile(fh *multipart.FileHeader, destDirectory string) (int64, error) {
	src, err := fh.Open()
	if err != nil {
		return 0, err
	}
	defer src.Close()

	out, err := os.OpenFile(filepath.Join(destDirectory, fh.Filename),
		os.O_WRONLY|os.O_CREATE, os.FileMode(0666))
	if err != nil {
		return 0, err
	}
	defer out.Close()

	return io.Copy(out, src)
}