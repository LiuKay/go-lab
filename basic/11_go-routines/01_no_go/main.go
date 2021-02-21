package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	foo()
	bar()
	end := time.Now()
	fmt.Println(end.Sub(start))
	//total: 6s
}

func foo() {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("foo")
	}
}

func bar() {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("bar")
	}
}
