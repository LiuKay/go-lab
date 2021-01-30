package main

import "fmt"

type foo int

func main() {
	var myFoo foo
	myFoo = 12
	fmt.Printf("%T %v \n", myFoo, myFoo)

	var yourFoo int
	yourFoo = 13
	fmt.Printf("%T %v \n", yourFoo, yourFoo)

	//fmt.Println(myFoo + yourFoo) // invalid operation: myFoo + yourFoo (mismatched types foo and int)

	fmt.Println(int(myFoo) + yourFoo)
}
