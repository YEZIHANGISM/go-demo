package deadlock

import "fmt"

func f1(in chan int) {
	// 通道in发送一个值，但没有接收者，造成死锁
	fmt.Println(<-in)
}

func DeadLock() {
	// 死锁：
	// 无缓冲通道out接受了一个值，造成了通道阻塞
	out := make(chan int)
	out <- 2
	go f1(out)
}
