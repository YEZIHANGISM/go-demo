package semaphore

import "fmt"

func sum(i, j int, ch chan int) {
	ch <- i + j
}

// GoSum 使用信号量计算两数之和
func GoSum() {
	// 信号量
	ch := make(chan int)
	go sum(12, 13, ch)
	fmt.Println(<-ch)
}

func producer(start, count int, ch1 chan int) {
	for i := 0; i < count; i++ {
		ch1 <- start
		start += count
	}
	close(ch1)
}

func consumer(ch1 chan int, done chan bool) {
	for num := range ch1 {
		fmt.Printf("%d\n", num)
	}
	done <- true
}

func ProducerAndConSumer() {
	ch1 := make(chan int)
	done := make(chan bool)
	go producer(0, 10, ch1)
	go consumer(ch1, done)

	<-done
}
