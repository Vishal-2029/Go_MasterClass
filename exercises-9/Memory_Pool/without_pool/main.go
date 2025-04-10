package main

import (
	"fmt"
	"time"
)

type MyObject struct {
	Value int
}

func main() {
	start := time.Now()

	for i := 0; i < 1000000; i++ {
		obj := &MyObject{Value: i}
		_ = obj 
	}

	fmt.Println("Without pool:", time.Since(start))
}
