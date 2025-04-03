package main

import (
	"math/rand"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	_,err = nc.QueueSubscribe("task","workers", func (msg *nats.Msg)  {
		log.Println("Received:", string(msg.Data))

		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

		log.Println("Completed:", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	select{}
}