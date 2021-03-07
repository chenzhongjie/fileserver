package version

const (
	NAME string = "fileserver"
	VERSION string = "1.0.0"
)

func FullVersion() string {
	return NAME + " v" + VERSION
}