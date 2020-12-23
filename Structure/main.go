package main

import "fmt"

type contactInfo struct {
	zip   int
	email string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	// Use Case 1
	alex := person{
		firstName: "Shagun",
		lastName:  "Bandi",
		contact: contactInfo{
			zip:   1234,
			email: "sahguN@kjjk.com",
		},
	}
	fmt.Println(alex)

	// Use Case 2
	var ben person
	fmt.Printf("%+v\n", ben)

	ben.firstName = "Ben"
	fmt.Printf("%+v\n", ben)

	// Use Case 3
	jim := person{
		firstName: "Jim",
	}
	jim.updateFirstName("Jimmy")
	fmt.Printf("%+v", jim)
}

func (pointerToPerson *person) updateFirstName(firstName string) {
	(*pointerToPerson).firstName = firstName
}
