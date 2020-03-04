package main

import "fmt"

type contact struct {
	emai     string
	password string
}

type person struct {
	firstName string
	lastName  string
	age       int
	heigth    float32
	contact
	verified bool
}

func main() {
	theFlash := person{
		firstName: "Barry",
		lastName:  "Allen",
		contact: contact{
			emai:     "email.com",
			password: "123abc",
		},
	}
	// var theFlash person
	// perPointer := &theFlash
	theFlash.updateName("Wally West")
	theFlash.print()
}

// GO is a pass by value language
func (p person) print() {
	fmt.Println("this is an object", p)
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newName string) {
	(*p).firstName = newName
}
