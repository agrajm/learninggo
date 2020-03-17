package main

import "fmt"

func main() {

	todos := map[string]string{
		"red":   "#FF0000",
		"green": "#452312", // Note we need a , even on the last line
	}
	//fmt.Println(todos)

	// var peeps map[string]int
	// fmt.Print(peeps)
	// peeps2 := make(map[string]int)
	// fmt.Print(peeps2)

	// Iterate over Maps
	for key, value := range todos {
		fmt.Println("key: " + key)
		fmt.Println("value: " + value)
	}
}
