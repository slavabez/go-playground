package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println("Area of the shape:", s.getArea())
}

func main() {
	tri := triangle{base: 15, height: 10}
	sq := square{sideLength: 25}

	printArea(tri)
	printArea(sq)
}
