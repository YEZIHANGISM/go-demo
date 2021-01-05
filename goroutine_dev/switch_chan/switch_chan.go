package switchchan

import (
	"fmt"
	"math"
	"os"
)

func tel(ch chan int, ch2 chan bool) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	ch2 <- true
}

// SwitchChan 切换协程
func SwitchChan() {
	ch := make(chan int)
	ch2 := make(chan bool)
	go tel(ch, ch2)
	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		case <-ch2:
			fmt.Println("done")
			os.Exit(0)
		}
	}
}

func fib(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fib(n-1) + fib(n-2)
	}
	return
}

func generateFib(ch chan int, quit chan bool) {
	for i := 2; i < 20; i++ {
		ch <- fib(i)
	}
	quit <- true
}

// Fibonacci 使用协程生成斐波那契数列
func Fibonacci() {
	ch := make(chan int)
	quit := make(chan bool)
	go generateFib(ch, quit)
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-quit:
			fmt.Println("done")
			os.Exit(0)
		}
	}
}

const constant float64 = 4

func expression(ch chan float64, k float64) {
	temp := constant * (math.Pow(-1, k) / (2*k + 1))
	fmt.Println(temp)
	ch <- temp
}

func concurrentPi(n int) float64 {
	ch := make(chan float64)
	for i := 0; i < n; i++ {
		go expression(ch, float64(i))
	}
	f := 0.0
	for j := 0; j < 1000; j++ {
		f += <-ch
	}
	return f
}

// Pi 使用协程计算pi的近似值
func Pi() {
	fmt.Println(concurrentPi(1000))
}
