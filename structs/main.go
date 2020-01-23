package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo //// same as contactInfo contactInfo -- no need for type here
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

// Call by Value - no effect on the actual person called with
func (p person) updateFirstName(newName string) {
	p.firstName = newName
}

// Pointer to Person -- Call by Name -- address is passed and original person updated
func (p *person) actuallyUpdateName(newName string) {
	(*p).firstName = newName
}

func main() {

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex.anderson@gmail.com",
			zipcode: 2145, // NOTE the comma after the last line -- IMPORTANT
		},
	}

	alex.updateFirstName("agraj")
	alex.print()

	// agraj := person{"Agraj", "Mangal"}
	// fmt.Printf("%+v", agraj)

	alex.actuallyUpdateName("Agraj")
	alex.print()

	// While calling by Pass by name (using pointer) both ways of calling are valid
	alexP := &alex
	alexP.actuallyUpdateName("This also works")
	alex.print()

}
