package main

import (
	"fmt"
	"sync"
	"time"
)

// 用来使 main 线程等待其他2个协程执行完毕，然后打印执行时间，不然main 程序直接退出了，看不到其他协程执行结果
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	start := time.Now()
	go foo()
	go bar()
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))

	//total: 6s
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
