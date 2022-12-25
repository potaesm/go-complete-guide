package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	side float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	sq := square{side: 10}
	t := triangle{base: 10, height: 5}
	fmt.Println(printArea(sq))
	fmt.Println(printArea(t))
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) float64 {
	return s.getArea()
}
