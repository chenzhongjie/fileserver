package main

import (
	"fileserver/frame"
	"fileserver/handler"
	_ "fileserver/html"
	"flag"
	"github.com/kataras/iris/v12"
)

func main() {
	port := flag.String("port", "8080", "Listen port.")
	maxFileSize := flag.Int64("maxsize", 2, "Max upload file size. Unit: G")
	pageRemote := flag.Bool("remote", false, "Remote access for page upload.")
	flag.Parse()

	handler.SetPageRemoteUpload(*pageRemote)
	frame.RegisterMiddleware(iris.LimitRequestBodySize(*maxFileSize<<30 + 1<<20))
	frame.Run(*port)
}
