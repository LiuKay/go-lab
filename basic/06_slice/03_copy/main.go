package main

import "fmt"

func main() {
	fmt.Println("Test copy:")
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6, 7, 8, 9}
	fmt.Println("before copy:", slice1)
	copy(slice1, slice2)
	fmt.Println("After copy:", slice1)
}
