package main

import (
	"runtime"
	"syscall"
	"time"

	"github.com/lxn/win"
)

func copyFileToClipboardForWin(filePath string) error {
	return nil
}

func writrFile(filePath string) error {
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
	// 读取文件内容
	fileContent := make([]byte, 1024)
	var bytesRead uint32
	err = syscall.ReadFile(syscall.Handle(fileHandle), fileContent, &bytesRead, nil)
	if err != nil {
		return err
	}
	// 打开剪贴板
	if !openC(0) {
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

func write(file string) (<-chan struct{}, error) {
	errch := make(chan error)
	changed := make(chan struct{}, 1)
	go func() {
		// make sure GetClipboardSequenceNumber happens with
		// OpenClipboard on the same thread.
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		for {
			r, _, _ := openClipboard.Call(0)
			if r == 0 {
				continue
			}
			break
		}

		// var param uintptr
		err := writrFile(file)
		if err != nil {
			errch <- err
			closeClipboard.Call()
			return
		}
		// Close the clipboard otherwise other applications cannot
		// paste the data.
		closeClipboard.Call()

		cnt, _, _ := getClipboardSequenceNumber.Call()
		errch <- nil
		for {
			time.Sleep(time.Second)
			cur, _, _ := getClipboardSequenceNumber.Call()
			if cur != cnt {
				changed <- struct{}{}
				close(changed)
				return
			}
		}
	}()
	err := <-errch
	if err != nil {
		return nil, err
	}
	return changed, nil
}

const (
	cFmtBitmap      = 2 // Win+PrintScreen
	cFmtUnicodeText = 13
	cFmtDIBV5       = 17
	// Screenshot taken from special shortcut is in different format (why??), see:
	// https://jpsoft.com/forums/threads/detecting-clipboard-format.5225/
	cFmtDataObject = 49161 // Shift+Win+s, returned from enumClipboardFormats
	gmemMoveable   = 0x0002
)

// BITMAPV5Header structure, see:
// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapv5header
type bitmapV5Header struct {
	Size          uint32
	Width         int32
	Height        int32
	Planes        uint16
	BitCount      uint16
	Compression   uint32
	SizeImage     uint32
	XPelsPerMeter int32
	YPelsPerMeter int32
	ClrUsed       uint32
	ClrImportant  uint32
	RedMask       uint32
	GreenMask     uint32
	BlueMask      uint32
	AlphaMask     uint32
	CSType        uint32
	Endpoints     struct {
		CiexyzRed, CiexyzGreen, CiexyzBlue struct {
			CiexyzX, CiexyzY, CiexyzZ int32 // FXPT2DOT30
		}
	}
	GammaRed    uint32
	GammaGreen  uint32
	GammaBlue   uint32
	Intent      uint32
	ProfileData uint32
	ProfileSize uint32
	Reserved    uint32
}

type bitmapHeader struct {
	Size          uint32
	Width         uint32
	Height        uint32
	PLanes        uint16
	BitCount      uint16
	Compression   uint32
	SizeImage     uint32
	XPelsPerMeter uint32
	YPelsPerMeter uint32
	ClrUsed       uint32
	ClrImportant  uint32
}

// Calling a Windows DLL, see:
// https://github.com/golang/go/wiki/WindowsDLLs
var (
	user32 = syscall.MustLoadDLL("user32")
	// Opens the clipboard for examination and prevents other
	// applications from modifying the clipboard content.
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-openclipboard
	openClipboard = user32.MustFindProc("OpenClipboard")
	// Closes the clipboard.
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-closeclipboard
	closeClipboard = user32.MustFindProc("CloseClipboard")
	// Empties the clipboard and frees handles to data in the clipboard.
	// The function then assigns ownership of the clipboard to the
	// window that currently has the clipboard open.
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-emptyclipboard
	emptyClipboard = user32.MustFindProc("EmptyClipboard")
	// Retrieves data from the clipboard in a specified format.
	// The clipboard must have been opened previously.
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclipboarddata
	getClipboardData = user32.MustFindProc("GetClipboardData")
	// Places data on the clipboard in a specified clipboard format.
	// The window must be the current clipboard owner, and the
	// application must have called the OpenClipboard function. (When
	// responding to the WM_RENDERFORMAT message, the clipboard owner
	// must not call OpenClipboard before calling SetClipboardData.)
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setclipboarddata
	setClipboardData = user32.MustFindProc("SetClipboardData")
	// Determines whether the clipboard contains data in the specified format.
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isclipboardformatavailable
	isClipboardFormatAvailable = user32.MustFindProc("IsClipboardFormatAvailable")
	// Clipboard data formats are stored in an ordered list. To perform
	// an enumeration of clipboard data formats, you make a series of
	// calls to the EnumClipboardFormats function. For each call, the
	// format parameter specifies an available clipboard format, and the
	// function returns the next available clipboard format.
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isclipboardformatavailable
	enumClipboardFormats = user32.MustFindProc("EnumClipboardFormats")
	// Retrieves the clipboard sequence number for the current window station.
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclipboardsequencenumber
	getClipboardSequenceNumber = user32.MustFindProc("GetClipboardSequenceNumber")
	// Registers a new clipboard format. This format can then be used as
	// a valid clipboard format.
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-registerclipboardformata
	registerClipboardFormatA = user32.MustFindProc("RegisterClipboardFormatA")

	kernel32 = syscall.NewLazyDLL("kernel32")

	// Locks a global memory object and returns a pointer to the first
	// byte of the object's memory block.
	// https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globallock
	gLock = kernel32.NewProc("GlobalLock")
	// Decrements the lock count associated with a memory object that was
	// allocated with GMEM_MOVEABLE. This function has no effect on memory
	// objects allocated with GMEM_FIXED.
	// https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalunlock
	gUnlock = kernel32.NewProc("GlobalUnlock")
	// Allocates the specified number of bytes from the heap.
	// https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalalloc
	gAlloc = kernel32.NewProc("GlobalAlloc")
	// Frees the specified global memory object and invalidates its handle.
	// https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalfree
	gFree   = kernel32.NewProc("GlobalFree")
	memMove = kernel32.NewProc("RtlMoveMemory")
)
