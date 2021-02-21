package main

import "fmt"

func main() {
	//name :="This name"
	//str,ok := name.(string) //invalid type assertion: name.(string) (non-interface type string on left)

	var name interface{} = "Name"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Println("value is not string")
	}

	var val interface{} = 3
	//fmt.Println(val + 1) //mismatched type
	fmt.Println(val.(int) + 1)

	rem := 7.24
	fmt.Printf("%T\n", rem)
	fmt.Printf("%T\n", int(rem))

	var val1 interface{} = 7
	fmt.Printf("%T\n", val1)
	// fmt.Printf("%T\n", int(val)) // need type assertion
	fmt.Printf("%T\n", val1.(int))
}
