package version

const (
	NAME string = "fileserver"
	VERSION string = "1.1.0"
)

func FullVersion() string {
	return NAME + " v" + VERSION
}