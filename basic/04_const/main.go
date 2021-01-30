package main

import "fmt"

// 这是一个枚举
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays // private
)

func main() {
	const name = "kay"
	fmt.Printf("const name:%s \n", name)

	fmt.Println(Sunday)
}
