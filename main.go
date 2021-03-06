package main

import (
	"fileserver/frame"
	_ "fileserver/handler"
	_ "fileserver/html"
	. "fileserver/log"
	"fileserver/version"
	"flag"
	"github.com/kataras/iris/v12"
)

func main() {
	port := flag.String("port", "8080", "Listen port.")
	maxFileSize := flag.Int64("maxsize", 2, "Max uploaded file size. Unit: G")
	flag.Parse()

	Log.Info("start %s %s ...", version.NAME, version.VERSION)
	frame.RegisterMiddleware(iris.LimitRequestBodySize((*maxFileSize + 1)<<30))
	frame.Run(*port)
}
