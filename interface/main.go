package main

import (
	"fmt"
	"interface/base_demo"
	"interface/empty_interface"
	"interface/go_reflect"
	"interface/methodset"
	"interface/sorting"
	"interface/type_assertion"
)

func main() {
	base_demo.BaseInterface()
	fmt.Println("====================")
	type_assertion.TypeAssertion()
	fmt.Println("====================")
	methodset.MethodSet()
	fmt.Println("====================")
	sorting.Sorting()
	fmt.Println("====================")
	empty_interface.EmptyInterface()
	fmt.Println("====================")
	go_reflect.ReflectDemo()
}
