package empty_interface

import "fmt"

type Element interface{}

type Vector struct {
	a []Element
}

func (p *Vector) At(i int) Element {
	return p.a[i]
}

func (p *Vector) Set(i int, e Element) {
	p.a[i] = e
}

func EmptyInterface() {
	e1 := Element(1)
	e2 := Element("hello")
	e3 := Element(complex(1, 2))
	data := []Element{e1, e2, e3}
	v := Vector{data}
	fmt.Println(v)
}
