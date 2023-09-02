# copy2file 工具

## 介绍

当复制长文本，想粘贴到其他应用时，某些社交或工作应用不支持粘贴全部文本或只能粘贴一部分（例如：某信和企业某信），这个时候就需要将文本保存成 txt 文件粘贴到这些应用里面，然后发送。

该工具提供了一个快捷键，当复制完长文本后，按下 “shift + alt + g” 键之后 (Mac 里面是 shift + option + g)，就能将内容转换成文件，然后直接将文件粘贴到其他应用中。

目前支持 MacOS 和 Windows


## MacOS 使用说明

1、在右侧 release 里面选择最新版本下载 `copy2file-darwin.tar.gz
` ，解压后得到 `copy2file` 和 `LICENSE` 两个文件


2、双击 `copy2file 文件`，可以直接在终端执行，第一次执行会出现两个权限允许项：

> 除了双击也可以直接拷贝到“应用程序”目录下，可以在启动台点击运行；或者在任意终端里面，直接执行 `./copy2file`

- 出现终端需要辅助功能权限，在“设置”里面允许终端即可
  
> 具体操作：系统偏好设置 -> 辅助功能 -> 终端 -> 勾选“使用辅助功能”

- 出现需要访问 Finder 的弹窗，点击“允许”

这两个权限加好或允许之后，重新执行 `copy2file` 文件

> 注意，如果提示文件已损坏，或无法验证开发者，可执行 `sudo xattr -rd com.apple.quarantine /path/to/copy2file
` ，这里的 `/path/to/copy2file` 是你 Mac 上脚本所在的路径，根据你解压出来的路径设置即可。建议将 `copy2file` 文件丢到 `/Applications` （应用程序）目录下，然后执行 `sudo xattr -rd com.apple.quarantine /Applications/copy2file` 

## Windows 使用说明

1、在右侧 release 里面选择最新版本下载 `copy2file-windows.tar.gz` ，解压后得到 `copy2file.exe`、 `file2clip.exe`、`LICENSE` 三个文件，将这三个文件放到一个目录下即可

2、双击 `clip2file.exe` 文件打开终端窗口，此时运行成功，注意不要关闭这个终端窗口


## 下载

由于 Github 在国内被限制，Releases 内下载很慢，可以选择下方地址

蓝奏云下载地址：

[MacOS](https://wwi.lanzoup.com/i5Jop1799fqd)

[Windows](https://wwi.lanzoup.com/iBqNV1799fjg)

## 文件缓存说明

Mac 系统中，文件缓存在 `~/.tmp/` 目录下

Windows 系统中，文件缓存在 `%USERPROFILE%\.tmp\` 目录下，`%USERPROFILE%` 表示当前用户目录，通常是 `C:\Users\用户名` 。

## 源码说明

1、安装好 go 环境，源码go版本 1.20.7

2、进入项目目录，执行 `go mod download`

3、项目目录下，执行 `go run .`

> 注意，需要开启 CGO
> 
