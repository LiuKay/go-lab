package main

import (
	"fmt"
)

func main() {
	x := 10
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "This is y"
		fmt.Println(y)
	}
	//	can not access y
	//fmt.Println(y)
}
