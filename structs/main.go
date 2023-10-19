package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// type person struct {
// 	firstName string
// 	lastName  string
// 	contact   contactInfo
// }

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func (pointerTOPerson *person) updateLastName(lastName string) {
	(*pointerTOPerson).lastName = lastName
	fmt.Printf("%+v \n", (*pointerTOPerson))
}

func main() {
	// harish := person{
	// 	firstName: "Harish",
	// 	lastName:  "Madi",
	// 	contactInfo: contactInfo{
	// 		email:   "afa@gmail.com",
	// 		zipCode: 123456,
	// 	},
	// }
	// fmt.Println(harish)
	// harish.print()

	var newPerson person

	newPerson.firstName = "Harish"
	newPerson.lastName = "Harish"
	// fmt.Println(newPerson)
	// fmt.Printf("%+v \n", newPerson)
	newPerson.print()

	// newPRef := &newPerson
	// newPRef.updateLastName("Madi")

	newPerson.updateLastName("Madi")
	newPerson.print()
}
