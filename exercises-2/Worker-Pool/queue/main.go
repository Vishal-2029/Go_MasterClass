package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil{
		log.Fatal(err)
	}
	defer nc.Drain()

	for i := 1; i <= 10; i++ {
		task := fmt.Sprintf("Task %d", i)
		err := nc.Publish("task", []byte(task))
		if err != nil {
			log.Printf("Failed to publish task: %v", err)
		}else {
			log.Println("Publish:", task)
		}
		time.Sleep(time.Second)
	}

	log.Println("All tasks published.")
}