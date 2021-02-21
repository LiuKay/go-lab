package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

var mutex sync.Mutex

func increment(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s, i, "counter:", counter)
		mutex.Unlock()
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
