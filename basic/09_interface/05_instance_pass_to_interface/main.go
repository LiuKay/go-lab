package main

import "fmt"

type Integer int

//value
func (i Integer) Less(b Integer) bool {
	return i < b
}

//pointer
func (i *Integer) Add(b Integer) {
	*i += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

type Less interface {
	Less(b Integer) bool
}

type Adder interface {
	Add(b Integer)
}

func main() {
	a := Integer(1)

	//use which one?
	//var b LessAdder = &a
	//var b LessAdder = a

	var l1 Less = &a
	var l2 Less = a
	fmt.Println(l1, l2)

	var a1 Adder = &a
	//integer does not implement Adder (Add method has pointer receiver)
	//var a2 Adder = a //Cannot use 'a' (type Integer) as type Adder Type does not implement 'Adder' as 'Add' method has a pointer receiver
	fmt.Println(a1)

}
