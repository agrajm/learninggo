package main

import "fmt"

func main() {

	// make( []T , size, capacity )
	// create a optimized slice if you know the size in advance
	// capacity is the maximum size of the underlying array allocated for slice

	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) // capacity

	x[0] = 12
	x[9] = 34
	// x[10] = 123 // Error
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) // capacity

	x = append(x, 1232)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) // capacity

	// if the number of elements increases the capacity due to more appends then
	/*the compiler creates another underlying array copies the old values and
	discards the original array. the capacity of the new array is double the
	capacity of original*/

	// slices of slices
	a := []string{"one", "two", "three"}
	b := []string{"1", "2", "3"}
	c := [][]string{a, b}
	fmt.Println(c)

}
