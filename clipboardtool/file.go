package clipboardtool

import (
	"runtime"
)

func CopyFileToClipboard(fileName string) {
	var osType = runtime.GOOS
	if osType == "darwin" {
		darwinCopy2file(fileName)
	}
	if osType == "windows" {
		winCopy2file(fileName)
	}
	if osType == "linux" {
		// TODO:
	}
}
