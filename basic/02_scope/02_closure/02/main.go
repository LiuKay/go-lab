package main

import "fmt"

/**
闭包提供了一种在方法内部重用代码的能力，当 increment 不想暴露给其他方法使用时可以使用这种方法将变量和方法的作用域限制在方法内
*/
func main() {
	x := 10
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
