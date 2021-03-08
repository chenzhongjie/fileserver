package main

import (
	"fileserver/frame"
	"fileserver/handler"
	. "fileserver/log"
	"fileserver/version"
	"flag"
	"fmt"
	"github.com/kataras/iris/v12"
	"os"
)

func main() {
	fPort := flag.String("port", "8080", "Listen port.")
	fLocalFilesDir := flag.String("fdir", "files", "Local relative dir to save all uploaded files.")
	fMaxFileSize := flag.Int64("maxsize", 2, "Max upload file size. Unit: G")
	fIsPageRemote := flag.Bool("remote", false, "Remote access for page upload.")
	fNoAuth := flag.Bool("noauth", false, "Remove token and no authentication.")
	fVersion := flag.Bool("version", false, "print filesever version.")
	flag.Parse()

	if *fVersion {
		fmt.Println(version.FullVersion())
		return
	}
	initLocalDir(*fLocalFilesDir)
	handler.SetLocalFilesDir(*fLocalFilesDir)
	if !*fNoAuth {
		handler.InitToken()
	}
	handler.SetPageRemoteUpload(*fIsPageRemote)
	frame.RegisterMiddleware(iris.LimitRequestBodySize(*fMaxFileSize<<30 + 1<<20))
	frame.Run(*fPort)
}

func initLocalDir(localFilesDir string) {
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