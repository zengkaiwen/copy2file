# copy2file 工具

目前只支持 MacOS 系统，其他两个系统的 TODO 中

# 使用构建好的执行文件

1、去 release 里面下载符合自己系统的版本，解压后得到 `copy2file` 和 `LICENSE` 两个文件，然后在终端内去执行 `copy2file`

### MacOS

出现需要访问 Finder 的弹窗，点击允许


```bash
./copy2file
```

2、执行成功后，不要关闭终端，保持终端程序的开启

3、复制文本，按下快捷键”shift + option + f“ 就可以了


# 执行代码

1、安装好 go 环境
2、进入项目目录，执行 `go mod download`
3、项目目录下，执行 `go run .`
