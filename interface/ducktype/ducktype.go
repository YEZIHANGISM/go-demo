package ducktype

import "fmt"

type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck) {
	for i := 1; i <= 3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

type Bird struct{}

func (b Bird) Quack() {
	fmt.Println("quacking")
}
func (b Bird) Walk() {
	fmt.Println("walking")
}

func DuckType() {
	b := new(Bird)
	DuckDance(b)
}
