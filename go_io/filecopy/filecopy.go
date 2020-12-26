package filecopy

import (
	"fmt"
	"io"
	"os"
)

// Copy 拷贝文件
func Copy() {
	copyFile("page_bak.txt", "page.txt")
	fmt.Println("copy done")
}

func copyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
