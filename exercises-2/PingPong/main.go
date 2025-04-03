package main

import (
	"fmt"
	"time"
)

func ping(Plyar1 chan bool, Plyar2 chan bool) {
	for i := 0; i < 10; i++ {
		<-Plyar1
		fmt.Println("Ping")
		time.Sleep(time.Second)
		Plyar2 <- true
	}
	close(Plyar2) 
}

func pong(Plyar2 chan bool, Plyar1 chan bool) {
	for range Plyar2 {
		fmt.Println("Pong")
		time.Sleep(time.Second)
		Plyar1 <- true
	}
}

func main() {
	Plyar1 := make(chan bool)
	Plyar2 := make(chan bool)

	go ping(Plyar1, Plyar2)
	go pong(Plyar2, Plyar1)

	Plyar1 <- true 

	time.Sleep(11 * time.Second) 
}