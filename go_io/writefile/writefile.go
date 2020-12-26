package writefile

import (
	"bufio"
	"fmt"
	"os"
)

// WriteFile 写文件
func WriteFile() {
	outputFile, outputErr := os.OpenFile("writer.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputErr != nil {
		fmt.Println("err occurred")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	fmt.Println("write done")
	outputWriter.Flush()
}
