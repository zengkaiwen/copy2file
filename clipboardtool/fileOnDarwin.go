package clipboardtool

import (
	"fmt"
	"strings"

	"github.com/josa42/go-applescript"
)

func darwinCopy2file(fileName string) error {
	scriptStr := strings.TrimSpace(
		fmt.Sprintf(`
			tell application "Finder"
					set the clipboard to POSIX file "%s"
			end tell`,
			fileName,
		),
	)
	applescript.Run(scriptStr)
	return nil
}
