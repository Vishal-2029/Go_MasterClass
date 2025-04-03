package main

import (
	"fmt"
	"time"
)

func main() {
	rate := time.Microsecond*1
	limiter := time.Tick(rate)

	for i := 0; i <= 10; i++ {
		<-limiter
		fmt.Println("Request",i, "at", time.TimeOnly)
	}
}