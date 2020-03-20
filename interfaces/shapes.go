package main

import "fmt"

type shape interface {
	getArea() float64
	printArea()
}

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (t triangle) printArea() {
	fmt.Println(t.getArea())
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (s square) printArea() {
	fmt.Println(s.getArea())
}

func main() {

	var s shape
	s = triangle{4.5, 5.5}
	s.printArea()
	s = square{5.0}
	s.printArea()
}
