package handler

import (
	"fileserver/frame"
)

func init() {
	frame.RegisterHandler("Post", "/page", checkPageToken, upload)
}