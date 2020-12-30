package main

import (
	primefilterfactory "goroutine_dev/prime_filter_factory"
)

func main() {
	// channels.ChanVar()
	// chanblock.ChanBlock()
	// deadlock.DeadLock()
	// bufchan.ChanWithBuf()
	// semaphore.GoSum()
	// semaphore.ProducerAndConSumer()
	// primefilter.PrimeFilter()
	primefilterfactory.PrimeFilterFactory()
}
