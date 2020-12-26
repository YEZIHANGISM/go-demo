package readstring

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// ReadStringOrLine 按字符串读取文件内容
func ReadStringOrLine() {
	inputFile, inputError := os.Open("input.txt")
	if inputError != nil {
		fmt.Printf("An Error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		// inputString, readerError := inputReader.ReadString('\n')
		// fmt.Printf("the input was: %s", inputString)
		// if readerError == io.EOF {
		// 	return
		// }

		line, isPrefix, readerError := inputReader.ReadLine()
		if isPrefix {
			fmt.Printf("file begin with: %s\n", line)
		} else {
			fmt.Printf("the line content was: %s\n", line)
		}
		if readerError == io.EOF {
			return
		}
	}
}
