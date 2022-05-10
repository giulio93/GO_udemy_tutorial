package main

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
		edge: 5.0,
	}

	tr := triangle{
		name:   "triangle",
		base:   5.0,
		height: 9.0,
	}

	printArea(sq)
	printArea(tr)

}

func printArea(s shape) {

	shape_name, area := s.getArea()
	println("The area of the", shape_name, "is", area)

}

func (s square) getArea() (string, float64) {
	return s.name, s.edge * s.edge
}

func (t triangle) getArea() (string, float64) {
	return t.name, (t.base * t.height) / 2
}
