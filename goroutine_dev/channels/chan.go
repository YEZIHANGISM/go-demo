package channels

import (
	"fmt"
	"time"
)

func ChanVar() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)

	// 如果不等待，sendData()执行后立刻执行getData()
	// getData()循环从chan中获取数据，没有得到就立刻返回
	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Println(input)
	}
}
