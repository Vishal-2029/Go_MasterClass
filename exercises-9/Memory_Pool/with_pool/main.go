package main

import (
	"fmt"
	"sync"
	"time"
)

type MyObject struct {
	Value int
}

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return &MyObject{}
		},
	}

	start := time.Now()

	for i := 0; i < 1_000_000; i++ {
		obj := pool.Get().(*MyObject)
		obj.Value = i
		
		pool.Put(obj)
	}

	fmt.Println("With pool:", time.Since(start))
}
