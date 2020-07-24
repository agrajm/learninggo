package main

import "fmt"

func main() {

	x := map[string]int{
		"James Bond":      32,
		"Miss MoneyPenny": 21,
	}

	fmt.Println(x)
	fmt.Println(x["James Bond"])
	fmt.Println(x["Not a Valid Value"]) // will give us ZERO

	// , ok idiom
	v, ok := x["Not a Valid Value"]
	fmt.Println(v)  // ZERO
	fmt.Println(ok) // False

	if x, ok := x["James Bond"]; ok {
		fmt.Println(x)
	}

}
