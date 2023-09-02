package main

import (
	"copy2file/clipboardtool"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/mitchellh/go-homedir"
	"golang.design/x/clipboard"
)

var separator = string(filepath.Separator)

func ClipboardFile() {
	data := string(clipboard.Read(clipboard.FmtText))
	if data == "" {
		log.Println("剪贴板内容为空")
		return
	}
	fileName, err := makeTmpFile(data)
	if err != nil {
		return
	}
	clipboardtool.CopyFileToClipboard(fileName)
}

func makeTmpFile(data string) (string, error) {
	var dir = "~/.tmp"
	var err error
	var expandedDir string
	expandedDir, err = homedir.Expand(dir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("expandedDir: %s\n", expandedDir)
	_, err = os.Stat(expandedDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(expandedDir, 0700)
		if err != nil {
			log.Fatalf("文件夹创建失败: %v", err)
		}
	}

	isJson := json.Valid([]byte(data))
	var fileName string
	currentTime := time.Now().UTC().Format("20060102150405")
	if isJson {
		fileName = fmt.Sprintf("%s%s%s.json", expandedDir, separator, currentTime)
	} else {
		fileName = fmt.Sprintf("%s%s%s.txt", expandedDir, separator, currentTime)
	}
	file, fileErr := os.Create(fileName)
	if fileErr != nil {
		log.Printf("文件创建失败: %v", fileErr)
		return "", fileErr
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		log.Printf("文件写入失败: %v", err)
		return "", err
	}

	err = file.Sync()
	if err != nil {
		log.Printf("文件同步失败: %v", err)
		return "", err
	}
	log.Println("文件创建成功:", file.Name())
	return file.Name(), nil
}
