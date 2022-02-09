package main

import "fmt"

type contactinfo struct {
	email   string
	zipcode int
}
type person struct {
	firstName string
	lastName  string
	contactinfo
}

func main() {
	jim := person{
		"Jim",
		"Party",
		contactinfo{"lmeow", 30},
	}
	//jimPointer := &jim

	jim.updateName("Roln")
	jim.print()
}

func (p *person) updateName(newName string) {
	(*p).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
