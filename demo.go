package main

import (
	"log"
	"syscall"

	"github.com/lxn/win"
)

func main() {
	if err := copyFileToClipboard("C:\\path\\to\\file.jpg"); err != nil {
		log.Fatal(err)
	}
}
func copyFileToClipboard(filePath string) error {
	// 打开文件
	file, err := syscall.UTF16PtrFromString(filePath)
	if err != nil {
		return err
	}
	fileHandle, err := syscall.CreateFile(file, syscall.GENERIC_READ, 0, nil, syscall.OPEN_EXISTING, syscall.FILE_ATTRIBUTE_NORMAL, 0)
	if err != nil {
		return err
	}
	defer syscall.CloseHandle(fileHandle)
	// 获取文件大小
	var fileSize int64
	err = syscall.GetFileSizeEx(syscall.Handle(fileHandle), &fileSize)
	if err != nil {
		return err
	}
	// 读取文件内容
	fileContent := make([]byte, fileSize)
	var bytesRead int
	err = syscall.ReadFile(syscall.Handle(fileHandle), fileContent, &bytesRead, nil)
	if err != nil {
		return err
	}
	// 打开剪贴板
	if !win.OpenClipboard(0) {
		return syscall.GetLastError()
	}
	defer win.CloseClipboard()
	// 清空剪贴板
	if !win.EmptyClipboard() {
		return syscall.GetLastError()
	}
	// 分配全局内存，并将文件内容复制到全局内存中
	hMem := win.GlobalAlloc(win.GMEM_MOVEABLE, uintptr(len(fileContent)))
	if hMem == 0 {
		return syscall.GetLastError()
	}
	defer win.GlobalFree(hMem)
	memPtr := win.GlobalLock(hMem)
	if memPtr == nil {
		return syscall.GetLastError()
	}
	defer win.GlobalUnlock(hMem)
	copy((*[1 << 30]byte)(memPtr)[:len(fileContent)], fileContent)
	// 设置剪贴板数据格式
	win.SetClipboardData(win.CF_DIB, win.HANDLE(hMem))
	if win.GetLastError() != 0 {
		return syscall.GetLastError()
	}
	return nil
}
