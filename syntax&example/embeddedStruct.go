package main

import "fmt"

// base struct
type person struct {
	firstName string
	lastName  string
}

// method on person
func (p *person) talk() {
	fmt.Println("hi this is", p.firstName, p.lastName) // in GO, after "," it always put an extra space
}

// embedded embeds person
type embedded struct {
	person // embedded field -- no name
	age    int
}

func main() {
	// create embedded instance
	c := embedded{
		person: person{firstName: "Urmi", lastName: "Biswas"},
		age:    25,
	}
	c.person.talk() // accessing embedded field
	c.talk()        // calls directly
}
