#include <Windows.h>

void WriteFilePathToClipboard(const char* filePath) {
    // 打开剪贴板
    if (!OpenClipboard(NULL)) {
        // 处理错误
        return;
    }

    // 清空剪贴板
    EmptyClipboard();

    // 分配全局内存用于存储文件路径
    HGLOBAL hMem = GlobalAlloc(GMEM_MOVEABLE, strlen(filePath) + 1);
    if (hMem == NULL) {
        // 处理错误
        CloseClipboard();
        return;
    }

    // 锁定分配的内存
    char* pMem = (char*)GlobalLock(hMem);
    if (pMem == NULL) {
        // 处理错误
        GlobalFree(hMem);
        CloseClipboard();
        return;
    }

    // 将文件路径复制到锁定的内存中
    strcpy(pMem, filePath);

    // 解锁分配的内存
    GlobalUnlock(hMem);

    // 设置数据格式为文件路径
    SetClipboardData(CF_HDROP, hMem);

    // 关闭剪贴板
    CloseClipboard();
}
