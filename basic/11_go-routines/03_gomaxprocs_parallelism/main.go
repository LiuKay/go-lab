package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	//http://c.biancheng.net/view/94.html
	fmt.Println("NumCPU:", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("foo")
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("bar")
	}
	wg.Done()
}
