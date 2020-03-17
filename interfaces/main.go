package main

import "fmt"

type shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	length float64
}

func (r Rect) Area() float64 {
	return r.length * r.width
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {

	var shap shape
	shap = Rect{5.0, 4.2}
	rect := Rect{5.0, 4.2}
	fmt.Println("Area is ", shap.Area())
	fmt.Println("Perimeter is ", shap.Perimeter())
	fmt.Printf("Type of shap is %T ", shap)
	fmt.Println("Value of shap is ", shap)
	fmt.Println("shap==rect ", shap == rect)
}
