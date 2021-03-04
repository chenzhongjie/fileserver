package main

import (
	"fileserver/frame"
	_ "fileserver/handler"
	_ "fileserver/html"
	"fileserver/version"
	"flag"
	"fmt"
	"github.com/kataras/iris/v12"
)

var Port = flag.String("port", "8080", "Listen port. Default: 8080")
var MaxFileSize = flag.Int64("maxsize", 1, "Max uploaded file size. Default: 1G")

func main() {
	fmt.Printf("start %s %s ...\n", version.NAME, version.VERSION)
	flag.Parse()

	frame.RegisterMiddleware(iris.LimitRequestBodySize((*MaxFileSize + 1)<<30))
	frame.Run(*Port)
}
