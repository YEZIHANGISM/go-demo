package readcmdargs

import (
	"fmt"
	"os"
	"strings"
)

// ReadCmdArgs 读取命令行参数
func ReadCmdArgs() {
	printString := "Hello "
	if len(os.Args) > 1 {
		printString += strings.Join(os.Args[1:], " ")
	}
	fmt.Println(printString)
}
