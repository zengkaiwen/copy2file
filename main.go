package main

import (
	"fmt"
	"log"
	"runtime"

	hook "github.com/robotn/gohook"
	"golang.design/x/clipboard"
)

func main() {
	fmt.Println("程序启动中...")
	fmt.Println("目前仅支持 MacOS 系统")

	err := clipboard.Init()
	if err != nil {
		log.Fatalf("剪贴板初始化失败: %v", err)
	}
	osType := runtime.GOOS
	optKey := "alt"
	if osType == "darwin" {
		optKey = "option"
	}
	hook.Register(hook.KeyDown, []string{"f", optKey, "shift"}, func(e hook.Event) {
		log.Printf("按下了快捷键: shift+%s+f", optKey)
		ClipboardFile()
	})
	s := hook.Start()

	fmt.Printf("复制任意文本之后，按下快捷键 shift+%s+f 即可生成文件并拷贝到剪贴板中\n", optKey)
	<-hook.Process(s)
}
