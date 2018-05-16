package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "James",
		lastName:  "Kirk",
		contactInfo: contactInfo{
			email:   "jkirk@starfleet.mil",
			zipCode: 80026,
		},
	}

	jim.updateName("Jim")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Println(p)
}
