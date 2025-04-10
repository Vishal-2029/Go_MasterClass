package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0
var mutex sync.Mutex

func increment() {
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		counter = counter + 1 
		mutex.Unlock()
	}
}

func main() {
	go increment()
	go increment()

	time.Sleep(1 * time.Second)
	fmt.Println("Final Counter:", counter)
}
