package readentirefile

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ReadEntirefile 读取整个文件内容至字符串中
func ReadEntirefile() {
	inputFile := "input.txt"
	outputFile := "input_bak.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, os.FileMode(0644))
	if err != nil {
		panic(err.Error())
	}
}
