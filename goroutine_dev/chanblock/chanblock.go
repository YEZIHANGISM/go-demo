package chanblock

import (
	"fmt"
	"time"
)

func ChanBlock() {
	ch := make(chan int)
	go pump(ch)

	// 无缓冲通道
	// 因为通道值没有接收者，发送者当前不可用，造成阻塞
	// fmt.Println(<-ch)

	// 解决办法
	go suck(ch)
	time.Sleep(1e9)
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
