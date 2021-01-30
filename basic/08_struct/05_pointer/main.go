package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := &person{
		name: "kay",
		age:  18,
	}
	fmt.Println(p1)
	fmt.Printf("type is %T \n", p1)
	fmt.Println(p1.name, p1.age)

	p2 := person{
		name: "jack",
		age:  19,
	}
	fmt.Println(p2)
	fmt.Printf("type is %T \n", p2)
	fmt.Println(p2.name, p2.age)

}
