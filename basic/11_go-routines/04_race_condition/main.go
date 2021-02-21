package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func increment(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "counter:", counter)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go increment("Foo:")
	go increment("Bar:")
	wg.Wait()
	fmt.Println("Finale counter:", counter)
}

//try command:
// go run -race main.go
// go run main.go
