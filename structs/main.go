package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	jim := person{firstName: "Jim",
		lastName: "Party",
		contact: contactInfo{
			email: "j@j.com",
			zip:   23423, //In a multi-line struct, slice, array or map literal, every line must end with a comma.
		},
	}

	//using pointers -------------------------------------------
	//approach 1
	pJim := &jim //get address of the Jim, assign to pJim
	pJim.updateName("NewJim")
	jim.print()

	//approach 2
	jim.updateName("NewJim2") //this is a shortcut in go compiler, so it automatically gets the address of the type based on the receiver
	jim.print()
}

// Receiver func with pointer
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName //this also works, a short cut in go compiler: p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
