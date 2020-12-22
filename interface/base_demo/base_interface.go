package base_demo

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	value int
}

func (s *Simple) Get() int {
	return s.value
}

func (s *Simple) Set(v int) {
	s.value = v
	return
}

func SimpleDemo(s *Simple) {
	value := s.Get()
	fmt.Println(value)
	s.Set(3)
	fmt.Println(s.Get())
}

// BaseInterface 基础接口DEMO
func BaseInterface() {
	s := &Simple{7}
	SimpleDemo(s)
}