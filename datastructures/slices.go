package main

import "fmt"

func main() {

	x := []string{"agraj", "says", "hello", "to", "you"}
	fmt.Println(x)

	// looping
	for index, element := range x {
		fmt.Println(index, element)
	}

	// slicing a slice
	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	// Append to a slice
	x = append(x, "hello", "again")
	fmt.Println(x)

	// Appending one slice to another using variadic operand
	y := []string{"another", "slice"}
	x = append(x, y...)
	fmt.Println(x)

	// slices of slices
	a := []string{"one", "two", "three"}
	b := []string{"1", "2", "3"}
	c := [][]string{a, b}
	fmt.Println(c)

}
