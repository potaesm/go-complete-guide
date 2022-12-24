package main

import "fmt"

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
	suthinan := person{
		firstName: "Suthinan",
		lastName:  "Musitmani",
		contactInfo: contactInfo{
			email:   "suthinan@gmail.com",
			zipCode: 20250,
		},
	}
	suthinan.updateName("Potae")
	suthinan.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
