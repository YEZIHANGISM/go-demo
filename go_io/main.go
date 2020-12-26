package main

import (
	"fmt"
	"go_io/filecopy"
	readpractice "go_io/read_practice"
	"go_io/readentirefile"
	"go_io/readstring"
	writepractice "go_io/write_practice"
	"go_io/writefile"
)

func main() {
	readstring.ReadStringOrLine()
	fmt.Println("===================")
	readentirefile.ReadEntirefile()
	fmt.Println("===================")
	readpractice.SplitString()
	fmt.Println("===================")
	writefile.WriteFile()
	fmt.Println("===================")
	writepractice.WriteFromStruct()
	fmt.Println("===================")
	filecopy.Copy()

}
