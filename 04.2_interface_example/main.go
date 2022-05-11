package main

import "fmt"

type shape interface {
	getArea() (string, float64)
}
type square struct {
	name string
	edge float64
}
type triangle struct {
	name   string
	base   float64
	height float64
}

func main() {
	sq := square{
		name: "square",
		edge: 10,
	}

	tr := triangle{
		name:   "triangle",
		base:   10,
		height: 10,
	}

	printArea(sq)
	printArea(tr)

}

func printArea(s shape) {

	shape_name, area := s.getArea()
	fmt.Println("The area of the", shape_name, "is", area)

}

func (s square) getArea() (string, float64) {
	return s.name, s.edge * s.edge
}

func (t triangle) getArea() (string, float64) {
	return t.name, (t.base * t.height) * 0.5
}
