package demo1

import "fmt"

func decorator(f func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("started")
		f(s)
		fmt.Println("ended")
	}
}

func hello(s string) {
	fmt.Println(s)
}

// Demo1 demo1
func Demo1() {
	// GO没有python @decorator类似的语法糖
	// 但本质是一样的，都是用高阶函数包装目标函数
	decorator(hello)("hello world")

	decorator := decorator(hello)
	decorator("hello world")
}
