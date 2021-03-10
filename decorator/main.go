package main

import (
	"decorator/demo1"
	"decorator/timer"
	"fmt"
)

func main() {
	demo1.Demo1()
	fmt.Println("==================================")
	timer.Timer()
	fmt.Println("============================")

}
