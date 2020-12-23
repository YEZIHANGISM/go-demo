package extract_interface

type Shaper interface {
	Area() float32
}

type TopologicalGenus interface {
	Rank() int
}

type Square