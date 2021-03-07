package main

import (
	"fileserver/frame"
	"fileserver/handler"
	"fileserver/version"
	"flag"
	"fmt"
	"github.com/kataras/iris/v12"
)

func main() {
	fPort := flag.String("port", "8080", "Listen port.")
	fMaxFileSize := flag.Int64("maxsize", 2, "Max upload file size. Unit: G")
	fPageRemote := flag.Bool("remote", false, "Remote access for page upload.")
	fVersion := flag.Bool("version", false, "print filesever version.")
	flag.Parse()

	if *fVersion {
		fmt.Println(version.FullVersion())
		return
	}

	handler.SetPageRemoteUpload(*fPageRemote)
	frame.RegisterMiddleware(iris.LimitRequestBodySize(*fMaxFileSize<<30 + 1<<20))
	frame.Run(*fPort)
}
