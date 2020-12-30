package primefilter

import "fmt"

func generate(ch chan int) {
	for i := 2; i < 1000; i++ {
		ch <- i
	}
	close(ch)
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

// PrimeFilter 素数筛选
func PrimeFilter() {
	ch := make(chan int)
	go generate(ch)
	for {
		prime := <-ch
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}
