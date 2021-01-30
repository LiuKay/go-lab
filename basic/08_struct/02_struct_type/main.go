package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

/**
(p person) is the "receiver"
it is another parameter
not idiomatic to use "this" or "self"
*/
func (p person) fullName() string {
	return p.firstName + p.lastName
}

func main() {
	p1 := person{
		firstName: "Kay",
		lastName:  "Liu",
		age:       18,
	}

	fmt.Println(p1)
	fmt.Println(p1.fullName())
}
