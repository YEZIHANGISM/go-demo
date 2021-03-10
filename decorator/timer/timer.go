package timer

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// TimeFunc timefunc
type TimeFunc func(int64, int64) int64

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func timeSumFunc(f TimeFunc) TimeFunc {
	return func(start, end int64) int64 {

		defer func(t time.Time) {
			fmt.Printf("--- Time Elapsed (%s): %v ---\n",
				getFunctionName(f), time.Since(t))
		}(time.Now())

		return f(start, end)
	}
}

func sum1(start, end int64) int64 {
	var sum int64
	sum = 0
	if start > end {
		start, end = end, start
	}
	for i := start; i <= end; i++ {
		sum += i
	}
	return sum
}

func sum2(start, end int64) int64 {
	if start > end {
		start, end = end, start
	}
	return (end - start + 1) * (end + start) / 2
}

// Timer 计算函数执行时间
func Timer() {
	sum1 := timeSumFunc(sum1)
	sum2 := timeSumFunc(sum2)

	fmt.Printf("%d, %d\n", sum1(-10000, 10000000), sum2(-10000, 10000000))
}
