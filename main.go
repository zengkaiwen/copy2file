package main

import (
	"fmt"
	"log"
	"runtime"

	hook "github.com/robotn/gohook"
	"golang.design/x/clipboard"
)

func main() {
	fmt.Println("")
	fmt.Println("================= 程序启动中 ==============")

	err := clipboard.Init()
	if err != nil {
		log.Fatalf("剪贴板初始化失败: %v", err)
	}
	osType := runtime.GOOS
	optKey := "alt"
	if osType == "darwin" {
		optKey = "option"
	}
	hook.Register(hook.KeyDown, []string{"g", optKey, "shift"}, func(e hook.Event) {
		// log.Printf("按下了快捷键: shift+%s+g", optKey)
		ClipboardFile()
	})
	s := hook.Start()

	fmt.Println("=【使用说明】")
	fmt.Println("= 1、复制任意文本")
	fmt.Printf("= 2、按下快捷键 shift + %s + g \n", optKey)
	fmt.Println("= 3、在其他支持文件粘贴的应用内粘贴")
	fmt.Println("================ 程序启动成功 =============")
	<-hook.Process(s)
}
