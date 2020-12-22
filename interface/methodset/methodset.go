package methodset

import "fmt"

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func MethodSet() {
	var lst List

	// 不能调用CountInto，因为Append()方法只定义在指针上，但lst是一个值
	//CountInto(lst, 1, 10)
	if LongEnough(lst) {
		fmt.Printf("- lst is long enough\n")
	}

	// 可以调用CountInto，因为new()返回一个指针对象，plst是一个指针对象
	plst := new(List)
	CountInto(plst, 1, 10)
	if LongEnough(plst) {
		fmt.Printf("- plst is long enough\n")
	}
}