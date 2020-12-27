package echo

import (
	"flag"
)

// NewLine 命令行参数帮助信息
var NewLine = flag.Bool("n", false, "print newline")

const (
	Space   = " "
	Newline = "\n"
)

// Echo 模拟Unix的echo命令
//
// -n 表示换行输出参数
func Echo() {
	// 打印帮助信息
	// flag.PrintDefaults()
	// flag.Parse()
	// var s string = ""
	// for i := 0; i < flag.NArg(); i++ {
	// 	if i > 0 {
	// 		s += " "
	// 		if *NewLine {
	// 			s += Newline
	// 		}
	// 	}
	// 	s += flag.Arg(i)
	// }
	// os.Stdout.WriteString(s)
}
