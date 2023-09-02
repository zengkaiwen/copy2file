package clipboardtool

import (
	"fmt"
	"os/exec"
)

func winCopy2file(fileName string) error {
	cmd := exec.Command("./file2clip.exe", fileName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(output))
		return err
	}
	return nil
}
