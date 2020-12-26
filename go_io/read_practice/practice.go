package readpractice

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type result struct {
	title   string
	price   string
	quntity string
}

// SplitString 分割字符串
func SplitString() {
	file := "products.txt"
	inputFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	res := []*result{}
	for {
		inputString, readerErr := inputReader.ReadString('\n')
		stringSlice := strings.Split(inputString, ";")
		re := &result{stringSlice[0], stringSlice[1], stringSlice[2]}
		fmt.Println(re)
		res = append(res, re)
		fmt.Println(res)
		if readerErr == io.EOF {
			return
		}
	}
}
