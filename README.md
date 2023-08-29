# copy2file 工具

目前只支持 MacOS 系统，其他两个系统的 TODO 中

# 使用构建好的执行文件

1、去 release 里面下载符合自己系统的版本，解压后得到 `copy2file` 和 `LICENSE` 两个文件，然后在终端内去执行 `copy2file`

### MacOS

双击 `copy2file 文件`，可以直接在终端执行，第一次执行会出现两个权限允许项：

- 出现终端需要辅助功能权限，允许终端即可

- 出现需要访问 Finder 的弹窗，点击允许

这两个权限加好或允许之后，重新执行 `copy2file` 文件

> 注意，如果提示文件已损坏，或无法验证开发者，可执行 `sudo xattr -rd com.apple.quarantine /path/to/copy2file
` ，这里的 `/path/to/copy2file` 是你 Mac 上脚本所在的路径，根据你解压出来的路径设置即可。建议将 `copy2file` 文件丢到 `/Applications` （应用程序）目录下，然后执行 `sudo xattr -rd com.apple.quarantine /Applications/copy2file` 


或者在你自己喜欢的终端程序里面去找到文件路径，通过命令执行：

```bash
./copy2file
```


2、执行成功后，不要关闭终端，保持终端程序的开启

3、复制文本，按下快捷键”shift + option + f“ 就可以了


# 执行代码

1、安装好 go 环境
2、进入项目目录，执行 `go mod download`
3、项目目录下，执行 `go run .`
