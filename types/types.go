package main

import "fmt"

var x int

type hotdog int

var y hotdog

func main() {

	x = 12
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)

	y = 112
	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)

	y = hotdog(x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)

}
