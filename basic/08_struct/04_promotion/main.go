package main

import "fmt"

type person struct {
	Name string
}

func (p person) Greeting() string {
	return "person Greeting"
}

type teacher struct {
	person
	Name string
}

//func (t teacher) Greeting() string {
//	return "teacher Greeting"
//}

// fields and methods of the inner-type are promoted to the outer-type

func main() {
	p1 := person{Name: "p1"}

	t1 := teacher{
		person: person{
			Name: "p1",
		},
		Name: "t1",
	}

	fmt.Println(p1.Name)
	fmt.Println(t1.Name, t1.person.Name)
	fmt.Println(p1.Greeting())
	fmt.Println(t1.Greeting())

}
